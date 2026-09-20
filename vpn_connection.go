package nm

import (
	"context"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	vpnConnectionPropertyVpnState = dbusext.VpnConnectionInterface + ".VpnState"
	vpnConnectionPropertyBanner   = dbusext.VpnConnectionInterface + ".Banner"
)

// VpnConnection is an eagerly-populated snapshot of a VPN connection's
// properties.
type VpnConnection struct {
	Path     dbus.ObjectPath
	Banner   string
	VpnState uint32
}

// NewVpnConnection fetches all VPN connection properties from D-Bus and
// returns a populated snapshot.
func NewVpnConnection(ctx context.Context, conn Conn, objectPath dbus.ObjectPath) (*VpnConnection, error) {
	b := dbusext.NewBase(conn, objectPath)

	state, err := b.GetUint32(ctx, vpnConnectionPropertyVpnState)
	if err != nil {
		return nil, err
	}
	banner, err := b.GetString(ctx, vpnConnectionPropertyBanner)
	if err != nil {
		return nil, err
	}

	return &VpnConnection{
		Path:     objectPath,
		VpnState: state,
		Banner:   banner,
	}, nil
}
