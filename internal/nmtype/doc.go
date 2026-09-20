// Package nmtype holds NetworkManager's enums and their String methods.
//
// It exists to keep seventeen files of stringer output out of the package
// callers actually use: [github.com/hnnsgstfssn/nm] aliases every type and constant here,
// so nm.DeviceState is this DeviceState and nothing outside needs to name
// this package.
//
// Adding a value means adding the constant and re-running the go:generate
// line above its type. The generated tables are indexed by value, so they
// are wrong rather than stale if that is skipped, and the compiler check
// stringer emits is what catches it.
package nmtype
