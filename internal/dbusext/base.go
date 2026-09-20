package dbusext

import (
	"context"
	"fmt"
	"iter"
	"strings"
	"sync"

	"github.com/godbus/dbus/v5"
)

const (
	propertiesGet = "org.freedesktop.DBus.Properties.Get"
	propertiesSet = "org.freedesktop.DBus.Properties.Set"

	// NetworkManager delivers a burst of PropertiesChanged signals around a
	// scan or an activation. godbus spawns a goroutine per signal that does
	// not fit, so the buffer is sized to absorb a burst rather than to queue
	// indefinitely.
	signalBuffer = 8
)

// Conn is the part of *dbus.Conn the bindings use.
type Conn interface {
	Object(dest string, path dbus.ObjectPath) dbus.BusObject
	Signal(ch chan<- *dbus.Signal)
	RemoveSignal(ch chan<- *dbus.Signal)
	AddMatchSignalContext(ctx context.Context, options ...dbus.MatchOption) error
	RemoveMatchSignalContext(ctx context.Context, options ...dbus.MatchOption) error
}

// Base is a NetworkManager object addressed by its path.
type Base struct {
	obj dbus.BusObject
}

// NewBase returns a Base for the object at path on the NetworkManager bus
// name.
func NewBase(conn Conn, path dbus.ObjectPath) Base {
	return Base{obj: conn.Object(NetworkManagerInterface, path)}
}

// Path reports the object path Base addresses.
func (b Base) Path() dbus.ObjectPath { return b.obj.Path() }

// Call invokes method and discards its reply.
func (b Base) Call(ctx context.Context, method string, args ...any) error {
	return b.call(ctx, method, nil, args...)
}

// CallWithReturn invokes method and stores its single return value in ret.
func (b Base) CallWithReturn(ctx context.Context, ret any, method string, args ...any) error {
	return b.call(ctx, method, []any{ret}, args...)
}

// CallWithReturn2 invokes method and stores its two return values.
func (b Base) CallWithReturn2(ctx context.Context, ret1, ret2 any, method string, args ...any) error {
	return b.call(ctx, method, []any{ret1, ret2}, args...)
}

func (b Base) call(ctx context.Context, method string, rets []any, args ...any) error {
	if err := b.obj.CallWithContext(ctx, method, 0, args...).Store(rets...); err != nil {
		return fmt.Errorf("%s on %s: %w", method, b.obj.Path(), err)
	}
	return nil
}

// SetProperty writes value to the property named iface.Member.
func (b Base) SetProperty(ctx context.Context, name string, value any) error {
	iface, prop, ok := splitProperty(name)
	if !ok {
		return fmt.Errorf("invalid property %q", name)
	}
	if err := b.obj.CallWithContext(ctx, propertiesSet, 0, iface, prop, dbus.MakeVariant(value)).Store(); err != nil {
		return fmt.Errorf("set %s on %s: %w", name, b.obj.Path(), err)
	}
	return nil
}

func (b Base) property(ctx context.Context, name string) (any, error) {
	iface, prop, ok := splitProperty(name)
	if !ok {
		return nil, fmt.Errorf("invalid property %q", name)
	}
	var v dbus.Variant
	if err := b.obj.CallWithContext(ctx, propertiesGet, 0, iface, prop).Store(&v); err != nil {
		return nil, fmt.Errorf("get %s on %s: %w", name, b.obj.Path(), err)
	}
	return v.Value(), nil
}

// splitProperty splits an interface.Member property name at its last dot.
func splitProperty(name string) (iface, prop string, ok bool) {
	i := strings.LastIndexByte(name, '.')
	if i <= 0 || i+1 == len(name) {
		return "", "", false
	}
	return name[:i], name[i+1:], true
}

// property reads a property and asserts its variant to T.
func property[T any](ctx context.Context, b Base, name string) (T, error) {
	var zero T
	v, err := b.property(ctx, name)
	if err != nil {
		return zero, err
	}
	t, ok := v.(T)
	if !ok {
		return zero, fmt.Errorf("property %s on %s: got %T, want %T", name, b.obj.Path(), v, zero)
	}
	return t, nil
}

func (b Base) GetBool(ctx context.Context, name string) (bool, error) {
	return property[bool](ctx, b, name)
}

func (b Base) GetString(ctx context.Context, name string) (string, error) {
	return property[string](ctx, b, name)
}

