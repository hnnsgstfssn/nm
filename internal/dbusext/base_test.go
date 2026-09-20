package dbusext_test

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
	"github.com/hnnsgstfssn/nm/internal/nmtest"
)

const (
	devicePath  = dbus.ObjectPath("/org/freedesktop/NetworkManager/Devices/1")
	stateProp   = dbusext.DeviceInterface + ".State"
	ip4Prop     = dbusext.DeviceInterface + ".Ip4Config"
	stateSignal = dbusext.DeviceInterface + ".StateChanged"
)

func TestPropertyTypes(t *testing.T) {
	bus := &nmtest.Bus{Properties: map[string]any{
		stateProp:                                 uint32(100),
		dbusext.DeviceInterface + ".Interface":    "wlan0",
		dbusext.DeviceInterface + ".Managed":      true,
		dbusext.DeviceInterface + ".Ports":        []dbus.ObjectPath{devicePath},
		dbusext.AccessPointInterface + ".Ssid":    []byte("kitchen"),
		dbusext.DeviceWirelessInterface + ".Last": int64(-1),
	}}
	b := dbusext.NewBase(bus, devicePath)
	ctx := t.Context()

	t.Run("uint32", func(t *testing.T) {
		got, err := b.GetUint32(ctx, stateProp)
		if err != nil || got != 100 {
			t.Fatalf("GetUint32 = %v, %v; want 100, nil", got, err)
		}
	})
	t.Run("string", func(t *testing.T) {
		got, err := b.GetString(ctx, dbusext.DeviceInterface+".Interface")
		if err != nil || got != "wlan0" {
			t.Fatalf("GetString = %q, %v; want \"wlan0\", nil", got, err)
		}
	})
	t.Run("bool", func(t *testing.T) {
		got, err := b.GetBool(ctx, dbusext.DeviceInterface+".Managed")
		if err != nil || !got {
			t.Fatalf("GetBool = %v, %v; want true, nil", got, err)
		}
	})
	t.Run("objects", func(t *testing.T) {
		got, err := b.GetObjects(ctx, dbusext.DeviceInterface+".Ports")
		if err != nil || len(got) != 1 || got[0] != devicePath {
			t.Fatalf("GetObjects = %v, %v; want [%s], nil", got, err, devicePath)
		}
	})
	t.Run("bytes", func(t *testing.T) {
		got, err := b.GetBytes(ctx, dbusext.AccessPointInterface+".Ssid")
		if err != nil || string(got) != "kitchen" {
			t.Fatalf("GetBytes = %q, %v; want \"kitchen\", nil", got, err)
		}
	})
	t.Run("int64", func(t *testing.T) {
		got, err := b.GetInt64(ctx, dbusext.DeviceWirelessInterface+".Last")
		if err != nil || got != -1 {
			t.Fatalf("GetInt64 = %v, %v; want -1, nil", got, err)
		}
	})
}

// A property that comes back as the wrong type is a binding bug, and the
// message has to say which property and what arrived: every accessor shares
// one implementation, so without the name the report is useless.
func TestPropertyWrongType(t *testing.T) {
	bus := &nmtest.Bus{Properties: map[string]any{stateProp: "one hundred"}}
	b := dbusext.NewBase(bus, devicePath)

	_, err := b.GetUint32(t.Context(), stateProp)
	if err == nil {
		t.Fatal("GetUint32 on a string property succeeded")
	}
	for _, want := range []string{stateProp, string(devicePath), "string", "uint32"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q does not mention %q", err, want)
		}
	}
}

// A failed call has to name the member and the object, because the alternative
// is a bare D-Bus error after ten sequential property reads with nothing to
// say which one failed.
func TestCallErrorContext(t *testing.T) {
	const boom = nmtest.Error("org.freedesktop.NetworkManager.Error.NotAllowed")
	bus := &nmtest.Bus{Errors: map[string]error{
		dbusext.DeviceInterface + ".Disconnect": boom,
		stateProp:                               boom,
	}}
	b := dbusext.NewBase(bus, devicePath)

	callErr := b.Call(t.Context(), dbusext.DeviceInterface+".Disconnect")
	_, propErr := b.GetUint32(t.Context(), stateProp)

	for _, tc := range []struct {
		name string
		err  error
		want string
	}{
		{"method", callErr, "Disconnect"},
		{"property", propErr, "State"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if !errors.Is(tc.err, boom) {
				t.Fatalf("error %v does not wrap the bus error", tc.err)
			}
			if !strings.Contains(tc.err.Error(), tc.want) || !strings.Contains(tc.err.Error(), string(devicePath)) {
				t.Errorf("error %q does not name %q and %q", tc.err, tc.want, devicePath)
			}
		})
	}
}

