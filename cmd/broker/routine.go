package main

import (
	bus "monitor/dbus"
	"monitor/dbus/systemd"
	"monitor/service"
	"monitor/slices"
	"path"
	"strings"

	"github.com/godbus/dbus/v5"
)

func getServiceUnitFiles() []systemd.UnitFile {
	unitFiles, err := bus.CallSystem[systemd.ListUnitFiles]()
	if err != nil {
		panic(err)
	}

	return slices.Filter(unitFiles, func(uf systemd.UnitFile) bool {
		return !strings.HasSuffix(uf.Path, `@.service`) && strings.HasSuffix(uf.Path, `.service`)
	})
}

func updateServices(services []service.Service) []service.Service {
	unitFiles := getServiceUnitFiles()
	units, err := bus.CallSystem[systemd.ListUnits]()
	if err != nil {
		panic(err)
	}

	updates := slices.MapSlice(unitFiles, func(unitFile systemd.UnitFile) []service.Service {
		n := path.Base(unitFile.Path)

		unitPath, err := bus.CallSystem[systemd.LoadUnit](n)
		if err != nil {
			panic(err)
		}

		filtered := slices.Filter(services, func(s service.Service) bool {
			return s.ObjectPath == string(unitPath)
		})
		if len(filtered) == 0 {
			filtered = []service.Service{{
				Name:       n,
				ObjectPath: string(unitPath),
				State:      service.StateDead,
			}}
		}

		unit, found := slices.First(units, func(unit systemd.Unit) bool {
			return string(unit.UnitPath) == string(unitPath)
		})
		if !found {
			return slices.Map(filtered, func(s service.Service) service.Service {
				s.State = service.StateDead
				return s
			})
		}

		state, err := service.StateString(unit.SubState)
		if err != nil {
			return filtered
		}

		filtered = slices.Map(filtered, func(s service.Service) service.Service {
			s.State = state
			return s
		})

		return filtered
	})

	tx := qdb.MustBegin()
	defer tx.Rollback() //nolint: errcheck
	for _, update := range updates {
		service.Upsert(tx, update)
	}
	tx.MustCommit()

	return updates
}

func markUnavailableAsDead(services []service.Service) {
	unitFiles := getServiceUnitFiles()
	unitPaths := slices.Map(unitFiles, func(unitFile systemd.UnitFile) dbus.ObjectPath {
		n := path.Base(unitFile.Path)

		unitPath, err := bus.CallSystem[systemd.LoadUnit](n)
		if err != nil {
			panic(err)
		}

		return dbus.ObjectPath(unitPath)
	})

	unavailable := slices.Filter(services, func(s service.Service) bool {
		return !slices.Contains(unitPaths, func(unitPath dbus.ObjectPath) bool {
			return s.ObjectPath == string(unitPath)
		})
	})
	unavailable = slices.Map(unavailable, func(s service.Service) service.Service {
		s.State = service.StateDead
		return s
	})

	tx := qdb.MustBegin()
	defer tx.Rollback() //nolint: errcheck
	for _, update := range unavailable {
		service.Upsert(tx, update)
	}
	tx.MustCommit()
}
