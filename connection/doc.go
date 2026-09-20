// Package connection binds
// org.freedesktop.NetworkManager.Settings.Connection, one saved profile.
//
// Settings arrive as a two-level map of D-Bus variants and are decoded into
// [github.com/hnnsgstfssn/nm.ConnectionSettings], plain Go values a caller can index
// without unwrapping anything. Secrets are not part of that: NetworkManager
// keeps them out of GetSettings, so [GetSecrets] is a separate call and a
// separate decision about what to hold in memory.
package connection
