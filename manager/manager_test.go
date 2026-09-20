package manager_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm"
	"github.com/hnnsgstfssn/nm/internal/nmtest"
	"github.com/hnnsgstfssn/nm/manager"
)

const (
	nmInterface  = "org.freedesktop.NetworkManager"
	addAndActive = nmInterface + ".AddAndActivateConnection"

	profilePath = dbus.ObjectPath("/org/freedesktop/NetworkManager/Settings/1")
	activePath  = dbus.ObjectPath("/org/freedesktop/NetworkManager/ActiveConnection/1")
	devicePath  = dbus.ObjectPath("/org/freedesktop/NetworkManager/Devices/1")
)

// AddAndActivateConnection sent the empty path for a nil device, which is not
// a valid object path: godbus refuses to marshal it, so the call failed before
// it reached NetworkManager instead of letting it pick a device.
func TestAddAndActivateConnectionNilDevice(t *testing.T) {
	for _, tc := range []struct {
		name string
		dev  *nm.Device
		want dbus.ObjectPath
	}{
		{"nil device", nil, nm.NoObject},
		{"named device", nm.NewDevice(devicePath), devicePath},
	} {
		t.Run(tc.name, func(t *testing.T) {
			bus := &nmtest.Bus{Replies: map[string][]any{
				addAndActive: {profilePath, activePath},
			}}

			ac, err := manager.AddAndActivateConnection(t.Context(), bus, nil, tc.dev)
			if err != nil {
				t.Fatalf("AddAndActivateConnection: %v", err)
			}
			if ac.Path != activePath {
				t.Errorf("active connection = %q; want %q", ac.Path, activePath)
			}

			args := sentArgs(t, bus, addAndActive)
			if len(args) != 3 {
				t.Fatalf("sent %d arguments; want 3", len(args))
			}
			if got := args[1]; got != tc.want {
				t.Errorf("device argument = %#v; want %q", got, tc.want)
			}
			if got := args[2]; got != nm.NoObject {
				t.Errorf("specific object argument = %#v; want %q", got, nm.NoObject)
			}
		})
	}
}

func TestTypedProperties(t *testing.T) {
	bus := &nmtest.Bus{Properties: map[string]any{
		nmInterface + ".State":        uint32(70),
		nmInterface + ".Connectivity": uint32(4),
		nmInterface + ".Metered":      uint32(3),
		nmInterface + ".Capabilities": []uint32{1},
	}}
	ctx := t.Context()

	t.Run("state", func(t *testing.T) {
		got, err := manager.GetPropertyState(ctx, bus)
		if err != nil {
			t.Fatal(err)
		}
		if got != nm.StateConnectedGlobal {
			t.Errorf("state = %v; want %v", got, nm.StateConnectedGlobal)
		}
		// The enum is what makes a log line readable, so the String method is
		// part of the reason for returning one.
		if got.String() != "StateConnectedGlobal" {
			t.Errorf("state string = %q; want %q", got, "StateConnectedGlobal")
		}
	})
	t.Run("connectivity", func(t *testing.T) {
		got, err := manager.GetPropertyConnectivity(ctx, bus)
		if err != nil {
			t.Fatal(err)
		}
		if got != nm.ConnectivityFull {
			t.Errorf("connectivity = %v; want %v", got, nm.ConnectivityFull)
		}
	})
	t.Run("metered", func(t *testing.T) {
		got, err := manager.GetPropertyMetered(ctx, bus)
		if err != nil {
			t.Fatal(err)
		}
		if got != nm.MeteredGuessYes {
			t.Errorf("metered = %v; want %v", got, nm.MeteredGuessYes)
		}
	})
	t.Run("capabilities", func(t *testing.T) {
		got, err := manager.GetPropertyCapabilities(ctx, bus)
		if err != nil {
			t.Fatal(err)
		}
		if len(got) != 1 || got[0] != nm.CapabilityTeam {
			t.Errorf("capabilities = %v; want [%v]", got, nm.CapabilityTeam)
		}
	})
}

// NetworkManager reports no primary connection as NoObject. Handing that back
// as a connection produced a handle whose every call failed.
func TestOptionalConnections(t *testing.T) {
	for _, tc := range []struct {
		name  string
		value dbus.ObjectPath
		want  *nm.ActiveConnection
	}{
		{"absent", nm.NoObject, nil},
		{"present", activePath, nm.NewActiveConnection(activePath)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			bus := &nmtest.Bus{Properties: map[string]any{
				nmInterface + ".PrimaryConnection":    tc.value,
				nmInterface + ".ActivatingConnection": tc.value,
			}}

			primary, err := manager.GetPropertyPrimaryConnection(t.Context(), bus)
			if err != nil {
				t.Fatalf("GetPropertyPrimaryConnection: %v", err)
			}
			activating, err := manager.GetPropertyActivatingConnection(t.Context(), bus)
			if err != nil {
				t.Fatalf("GetPropertyActivatingConnection: %v", err)
			}

			for name, got := range map[string]*nm.ActiveConnection{"primary": primary, "activating": activating} {
				switch {
				case tc.want == nil && got != nil:
					t.Errorf("%s connection = %q; want nil", name, got.Path)
				case tc.want != nil && (got == nil || got.Path != tc.want.Path):
					t.Errorf("%s connection = %v; want %q", name, got, tc.want.Path)
				}
			}
		})
	}
}

// Snapshot used to discard every error, so a NetworkManager that was not
// answering produced a map full of zero values indistinguishable from real
// state.
func TestSnapshotReportsFailures(t *testing.T) {
	const denied = nmtest.Error("org.freedesktop.NetworkManager.Error.NotAllowed")
	bus := &nmtest.Bus{
		Properties: map[string]any{nmInterface + ".Version": "1.42.4"},
		Errors:     map[string]error{nmInterface + ".State": denied},
	}

	values, err := manager.Snapshot(t.Context(), bus)
	if err == nil {
		t.Fatal("Snapshot succeeded with an unreadable property")
	}
	if !errors.Is(err, denied) || !strings.Contains(err.Error(), "State") {
		t.Errorf("error %q does not name the failed property", err)
	}
	if got, ok := values["Version"]; !ok || got != "1.42.4" {
		t.Errorf("Version = %v, %v; want the value that was readable", got, ok)
	}
	if _, ok := values["State"]; ok {
		t.Error("State is present in the snapshot despite failing to read")
	}
}

// sentArgs returns the arguments of the last call to member.
func sentArgs(t *testing.T, bus *nmtest.Bus, member string) []any {
	t.Helper()
	for _, c := range bus.Calls() {
		if c.Member == member {
			return c.Args
		}
	}
	t.Fatalf("no call to %s; calls were %+v", member, bus.Calls())
	return nil
}
