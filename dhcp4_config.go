package nm

import (
	"context"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	dhcp4ConfigPropertyOptions = dbusext.DHCP4ConfigInterface + ".Options"
)

// DHCP4Options holds the key/value pairs returned by the DHCPv4 server.
type DHCP4Options map[string]any

// DHCP4Config is an eagerly-populated snapshot of DHCPv4 configuration.
type DHCP4Config struct {
	Options DHCP4Options
}

// NewDHCP4Config fetches all DHCPv4 properties from D-Bus and returns a
// populated snapshot.
func NewDHCP4Config(ctx context.Context, conn Conn, objectPath dbus.ObjectPath) (*DHCP4Config, error) {
	b := dbusext.NewBase(conn, objectPath)

	raw, err := b.GetVariantMap(ctx, dhcp4ConfigPropertyOptions)
	if err != nil {
		return nil, err
	}

	opts := make(DHCP4Options, len(raw))
	for k, v := range raw {
		opts[k] = v.Value()
	}

	return &DHCP4Config{Options: opts}, nil
}