func (b Base) GetStrings(ctx context.Context, name string) ([]string, error) {
	return property[[]string](ctx, b, name)
}

func (b Base) GetUint8(ctx context.Context, name string) (uint8, error) {
	return property[uint8](ctx, b, name)
}

func (b Base) GetUint32(ctx context.Context, name string) (uint32, error) {
	return property[uint32](ctx, b, name)
}

func (b Base) GetUint32s(ctx context.Context, name string) ([]uint32, error) {
	return property[[]uint32](ctx, b, name)
}

func (b Base) GetInt32(ctx context.Context, name string) (int32, error) {
	return property[int32](ctx, b, name)
}

func (b Base) GetInt64(ctx context.Context, name string) (int64, error) {
	return property[int64](ctx, b, name)
}

func (b Base) GetUint64(ctx context.Context, name string) (uint64, error) {
	return property[uint64](ctx, b, name)
}

func (b Base) GetBytes(ctx context.Context, name string) ([]byte, error) {
	return property[[]byte](ctx, b, name)
}

func (b Base) GetByteSlices(ctx context.Context, name string) ([][]byte, error) {
	return property[[][]byte](ctx, b, name)
}

func (b Base) GetVariantMap(ctx context.Context, name string) (map[string]dbus.Variant, error) {
	return property[map[string]dbus.Variant](ctx, b, name)
}

func (b Base) GetVariantMaps(ctx context.Context, name string) ([]map[string]dbus.Variant, error) {
	return property[[]map[string]dbus.Variant](ctx, b, name)
}

func (b Base) GetObjects(ctx context.Context, name string) ([]dbus.ObjectPath, error) {
	return property[[]dbus.ObjectPath](ctx, b, name)
}

// GetObject reads an object path property.
func (b Base) GetObject(ctx context.Context, name string) (dbus.ObjectPath, error) {
	return property[dbus.ObjectPath](ctx, b, name)
}

// GetOptionalObject reads an object path property that NetworkManager sets to
// NoObject when the relationship does not currently exist, reporting whether
// it does. Passing NoObject on to another call yields an object whose every
// method fails, so the distinction belongs here rather than at each call
// site.
func (b Base) GetOptionalObject(ctx context.Context, name string) (dbus.ObjectPath, bool, error) {
	path, err := b.GetObject(ctx, name)
	if err != nil {
		return "", false, err
	}
	return path, path != NoObject, nil
}

// Signals subscribes to the NetworkManager signals selected by options and
// returns an iterator over them. Iteration ends when ctx is done.
//
// The subscription is registered before Signals returns, so a caller can read
// current state afterwards without racing a transition. It is removed when
// the range loop ends or when ctx is done, whichever happens first, so an
// iterator that is never ranged over still stops costing anything once its
// context does.
func Signals(ctx context.Context, conn Conn, options ...dbus.MatchOption) (iter.Seq[*dbus.Signal], error) {
	// Matching on the sender keeps any other process on the system bus from
	// emitting a signal that looks like NetworkManager's.
	match := make([]dbus.MatchOption, 0, len(options)+1)
	match = append(match, dbus.WithMatchSender(NetworkManagerInterface))
	match = append(match, options...)

	ch := make(chan *dbus.Signal, signalBuffer)
	conn.Signal(ch)
	if err := conn.AddMatchSignalContext(ctx, match...); err != nil {
		conn.RemoveSignal(ch)
		return nil, fmt.Errorf("subscribe: %w", err)
	}

	// Removal outlives ctx: tearing the rule down is itself a bus call, and
	// the usual reason to stop is that ctx is already done. The error has no
	// caller to reach and the subscription is unusable either way.
	var once sync.Once
	cleanup := func() {
		once.Do(func() {
			_ = conn.RemoveMatchSignalContext(context.WithoutCancel(ctx), match...) //nolint:errcheck
			conn.RemoveSignal(ch)
		})
	}
	stopCleanup := context.AfterFunc(ctx, cleanup)

	return func(yield func(*dbus.Signal) bool) {
		defer func() {
			stopCleanup()
			cleanup()
		}()

		for {
			select {
			case signal, ok := <-ch:
				// Every registered channel sees every signal the connection
				// receives, so callers filter on name and path.
				if !ok || !yield(signal) {
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}, nil
}