func TestOptionalObject(t *testing.T) {
	for _, tc := range []struct {
		name     string
		value    dbus.ObjectPath
		wantOK   bool
		wantPath dbus.ObjectPath
	}{
		{"present", devicePath, true, devicePath},
		{"absent", dbusext.NoObject, false, dbusext.NoObject},
	} {
		t.Run(tc.name, func(t *testing.T) {
			bus := &nmtest.Bus{Properties: map[string]any{ip4Prop: tc.value}}
			b := dbusext.NewBase(bus, devicePath)

			path, ok, err := b.GetOptionalObject(t.Context(), ip4Prop)
			if err != nil {
				t.Fatalf("GetOptionalObject: %v", err)
			}
			if ok != tc.wantOK || path != tc.wantPath {
				t.Errorf("GetOptionalObject = %q, %v; want %q, %v", path, ok, tc.wantPath, tc.wantOK)
			}
		})
	}
}

// A context that is already done must not reach the bus at all: the point of
// threading one through is that a wedged NetworkManager can be abandoned.
func TestContextHonored(t *testing.T) {
	bus := &nmtest.Bus{Properties: map[string]any{stateProp: uint32(100)}}
	b := dbusext.NewBase(bus, devicePath)

	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	if _, err := b.GetUint32(ctx, stateProp); !errors.Is(err, context.Canceled) {
		t.Fatalf("GetUint32 on a cancelled context = %v; want context.Canceled", err)
	}
}

func TestSignals(t *testing.T) {
	t.Run("yields and removes the rule", func(t *testing.T) {
		bus := &nmtest.Bus{}
		ctx, cancel := context.WithCancel(t.Context())
		defer cancel()

		signals, err := dbusext.Signals(ctx, bus, dbus.WithMatchObjectPath(devicePath))
		if err != nil {
			t.Fatal(err)
		}
		// The rule is in place before Signals returns, which is what lets a
		// caller read current state without racing a transition.
		if added, _ := bus.Subscriptions(); added != 1 {
			t.Fatalf("match rules added before ranging = %d; want 1", added)
		}

		bus.Emit(&dbus.Signal{Path: devicePath, Name: stateSignal, Body: []any{uint32(100)}})

		for signal := range signals {
			if signal.Name != stateSignal {
				t.Errorf("signal name = %q; want %q", signal.Name, stateSignal)
			}
			break
		}

		if _, removed := bus.Subscriptions(); removed != 1 {
			t.Errorf("match rules removed after the loop = %d; want 1", removed)
		}
	})

	t.Run("ends when the context does", func(t *testing.T) {
		bus := &nmtest.Bus{}
		ctx, cancel := context.WithCancel(t.Context())

		signals, err := dbusext.Signals(ctx, bus, dbus.WithMatchObjectPath(devicePath))
		if err != nil {
			t.Fatal(err)
		}

		cancel()
		for range signals {
			t.Fatal("iterated a signal after cancellation")
		}
		if _, removed := bus.Subscriptions(); removed != 1 {
			t.Errorf("match rules removed = %d; want 1", removed)
		}
	})

	// An iterator the caller never ranges over would otherwise hold its match
	// rule for the life of the connection.
	t.Run("abandoned iterator is cleaned up", func(t *testing.T) {
		bus := &nmtest.Bus{}
		removed := bus.Removed()
		ctx, cancel := context.WithCancel(t.Context())

		if _, err := dbusext.Signals(ctx, bus, dbus.WithMatchObjectPath(devicePath)); err != nil {
			t.Fatal(err)
		}
		cancel()

		// The cleanup runs in the goroutine context.AfterFunc starts, so wait
		// for the removal itself. The deadline is here to fail the test.
		select {
		case <-removed:
		case <-time.After(5 * time.Second):
			t.Fatal("cancelling the context did not remove the match rule")
		}
	})

	t.Run("failed subscription is not registered", func(t *testing.T) {
		bus := &nmtest.Bus{MatchErr: nmtest.Error("org.freedesktop.DBus.Error.AccessDenied")}

		if _, err := dbusext.Signals(t.Context(), bus, dbus.WithMatchObjectPath(devicePath)); err == nil {
			t.Fatal("Signals succeeded with a refused match rule")
		}
		bus.Emit(&dbus.Signal{Path: devicePath, Name: stateSignal})
	})
}
