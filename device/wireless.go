package device

import (
	"context"
	"iter"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	deviceWirelessGetAccessPoints    = dbusext.DeviceWirelessInterface + ".GetAccessPoints"
	deviceWirelessGetAllAccessPoints = dbusext.DeviceWirelessInterface + ".GetAllAccessPoints"
	deviceWirelessRequestScan        = dbusext.DeviceWirelessInterface + ".RequestScan"

	deviceWirelessPropertyPermHwAddress        = dbusext.DeviceWirelessInterface + ".PermHwAddress"
	deviceWirelessPropertyMode                 = dbusext.DeviceWirelessInterface + ".Mode"
	deviceWirelessPropertyBitrate              = dbusext.DeviceWirelessInterface + ".Bitrate"
	deviceWirelessPropertyAccessPoints         = dbusext.DeviceWirelessInterface + ".AccessPoints"
	deviceWirelessPropertyActiveAccessPoint    = dbusext.DeviceWirelessInterface + ".ActiveAccessPoint"
	deviceWirelessPropertyWirelessCapabilities = dbusext.DeviceWirelessInterface + ".WirelessCapabilities"
	deviceWirelessPropertyLastScan             = dbusext.DeviceWirelessInterface + ".LastScan"

	propertiesInterface            = "org.freedesktop.DBus.Properties"
	propertiesSignalChanged        = "PropertiesChanged"
	deviceWirelessLastScanProperty = "LastScan"
)

// WirelessGetAccessPoints returns visible access points (excluding hidden SSIDs).
func WirelessGetAccessPoints(ctx context.Context, conn nm.Conn, dev *nm.Device) ([]*nm.AccessPoint, error) {
	b := dbusext.NewBase(conn, dev.Path)
	var paths []dbus.ObjectPath
	if err := b.CallWithReturn(ctx, &paths, deviceWirelessGetAccessPoints); err != nil {
		return nil, err
	}
	aps := make([]*nm.AccessPoint, len(paths))
	var err error
	for i, path := range paths {
		aps[i], err = nm.NewAccessPoint(ctx, conn, path)
		if err != nil {
			return aps, err
		}
	}
	return aps, nil
}

// WirelessGetAllAccessPoints returns all visible access points, including hidden.
func WirelessGetAllAccessPoints(ctx context.Context, conn nm.Conn, dev *nm.Device) ([]*nm.AccessPoint, error) {
	b := dbusext.NewBase(conn, dev.Path)
	var paths []dbus.ObjectPath
	if err := b.CallWithReturn(ctx, &paths, deviceWirelessGetAllAccessPoints); err != nil {
		return nil, err
	}
	aps := make([]*nm.AccessPoint, len(paths))
	var err error
	for i, path := range paths {
		aps[i], err = nm.NewAccessPoint(ctx, conn, path)
		if err != nil {
			return aps, err
		}
	}
	return aps, nil
}

// WirelessRequestScan asks NetworkManager to scan. It returns as soon as the
// request is accepted; use [LastScans] to learn when one finishes.
//
// NetworkManager refuses the request while a scan is already running and
// rate limits repeats, so an error here is ordinary rather than fatal.
func WirelessRequestScan(ctx context.Context, conn nm.Conn, dev *nm.Device) error {
	b := dbusext.NewBase(conn, dev.Path)
	// The options dictionary is required and empty: a nil map of the right
	// type marshals as the empty a{sv} NetworkManager expects.
	var options map[string]any
	return b.Call(ctx, deviceWirelessRequestScan, options)
}

func WirelessGetPropertyPermHwAddress(ctx context.Context, conn nm.Conn, dev *nm.Device) (string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetString(ctx, deviceWirelessPropertyPermHwAddress)
}

func WirelessGetPropertyMode(ctx context.Context, conn nm.Conn, dev *nm.Device) (nm.WifiMode, error) {
	b := dbusext.NewBase(conn, dev.Path)
	mode, err := b.GetUint32(ctx, deviceWirelessPropertyMode)
	return nm.WifiMode(mode), err
}

func WirelessGetPropertyBitrate(ctx context.Context, conn nm.Conn, dev *nm.Device) (uint32, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetUint32(ctx, deviceWirelessPropertyBitrate)
}

func WirelessGetPropertyAccessPoints(ctx context.Context, conn nm.Conn, dev *nm.Device) ([]*nm.AccessPoint, error) {
	b := dbusext.NewBase(conn, dev.Path)
	paths, err := b.GetObjects(ctx, deviceWirelessPropertyAccessPoints)
	if err != nil {
		return nil, err
	}
	aps := make([]*nm.AccessPoint, len(paths))
	for i, path := range paths {
		aps[i], err = nm.NewAccessPoint(ctx, conn, path)
		if err != nil {
			return aps, err
		}
	}
	return aps, nil
}

func WirelessGetPropertyActiveAccessPoint(ctx context.Context, conn nm.Conn, dev *nm.Device) (*nm.AccessPoint, error) {
	b := dbusext.NewBase(conn, dev.Path)
	path, ok, err := b.GetOptionalObject(ctx, deviceWirelessPropertyActiveAccessPoint)
	if err != nil || !ok {
		return nil, err
	}
	return nm.NewAccessPoint(ctx, conn, path)
}

func WirelessGetPropertyWirelessCapabilities(ctx context.Context, conn nm.Conn, dev *nm.Device) (uint32, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetUint32(ctx, deviceWirelessPropertyWirelessCapabilities)
}

func WirelessGetPropertyLastScan(ctx context.Context, conn nm.Conn, dev *nm.Device) (int64, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetInt64(ctx, deviceWirelessPropertyLastScan)
}

// LastScans returns an iterator over the device's LastScan property, which
// NetworkManager updates each time a scan finishes. The value is
// CLOCK_BOOTTIME milliseconds, or -1 for a device that has never scanned.
// Iteration ends when ctx is done.
//
// A scan requested with [WirelessRequestScan] completes asynchronously, so
// this is how a caller learns the access point list is worth re-reading. The
// subscription is in place before this returns, so reading the current value
// afterwards cannot miss a scan that finishes in between. The match rule is
// removed when the range loop ends or when ctx is done, whichever comes
// first.
func LastScans(ctx context.Context, conn nm.Conn, dev *nm.Device) (iter.Seq[int64], error) {
	signals, err := dbusext.Signals(ctx, conn,
		dbus.WithMatchInterface(propertiesInterface),
		dbus.WithMatchMember(propertiesSignalChanged),
		dbus.WithMatchObjectPath(dev.Path),
	)
	if err != nil {
		return nil, err
	}

	return func(yield func(int64) bool) {
		for signal := range signals {
			last, ok := lastScanFrom(signal)
			if !ok {
				continue
			}
			if !yield(last) {
				return
			}
		}
	}, nil
}

// lastScanFrom extracts a LastScan value from a PropertiesChanged signal,
// reporting whether the signal carried one. The signal body is
// (interface_name, changed_properties, invalidated_properties).
func lastScanFrom(signal *dbus.Signal) (int64, bool) {
	if signal.Name != propertiesInterface+"."+propertiesSignalChanged || len(signal.Body) < 2 {
		return 0, false
	}
	if iface, ok := signal.Body[0].(string); !ok || iface != dbusext.DeviceWirelessInterface {
		return 0, false
	}
	changed, ok := signal.Body[1].(map[string]dbus.Variant)
	if !ok {
		return 0, false
	}
	v, ok := changed[deviceWirelessLastScanProperty]
	if !ok {
		return 0, false
	}
	last, ok := v.Value().(int64)
	return last, ok
}
