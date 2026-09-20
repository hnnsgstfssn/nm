// Package dbusext adapts godbus to the shape the NetworkManager bindings
// need.
//
// Three things are missing from *dbus.Conn at this layer. Property reads go
// through org.freedesktop.DBus.Properties.Get by hand because the godbus
// helper takes no context, and a hung NetworkManager is exactly what an
// operator wants to abort. Every reply is a variant whose Go type has to be
// asserted, so one generic reader produces that failure message rather than
// a copy of it per type. And a subscription is a match rule plus a channel
// plus their removal, which callers get as an iterator that owns all three.
//
// [Conn] is the slice of *dbus.Conn all of that uses, so tests serve canned
// replies without a bus.
package dbusext
