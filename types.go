package nm

import (
	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

// Conn is the bus the bindings talk to. *dbus.Conn satisfies it.
type Conn = dbusext.Conn

// NoObject is the object path NetworkManager reports for a relationship that
// does not currently exist. Accessors that can return it do so as a nil
// result rather than as a handle that fails on use.
const NoObject = dbusext.NoObject

// Device is a handle to a realized NetworkManager device.
type Device struct {
	Path dbus.ObjectPath
}

// NewDevice returns a Device handle for the given D-Bus object path.
func NewDevice(path dbus.ObjectPath) *Device {
	return &Device{Path: path}
}

// Connection is a handle to a NetworkManager connection profile.
type Connection struct {
	Path dbus.ObjectPath
}

// NewConnection returns a Connection handle for the given D-Bus object path.
func NewConnection(path dbus.ObjectPath) *Connection {
	return &Connection{Path: path}
}

// ActiveConnection is a handle to an active network connection.
type ActiveConnection struct {
	Path dbus.ObjectPath
}

// NewActiveConnection returns an ActiveConnection handle for the given D-Bus
// object path.
func NewActiveConnection(path dbus.ObjectPath) *ActiveConnection {
	return &ActiveConnection{Path: path}
}

// ConnectionSettings holds a two-level map of NM connection settings.
type ConnectionSettings map[string]map[string]any

// DeviceStateChange is one device state transition. NetworkManager also
// reports the state being left, which is not kept here because nothing has
// needed it.
type DeviceStateChange struct {
	Path   dbus.ObjectPath
	State  DeviceState
	Reason DeviceStateReason
}

// StateChange carries information about an active connection state
// transition.
type StateChange struct {
	Path   dbus.ObjectPath
	State  ActiveConnectionState
	Reason ActiveConnectionStateReason
}
