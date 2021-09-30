package systemd

import "github.com/godbus/dbus/v5"

/*

   The primary unit name as string
   The human readable description string
   The load state (i.e. whether the unit file has been loaded successfully)
   The active state (i.e. whether the unit is currently started or not)
   The sub state (a more fine-grained version of the active state that is specific to the unit type, which the active state is not)
   A unit that is being followed in its state by this unit, if there is any, otherwise the empty string.
   The unit object path
   If there is a job queued for the job unit the numeric job id, 0 otherwise
   The job type as string
   The job object path

   ssssssouso
*/
type Unit struct {
	Name        string
	Description string

	// TODO: Probably ENUM
	LoadState   string
	ActiveState string
	SubState    string

	Path     string
	UnitPath dbus.ObjectPath

	JobQueueID uint32
	JobType    string
	JobPath    dbus.ObjectPath
}

/*
	Name
	Status

    ss
*/
type UnitFile struct {
	Path  string
	State string
}
