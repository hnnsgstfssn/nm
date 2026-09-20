package device

import (
	"context"
	"errors"
	"fmt"
	"iter"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	deviceReapply    = dbusext.DeviceInterface + ".Reapply"
	deviceDisconnect = dbusext.DeviceInterface + ".Disconnect"
	deviceDelete     = dbusext.DeviceInterface + ".Delete"

	devicePropertyUdi                  = dbusext.DeviceInterface + ".Udi"
	devicePropertyInterface            = dbusext.DeviceInterface + ".Interface"
	devicePropertyIPInterface          = dbusext.DeviceInterface + ".IpInterface"
	devicePropertyDriver               = dbusext.DeviceInterface + ".Driver"
	devicePropertyDriverVersion        = dbusext.DeviceInterface + ".DriverVersion"
	devicePropertyFirmwareVersion      = dbusext.DeviceInterface + ".FirmwareVersion"
	devicePropertyState                = dbusext.DeviceInterface + ".State"
	devicePropertyActiveConnection     = dbusext.DeviceInterface + ".ActiveConnection"
	devicePropertyIP4Config            = dbusext.DeviceInterface + ".Ip4Config"
	devicePropertyDhcp4Config          = dbusext.DeviceInterface + ".Dhcp4Config"
	devicePropertyIP6Config            = dbusext.DeviceInterface + ".Ip6Config"
	devicePropertyDhcp6Config          = dbusext.DeviceInterface + ".Dhcp6Config"
	devicePropertyManaged              = dbusext.DeviceInterface + ".Managed"
	devicePropertyAutoconnect          = dbusext.DeviceInterface + ".Autoconnect"
	devicePropertyFirmwareMissing      = dbusext.DeviceInterface + ".FirmwareMissing"
	devicePropertyNmPluginMissing      = dbusext.DeviceInterface + ".NmPluginMissing"
	devicePropertyDeviceType           = dbusext.DeviceInterface + ".DeviceType"
	devicePropertyAvailableConnections = dbusext.DeviceInterface + ".AvailableConnections"
	devicePropertyPhysicalPortID       = dbusext.DeviceInterface + ".PhysicalPortId"
	devicePropertyMtu                  = dbusext.DeviceInterface + ".Mtu"
	devicePropertyReal                 = dbusext.DeviceInterface + ".Real"
	devicePropertyIP4Connectivity      = dbusext.DeviceInterface + ".Ip4Connectivity"
	devicePropertyIP6Connectivity      = dbusext.DeviceInterface + ".Ip6Connectivity"
	devicePropertyInterfaceFlags       = dbusext.DeviceInterface + ".InterfaceFlags"
	devicePropertyHwAddress            = dbusext.DeviceInterface + ".HwAddress"
	devicePropertyPorts                = dbusext.DeviceInterface + ".Ports"

	deviceSignalStateChanged = "StateChanged"
)

// Reapply attempts to update the configuration of a device without
// deactivating it.
func Reapply(ctx context.Context, conn nm.Conn, dev *nm.Device, settings nm.ConnectionSettings, versionID uint64, flags uint32) error {
	b := dbusext.NewBase(conn, dev.Path)
	return b.Call(ctx, deviceReapply, settings, versionID, flags)
}

// Disconnect disconnects the device and prevents automatic reactivation.
func Disconnect(ctx context.Context, conn nm.Conn, dev *nm.Device) error {
	b := dbusext.NewBase(conn, dev.Path)
	return b.Call(ctx, deviceDisconnect)
}

// Delete removes a software device from NetworkManager.
func Delete(ctx context.Context, conn nm.Conn, dev *nm.Device) error {
	b := dbusext.NewBase(conn, dev.Path)
	return b.Call(ctx, deviceDelete)
}

func GetPropertyUdi(ctx context.Context, conn nm.Conn, dev *nm.Device) (string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetString(ctx, devicePropertyUdi)
}

func GetPropertyInterface(ctx context.Context, conn nm.Conn, dev *nm.Device) (string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetString(ctx, devicePropertyInterface)
}

func GetPropertyIPInterface(ctx context.Context, conn nm.Conn, dev *nm.Device) (string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetString(ctx, devicePropertyIPInterface)
}

func GetPropertyDriver(ctx context.Context, conn nm.Conn, dev *nm.Device) (string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetString(ctx, devicePropertyDriver)
}

func GetPropertyDriverVersion(ctx context.Context, conn nm.Conn, dev *nm.Device) (string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetString(ctx, devicePropertyDriverVersion)
}

func GetPropertyFirmwareVersion(ctx context.Context, conn nm.Conn, dev *nm.Device) (string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetString(ctx, devicePropertyFirmwareVersion)
}

func GetPropertyState(ctx context.Context, conn nm.Conn, dev *nm.Device) (nm.DeviceState, error) {
	b := dbusext.NewBase(conn, dev.Path)
	state, err := b.GetUint32(ctx, devicePropertyState)
	return nm.DeviceState(state), err
}

