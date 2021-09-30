package dbus

import (
	"github.com/godbus/dbus/v5"
)

// Call contains everything that a call needs
type Call interface {
	NS() string
	Path() dbus.ObjectPath
	Method() string
	Flags() dbus.Flags
}

func CallSession[T Call](args ...interface{}) (T, error) {
	return call[T](dbus.SessionBus, args)
}

func CallSystem[T Call](args ...interface{}) (T, error) {
	return call[T](dbus.SystemBus, args)
}

func call[T Call](open func() (*dbus.Conn, error), args []interface{}) (T, error) {
	conn, err := open()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Dereference since *T does not implement Call interface
	c := *(new(T))

	o := conn.Object(c.NS(), c.Path())
	return c, o.Call(c.Method(), c.Flags(), args...).Store(&c)
}
