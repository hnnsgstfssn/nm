// Package active binds org.freedesktop.NetworkManager.Connection.Active.
//
// An active connection is a profile that is currently applied to a device,
// which makes it the place to watch an activation succeed or fail:
// [StateChanges] is how a caller learns a wrong password was rejected
// rather than waiting for a timeout. The object disappears when the
// connection goes down, so its accessors are the ones most likely to fail
// on a path that was valid a moment ago.
package active