func GetPropertyActiveConnection(ctx context.Context, conn nm.Conn, dev *nm.Device) (*nm.ActiveConnection, error) {
	b := dbusext.NewBase(conn, dev.Path)
	path, ok, err := b.GetOptionalObject(ctx, devicePropertyActiveConnection)
	if err != nil || !ok {
		return nil, err
	}
	return nm.NewActiveConnection(path), nil
}

func GetPropertyIP4Config(ctx context.Context, conn nm.Conn, dev *nm.Device) (*nm.IP4Config, error) {
	b := dbusext.NewBase(conn, dev.Path)
	path, ok, err := b.GetOptionalObject(ctx, devicePropertyIP4Config)
	if err != nil || !ok {
		return nil, err
	}
	return nm.NewIP4Config(ctx, conn, path)
}

func GetPropertyDHCP4Config(ctx context.Context, conn nm.Conn, dev *nm.Device) (*nm.DHCP4Config, error) {
	b := dbusext.NewBase(conn, dev.Path)
	path, ok, err := b.GetOptionalObject(ctx, devicePropertyDhcp4Config)
	if err != nil || !ok {
		return nil, err
	}
	return nm.NewDHCP4Config(ctx, conn, path)
}

func GetPropertyIP6Config(ctx context.Context, conn nm.Conn, dev *nm.Device) (*nm.IP6Config, error) {
	b := dbusext.NewBase(conn, dev.Path)
	path, ok, err := b.GetOptionalObject(ctx, devicePropertyIP6Config)
	if err != nil || !ok {
		return nil, err
	}
	return nm.NewIP6Config(ctx, conn, path)
}

func GetPropertyDHCP6Config(ctx context.Context, conn nm.Conn, dev *nm.Device) (*nm.DHCP6Config, error) {
	b := dbusext.NewBase(conn, dev.Path)
	path, ok, err := b.GetOptionalObject(ctx, devicePropertyDhcp6Config)
	if err != nil || !ok {
		return nil, err
	}
	return nm.NewDHCP6Config(ctx, conn, path)
}

func GetPropertyManaged(ctx context.Context, conn nm.Conn, dev *nm.Device) (bool, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetBool(ctx, devicePropertyManaged)
}

func SetPropertyManaged(ctx context.Context, conn nm.Conn, dev *nm.Device, managed bool) error {
	b := dbusext.NewBase(conn, dev.Path)
	return b.SetProperty(ctx, devicePropertyManaged, managed)
}

func GetPropertyAutoConnect(ctx context.Context, conn nm.Conn, dev *nm.Device) (bool, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetBool(ctx, devicePropertyAutoconnect)
}

func SetPropertyAutoConnect(ctx context.Context, conn nm.Conn, dev *nm.Device, autoconnect bool) error {
	b := dbusext.NewBase(conn, dev.Path)
	return b.SetProperty(ctx, devicePropertyAutoconnect, autoconnect)
}

func GetPropertyFirmwareMissing(ctx context.Context, conn nm.Conn, dev *nm.Device) (bool, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetBool(ctx, devicePropertyFirmwareMissing)
}

func GetPropertyNmPluginMissing(ctx context.Context, conn nm.Conn, dev *nm.Device) (bool, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetBool(ctx, devicePropertyNmPluginMissing)
}

func GetPropertyDeviceType(ctx context.Context, conn nm.Conn, dev *nm.Device) (nm.DeviceType, error) {
	b := dbusext.NewBase(conn, dev.Path)
	typ, err := b.GetUint32(ctx, devicePropertyDeviceType)
	return nm.DeviceType(typ), err
}

func GetPropertyAvailableConnections(ctx context.Context, conn nm.Conn, dev *nm.Device) ([]*nm.Connection, error) {
	b := dbusext.NewBase(conn, dev.Path)
	paths, err := b.GetObjects(ctx, devicePropertyAvailableConnections)
	if err != nil {
		return nil, err
	}
	conns := make([]*nm.Connection, len(paths))
	for i, path := range paths {
		conns[i] = nm.NewConnection(path)
	}
	return conns, nil
}

func GetPropertyPhysicalPortID(ctx context.Context, conn nm.Conn, dev *nm.Device) (string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetString(ctx, devicePropertyPhysicalPortID)
}

func GetPropertyMtu(ctx context.Context, conn nm.Conn, dev *nm.Device) (uint32, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetUint32(ctx, devicePropertyMtu)
}

func GetPropertyReal(ctx context.Context, conn nm.Conn, dev *nm.Device) (bool, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetBool(ctx, devicePropertyReal)
}

func GetPropertyIP4Connectivity(ctx context.Context, conn nm.Conn, dev *nm.Device) (nm.Connectivity, error) {
	b := dbusext.NewBase(conn, dev.Path)
	connectivity, err := b.GetUint32(ctx, devicePropertyIP4Connectivity)
	return nm.Connectivity(connectivity), err
}

