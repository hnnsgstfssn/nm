package device_test

import (
	"context"
	"testing"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm"
	"github.com/hnnsgstfssn/nm/device"
	"github.com/hnnsgstfssn/nm/internal/nmtest"
)

const (
	deviceInterface = "org.freedesktop.NetworkManager.Device"
	stateChanged    = deviceInterface + ".StateChanged"

	devicePath = dbus.ObjectPath("/org/freedesktop/NetworkManager/Devices/1")
	ip4Path    = dbus.ObjectPath("/org/freedesktop/NetworkManager/IP4Config/1")
)

func TestTypedProperties(t *testing.T) {
	bus := &nmtest.Bus{Properties: map[string]any{
		deviceInterface + ".State":           uint32(60),
		deviceInterface + ".DeviceType":      uint32(2),
		deviceInterface + ".Ip4Connectivity": uint32(2),
		deviceInterface + ".InterfaceFlags":  uint32(0x3),
	}}
	dev := nm.NewDevice(devicePath)
	ctx := t.Context()

	state, err := device.GetPropertyState(ctx, bus, dev)
	if err != nil {
		t.Fatal(err)
	}
	if state != nm.DeviceStateNeedAuth {
		t.Errorf("state = %v; want %v", state, nm.DeviceStateNeedAuth)
	}

	typ, err := device.GetPropertyDeviceType(ctx, bus, dev)
	if err != nil {
		t.Fatal(err)
	}
	if typ != nm.DeviceTypeWifi {
		t.Errorf("device type = %v; want %v", typ, nm.DeviceTypeWifi)
	}

	connectivity, err := device.GetPropertyIP4Connectivity(ctx, bus, dev)
	if err != nil {
		t.Fatal(err)
	}
	if connectivity != nm.ConnectivityPortal {
		t.Errorf("connectivity = %v; want %v", connectivity, nm.ConnectivityPortal)
	}

	flags, err := device.GetPropertyInterfaceFlags(ctx, bus, dev)
	if err != nil {
		t.Fatal(err)
	}
	if flags&nm.DeviceInterfaceFlagsUp == 0 {
		t.Errorf("interface flags = %v; want the up bit set", flags)
	}
}

// A device with no address yet reports NoObject for its IP4Config, which used
// to come back as a config object whose every read failed.
func TestOptionalConfig(t *testing.T) {
	for _, tc := range []struct {
		name     string
		value    dbus.ObjectPath
		wantNil  bool
		wantPath dbus.ObjectPath
	}{
		{name: "absent", value: nm.NoObject, wantNil: true},
		{name: "present", value: ip4Path, wantPath: ip4Path},
	} {
		t.Run(tc.name, func(t *testing.T) {
			bus := &nmtest.Bus{Properties: map[string]any{
				deviceInterface + ".Ip4Config":                            tc.value,
				"org.freedesktop.NetworkManager.IP4Config.AddressData":    []map[string]dbus.Variant{},
				"org.freedesktop.NetworkManager.IP4Config.Gateway":        "192.168.1.1",
				"org.freedesktop.NetworkManager.IP4Config.RouteData":      []map[string]dbus.Variant{},
				"org.freedesktop.NetworkManager.IP4Config.NameserverData": []map[string]dbus.Variant{},
				"org.freedesktop.NetworkManager.IP4Config.Domains":        []string{},
				"org.freedesktop.NetworkManager.IP4Config.Searches":       []string{},
				"org.freedesktop.NetworkManager.IP4Config.DnsOptions":     []string{},
				"org.freedesktop.NetworkManager.IP4Config.DnsPriority":    int32(100),
				"org.freedesktop.NetworkManager.IP4Config.WinsServerData": []string{},
			}}

			cfg, err := device.GetPropertyIP4Config(t.Context(), bus, nm.NewDevice(devicePath))
			if err != nil {
				t.Fatalf("GetPropertyIP4Config: %v", err)
			}
			if tc.wantNil {
				if cfg != nil {
					t.Fatalf("config = %+v; want nil", cfg)
				}
				return
			}
			if cfg == nil || cfg.Path != tc.wantPath {
				t.Fatalf("config = %v; want one at %q", cfg, tc.wantPath)
			}
		})
	}
}

// Device.StateChanged carries (new_state, old_state, reason). Reading the
// reason from the second argument reported the previous state as the reason
// for every transition, and a shorter body than that indexed past the end.
func TestStateChanges(t *testing.T) {
	for _, tc := range []struct {
		name   string
		signal *dbus.Signal
		want   *nm.DeviceStateChange
	}{
		{
			name:   "state and reason",
			signal: &dbus.Signal{Path: devicePath, Name: stateChanged, Body: []any{uint32(120), uint32(60), uint32(7)}},
			want: &nm.DeviceStateChange{
				Path:   devicePath,
				State:  nm.DeviceStateFailed,
				Reason: nm.DeviceStateReasonNoSecrets,
			},
		},
		{
			name:   "short body",
			signal: &dbus.Signal{Path: devicePath, Name: stateChanged, Body: []any{uint32(120), uint32(60)}},
		},
		{
			name:   "another member",
			signal: &dbus.Signal{Path: devicePath, Name: deviceInterface + ".Something", Body: []any{uint32(120), uint32(60), uint32(7)}},
		},
		{
			name:   "wrong argument types",
			signal: &dbus.Signal{Path: devicePath, Name: stateChanged, Body: []any{"failed", "config", "no secrets"}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			bus := &nmtest.Bus{}
			ctx, cancel := context.WithCancel(t.Context())
			defer cancel()

			changes, err := device.StateChanges(ctx, bus, nm.NewDevice(devicePath))
			if err != nil {
				t.Fatal(err)
			}

			// A signal the iterator must skip is followed by one it must
			// yield, so the assertion is on what comes out first rather than
			// on nothing coming out, which would need a timeout to observe.
			bus.Emit(tc.signal)
			bus.Emit(&dbus.Signal{
				Path: devicePath,
				Name: stateChanged,
				Body: []any{uint32(100), uint32(90), uint32(2)},
			})
			want := nm.DeviceStateChange{
				Path:   devicePath,
				State:  nm.DeviceStateActivated,
				Reason: nm.DeviceStateReasonNowManaged,
			}
			if tc.want != nil {
				want = *tc.want
			}

			for got := range changes {
				if got != want {
					t.Errorf("first change = %+v; want %+v", got, want)
				}
				break
			}
		})
	}
}
