package device

import (
	"context"

	"github.com/hnnsgstfssn/nm"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	deviceWiredPropertyHwAddress       = dbusext.DeviceWiredInterface + ".HwAddress"
	deviceWiredPropertyPermHwAddress   = dbusext.DeviceWiredInterface + ".PermHwAddress"
	deviceWiredPropertySpeed           = dbusext.DeviceWiredInterface + ".Speed"
	deviceWiredPropertyS390Subchannels = dbusext.DeviceWiredInterface + ".S390Subchannels"
	deviceWiredPropertyCarrier         = dbusext.DeviceWiredInterface + ".Carrier"
)

func WiredGetPropertyHwAddress(ctx context.Context, conn nm.Conn, dev *nm.Device) (string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetString(ctx, deviceWiredPropertyHwAddress)
}

func WiredGetPropertyPermHwAddress(ctx context.Context, conn nm.Conn, dev *nm.Device) (string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetString(ctx, deviceWiredPropertyPermHwAddress)
}

func WiredGetPropertySpeed(ctx context.Context, conn nm.Conn, dev *nm.Device) (uint32, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetUint32(ctx, deviceWiredPropertySpeed)
}

func WiredGetPropertyS390Subchannels(ctx context.Context, conn nm.Conn, dev *nm.Device) ([]string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetStrings(ctx, deviceWiredPropertyS390Subchannels)
}

func WiredGetPropertyCarrier(ctx context.Context, conn nm.Conn, dev *nm.Device) (bool, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetBool(ctx, deviceWiredPropertyCarrier)
}
