package systemd

import (
	"github.com/godbus/dbus/v5"
)

// ListUnitFiles call for dbus
type ListUnitFiles []UnitFile

func (ListUnitFiles) NS() string {
	return "org.freedesktop.systemd1"
}

func (ListUnitFiles) Path() dbus.ObjectPath {
	return "/org/freedesktop/systemd1"
}

func (ListUnitFiles) Method() string {
	return "org.freedesktop.systemd1.Manager.ListUnitFiles"
}

func (ListUnitFiles) Flags() dbus.Flags {
	return 0
}

// ListUnit call for dbus
type ListUnits []Unit

func (ListUnits) NS() string {
	return "org.freedesktop.systemd1"
}

func (ListUnits) Path() dbus.ObjectPath {
	return "/org/freedesktop/systemd1"
}

func (ListUnits) Method() string {
	return "org.freedesktop.systemd1.Manager.ListUnits"
}

func (ListUnits) Flags() dbus.Flags {
	return 0
}

// LoadUnit() is similar to GetUnit() but will load the unit from disk if possible.
type LoadUnit dbus.ObjectPath

func (LoadUnit) NS() string {
	return "org.freedesktop.systemd1"
}

func (LoadUnit) Path() dbus.ObjectPath {
	return "/org/freedesktop/systemd1"
}

func (LoadUnit) Method() string {
	return "org.freedesktop.systemd1.Manager.LoadUnit"
}

func (LoadUnit) Flags() dbus.Flags {
	return 0
}
