// Package manager binds org.freedesktop.NetworkManager, the daemon object.
//
// It is the entry point: everything else in these packages needs an object
// path, and this is where the first one comes from. It also owns the calls
// that create state rather than read it, activation and checkpoints, so a
// caller that only wants to look at the current configuration can stay in
// the other packages.
package manager
