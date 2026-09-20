// Package nmtest is a NetworkManager that answers from a map.
//
// The bindings reach the bus through one small interface, which is what makes
// this possible: [Bus] serves canned property values and method replies, and
// records what was sent, so the parts worth testing - which member a call
// names, which arguments it carries, how a variant is decoded, what happens
// to a subscription - can be tested without a daemon.
//
// Replies travel as *dbus.Call values, so the real godbus Store path converts
// them and a test that decodes wrongly fails here too.
package nmtest

import (
	"context"
	"sync"

	"github.com/godbus/dbus/v5"
)

// Bus is a fake system bus. The zero value answers every property with an
// error; fill in the maps for the ones a test needs.
type Bus struct {
	// Properties answers Get, keyed by the full "interface.Member" name. A
	// value is returned as the variant NetworkManager would send.
	Properties map[string]any
	// Replies answers method calls, keyed by the full member name, with the
	// reply body. A member that is absent replies with an empty body, which
	// is what a void method does.
	Replies map[string][]any
	// Errors fails a property or method by full name, taking precedence over
	// Properties and Replies.
	Errors map[string]error
	// MatchErr fails subscription setup.
	MatchErr error

	// Field order below is betteralign's, not a grouping.
	removed  chan struct{}
	calls    []Call
	signals  []chan<- *dbus.Signal
	mu       sync.Mutex
	matches  int
	removals int
}

// Removed returns a channel that receives once per match rule removal, so a
// test can wait for a subscription to be torn down rather than for a
// duration. Removal can happen in a goroutine the test did not start:
// [dbusext.Signals] cleans up from a context.AfterFunc.
func (b *Bus) Removed() <-chan struct{} {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.removed == nil {
		b.removed = make(chan struct{}, 4)
	}
	return b.removed
}

// Call is one recorded method invocation. Property reads and writes are
// recorded as calls to org.freedesktop.DBus.Properties.Get and .Set.
type Call struct {
	Path   dbus.ObjectPath
	Member string
	Args   []any
}

// Calls returns the calls made so far, oldest first.
func (b *Bus) Calls() []Call {
	b.mu.Lock()
	defer b.mu.Unlock()
	return append([]Call(nil), b.calls...)
}

// Subscriptions reports the number of match rules added and removed.
func (b *Bus) Subscriptions() (added, removed int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.matches, b.removals
}

// Emit delivers a signal to every registered channel, as the bus does.
func (b *Bus) Emit(signal *dbus.Signal) {
	b.mu.Lock()
	channels := append([]chan<- *dbus.Signal(nil), b.signals...)
	b.mu.Unlock()

	for _, ch := range channels {
		ch <- signal
	}
}

func (b *Bus) Object(_ string, path dbus.ObjectPath) dbus.BusObject {
	return &object{bus: b, path: path}
}

func (b *Bus) Signal(ch chan<- *dbus.Signal) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.signals = append(b.signals, ch)
}

func (b *Bus) RemoveSignal(ch chan<- *dbus.Signal) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for i, s := range b.signals {
		if s == ch {
			b.signals = append(b.signals[:i], b.signals[i+1:]...)
			return
		}
	}
}

func (b *Bus) AddMatchSignalContext(_ context.Context, _ ...dbus.MatchOption) error {
	if b.MatchErr != nil {
		return b.MatchErr
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	b.matches++
	return nil
}

func (b *Bus) RemoveMatchSignalContext(_ context.Context, _ ...dbus.MatchOption) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.removals++
	// Non-blocking: a test that does not watch removals must not wedge the
	// caller's cleanup.
	select {
	case b.removed <- struct{}{}:
	default:
	}
	return nil
}

// object is one path on the fake bus.
type object struct {
	bus  *Bus
	path dbus.ObjectPath
}

func (o *object) CallWithContext(ctx context.Context, member string, _ dbus.Flags, args ...any) *dbus.Call {
	o.bus.mu.Lock()
	o.bus.calls = append(o.bus.calls, Call{Path: o.path, Member: member, Args: args})
	o.bus.mu.Unlock()

	if err := ctx.Err(); err != nil {
		return &dbus.Call{Err: err}
	}
	if member == propertiesGet {
		return o.property(args)
	}
	if err := o.bus.Errors[member]; err != nil {
		return &dbus.Call{Err: err}
	}
	return &dbus.Call{Body: o.bus.Replies[member]}
}

// property answers org.freedesktop.DBus.Properties.Get, whose arguments are
// the interface name and the member name.
func (o *object) property(args []any) *dbus.Call {
	if len(args) != 2 {
		return &dbus.Call{Err: errUnknown}
	}
	iface, isIface := args[0].(string)
	member, isMember := args[1].(string)
	if !isIface || !isMember {
		return &dbus.Call{Err: errUnknown}
	}
	name := iface + "." + member

	if err := o.bus.Errors[name]; err != nil {
		return &dbus.Call{Err: err}
	}
	v, ok := o.bus.Properties[name]
	if !ok {
		return &dbus.Call{Err: errUnknown}
	}
	return &dbus.Call{Body: []any{dbus.MakeVariant(v)}}
}

func (o *object) Path() dbus.ObjectPath { return o.path }

func (o *object) Call(member string, flags dbus.Flags, args ...any) *dbus.Call {
	return o.CallWithContext(context.Background(), member, flags, args...)
}

// The rest of dbus.BusObject is not reachable through the bindings: they call
// CallWithContext and nothing else.

func (o *object) Go(string, dbus.Flags, chan *dbus.Call, ...any) *dbus.Call {
	return &dbus.Call{Err: errUnsupported}
}

func (o *object) GoWithContext(context.Context, string, dbus.Flags, chan *dbus.Call, ...any) *dbus.Call {
	return &dbus.Call{Err: errUnsupported}
}

func (o *object) AddMatchSignal(string, string, ...dbus.MatchOption) *dbus.Call {
	return &dbus.Call{Err: errUnsupported}
}

func (o *object) RemoveMatchSignal(string, string, ...dbus.MatchOption) *dbus.Call {
	return &dbus.Call{Err: errUnsupported}
}

func (o *object) GetProperty(string) (dbus.Variant, error) {
	return dbus.Variant{}, errUnsupported
}

func (o *object) StoreProperty(string, any) error { return errUnsupported }

func (o *object) SetProperty(string, any) error { return errUnsupported }

func (o *object) Destination() string { return networkManager }

const (
	propertiesGet  = "org.freedesktop.DBus.Properties.Get"
	networkManager = "org.freedesktop.NetworkManager"
)

// Error is a bus-side failure. It matches the shape of a D-Bus error closely
// enough for callers that only report it.
type Error string

func (e Error) Error() string { return string(e) }

const (
	errUnknown     Error = "org.freedesktop.DBus.Error.UnknownProperty"
	errUnsupported Error = "nmtest: method not served"
)