func GetPropertyIP6Connectivity(ctx context.Context, conn nm.Conn, dev *nm.Device) (nm.Connectivity, error) {
	b := dbusext.NewBase(conn, dev.Path)
	connectivity, err := b.GetUint32(ctx, devicePropertyIP6Connectivity)
	return nm.Connectivity(connectivity), err
}

func GetPropertyInterfaceFlags(ctx context.Context, conn nm.Conn, dev *nm.Device) (nm.DeviceInterfaceFlags, error) {
	b := dbusext.NewBase(conn, dev.Path)
	flags, err := b.GetUint32(ctx, devicePropertyInterfaceFlags)
	return nm.DeviceInterfaceFlags(flags), err
}

func GetPropertyHwAddress(ctx context.Context, conn nm.Conn, dev *nm.Device) (string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetString(ctx, devicePropertyHwAddress)
}

func GetPropertyPorts(ctx context.Context, conn nm.Conn, dev *nm.Device) ([]*nm.Device, error) {
	b := dbusext.NewBase(conn, dev.Path)
	paths, err := b.GetObjects(ctx, devicePropertyPorts)
	if err != nil {
		return nil, err
	}
	devices := make([]*nm.Device, len(paths))
	for i, path := range paths {
		devices[i] = nm.NewDevice(path)
	}
	return devices, nil
}

// StateChanges returns an iterator over dev's state transitions. Iteration
// ends when ctx is done.
//
// The subscription is in place before this returns, so a caller can read the
// current state afterwards without a transition slipping through the gap.
// The match rule is removed when the range loop ends or when ctx is done,
// whichever comes first.
func StateChanges(ctx context.Context, conn nm.Conn, dev *nm.Device) (iter.Seq[nm.DeviceStateChange], error) {
	signals, err := dbusext.Signals(ctx, conn,
		dbus.WithMatchInterface(dbusext.DeviceInterface),
		dbus.WithMatchMember(deviceSignalStateChanged),
		dbus.WithMatchObjectPath(dev.Path),
	)
	if err != nil {
		return nil, err
	}

	return func(yield func(nm.DeviceStateChange) bool) {
		for signal := range signals {
			change, ok := stateChangeFrom(signal)
			if !ok {
				continue
			}
			if !yield(change) {
				return
			}
		}
	}, nil
}

// stateChangeFrom reads a StateChanged signal, reporting whether it was one.
// The body is (new_state, old_state, reason), and its length is checked
// because any process on the bus can emit a signal with this name.
func stateChangeFrom(signal *dbus.Signal) (nm.DeviceStateChange, bool) {
	if signal.Name != dbusext.DeviceInterface+"."+deviceSignalStateChanged || len(signal.Body) < 3 {
		return nm.DeviceStateChange{}, false
	}
	state, ok1 := signal.Body[0].(uint32)
	reason, ok2 := signal.Body[2].(uint32)
	if !ok1 || !ok2 {
		return nm.DeviceStateChange{}, false
	}
	return nm.DeviceStateChange{
		Path:   signal.Path,
		State:  nm.DeviceState(state),
		Reason: nm.DeviceStateReason(reason),
	}, true
}

// snapshot collects property reads, keeping the values that came back and the
// errors from the reads that did not.
type snapshot struct {
	values map[string]any
	errs   []error
}

// read adds one property to s. A failed read leaves its key absent rather
// than storing a zero value that reads as real state, and stops once ctx is
// done so that a cancelled snapshot does not make more doomed calls.
func read[T any](ctx context.Context, conn nm.Conn, dev *nm.Device, s *snapshot, key string, get func(context.Context, nm.Conn, *nm.Device) (T, error)) {
	if ctx.Err() != nil {
		return
	}
	v, err := get(ctx, conn, dev)
	if err != nil {
		s.errs = append(s.errs, fmt.Errorf("%s: %w", key, err))
		return
	}
	s.values[key] = v
}

// Snapshot reads the device properties worth having when diagnosing a link
// that will not come up. A property that cannot be read is absent from the
// map and its error is joined into the returned error.
func Snapshot(ctx context.Context, conn nm.Conn, dev *nm.Device) (map[string]any, error) {
	s := &snapshot{values: map[string]any{}}

	read(ctx, conn, dev, s, "Interface", GetPropertyInterface)
	read(ctx, conn, dev, s, "IPInterface", GetPropertyIPInterface)
	read(ctx, conn, dev, s, "State", GetPropertyState)
	read(ctx, conn, dev, s, "DeviceType", GetPropertyDeviceType)
	read(ctx, conn, dev, s, "IP4Config", GetPropertyIP4Config)
	read(ctx, conn, dev, s, "DHCP4Config", GetPropertyDHCP4Config)
	read(ctx, conn, dev, s, "AvailableConnections", GetPropertyAvailableConnections)

	return s.values, errors.Join(s.errs...)
}
