// Package device binds org.freedesktop.NetworkManager.Device.
//
// A NetworkManager device object carries one generic interface plus one for
// the kind of link it is: Wired, Wireless, Bridge, Generic, IPTunnel and
// Statistics all live on the same object path. The prefix on a function
// name says which interface it reads, because the argument is the same
// [github.com/hnnsgstfssn/nm.Device] either way, and asking a wired device for its access
// points is a runtime error rather than a compile-time one.
package device
