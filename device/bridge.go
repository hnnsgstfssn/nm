package device

import (
	"context"

	"github.com/hnnsgstfssn/nm"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	deviceBridgePropertySlaves  = dbusext.DeviceBridgeInterface + ".Slaves"
	deviceBridgePropertyCarrier = dbusext.DeviceBridgeInterface + ".Carrier"
)

func BridgeGetPropertySlaves(ctx context.Context, conn nm.Conn, dev *nm.Device) ([]*nm.Device, error) {
	b := dbusext.NewBase(conn, dev.Path)
	paths, err := b.GetObjects(ctx, deviceBridgePropertySlaves)
	if err != nil {
		return nil, err
	}
	devices := make([]*nm.Device, len(paths))
	for i, path := range paths {
		devices[i] = nm.NewDevice(path)
	}
	return devices, nil
}

func BridgeGetPropertyCarrier(ctx context.Context, conn nm.Conn, dev *nm.Device) (bool, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetBool(ctx, deviceBridgePropertyCarrier)
}
