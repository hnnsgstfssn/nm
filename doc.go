// Package nm binds NetworkManager's D-Bus API.
//
// NetworkManager exposes a dozen interfaces over object paths, and an object
// path is not a handle: it is a name that may already be gone, and reading a
// single property from it is a round trip to the daemon. The bindings are
// shaped around that. This package holds the value types every interface
// deals in, and one subpackage per interface holds its calls:
//
//	nm/manager      org.freedesktop.NetworkManager
//	nm/device       .Device and its per-link interfaces
//	nm/active       .Connection.Active
//	nm/settings     .Settings
//	nm/connection   .Settings.Connection
//
// Two kinds of type live here. [Device], [Connection] and [ActiveConnection]
// are handles: a path and nothing else, so holding one costs nothing and
// says nothing about whether the object still exists. [AccessPoint],
// [IP4Config] and the rest are snapshots, read in full at construction, so a
// caller that wants ten properties writes one error check instead of ten and
// the values cannot move under it afterwards. That is not an atomic view:
// NetworkManager offers no way to ask for one, so the ten reads are ten
// round trips and the device can change between them. Nothing here caches,
// and nothing retains the connection.
//
// Every call takes a context. A D-Bus call is I/O, the daemon can wedge, and
// a device onboarding over WiFi is exactly where someone wants to give up
// and retry. Signal subscriptions are iterators rather than channels with a
// goroutine behind them: the range loop is the only reader, and the match
// rule goes away when that loop ends or when the context does.
//
// The bus is [Conn] rather than *dbus.Conn so that the bindings can be
// tested against canned replies; *dbus.Conn satisfies it.
package nm
