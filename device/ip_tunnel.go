package device

import (
	"context"

	"github.com/hnnsgstfssn/nm"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	deviceIPTunnelPropertyMode               = dbusext.DeviceIPTunnelInterface + ".Mode"
	deviceIPTunnelPropertyParent             = dbusext.DeviceIPTunnelInterface + ".Parent"
	deviceIPTunnelPropertyLocal              = dbusext.DeviceIPTunnelInterface + ".Local"
	deviceIPTunnelPropertyRemote             = dbusext.DeviceIPTunnelInterface + ".Remote"
	deviceIPTunnelPropertyTTL                = dbusext.DeviceIPTunnelInterface + ".Ttl"
	deviceIPTunnelPropertyTos                = dbusext.DeviceIPTunnelInterface + ".Tos"
	deviceIPTunnelPropertyPathMtuDiscovery   = dbusext.DeviceIPTunnelInterface + ".PathMtuDiscovery"
	deviceIPTunnelPropertyInputKey           = dbusext.DeviceIPTunnelInterface + ".InputKey"
	deviceIPTunnelPropertyOutputKey          = dbusext.DeviceIPTunnelInterface + ".OutputKey"
	deviceIPTunnelPropertyEncapsulationLimit = dbusext.DeviceIPTunnelInterface + ".EncapsulationLimit"
	deviceIPTunnelPropertyFlowLabel          = dbusext.DeviceIPTunnelInterface + ".FlowLabel"
	deviceIPTunnelPropertyFlags              = dbusext.DeviceIPTunnelInterface + ".Flags"
)

func IPTunnelGetPropertyMode(ctx context.Context, conn nm.Conn, dev *nm.Device) (uint32, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetUint32(ctx, deviceIPTunnelPropertyMode)
}

func IPTunnelGetPropertyParent(ctx context.Context, conn nm.Conn, dev *nm.Device) (*nm.Device, error) {
	b := dbusext.NewBase(conn, dev.Path)
	path, ok, err := b.GetOptionalObject(ctx, deviceIPTunnelPropertyParent)
	if err != nil || !ok {
		return nil, err
	}
	return nm.NewDevice(path), nil
}

func IPTunnelGetPropertyLocal(ctx context.Context, conn nm.Conn, dev *nm.Device) (string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetString(ctx, deviceIPTunnelPropertyLocal)
}

func IPTunnelGetPropertyRemote(ctx context.Context, conn nm.Conn, dev *nm.Device) (string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetString(ctx, deviceIPTunnelPropertyRemote)
}

func IPTunnelGetPropertyTTL(ctx context.Context, conn nm.Conn, dev *nm.Device) (uint8, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetUint8(ctx, deviceIPTunnelPropertyTTL)
}

func IPTunnelGetPropertyTos(ctx context.Context, conn nm.Conn, dev *nm.Device) (uint8, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetUint8(ctx, deviceIPTunnelPropertyTos)
}

func IPTunnelGetPropertyPathMtuDiscovery(ctx context.Context, conn nm.Conn, dev *nm.Device) (bool, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetBool(ctx, deviceIPTunnelPropertyPathMtuDiscovery)
}

func IPTunnelGetPropertyInputKey(ctx context.Context, conn nm.Conn, dev *nm.Device) (string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetString(ctx, deviceIPTunnelPropertyInputKey)
}

func IPTunnelGetPropertyOutputKey(ctx context.Context, conn nm.Conn, dev *nm.Device) (string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetString(ctx, deviceIPTunnelPropertyOutputKey)
}

func IPTunnelGetPropertyEncapsulationLimit(ctx context.Context, conn nm.Conn, dev *nm.Device) (uint8, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetUint8(ctx, deviceIPTunnelPropertyEncapsulationLimit)
}

func IPTunnelGetPropertyFlowLabel(ctx context.Context, conn nm.Conn, dev *nm.Device) (uint32, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetUint32(ctx, deviceIPTunnelPropertyFlowLabel)
}

func IPTunnelGetPropertyFlags(ctx context.Context, conn nm.Conn, dev *nm.Device) (uint32, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetUint32(ctx, deviceIPTunnelPropertyFlags)
}
