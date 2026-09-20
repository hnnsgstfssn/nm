package nm

import (
	"context"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	dhcp6ConfigPropertyOptions = dbusext.DHCP6ConfigInterface + ".Options"
)

// DHCP6Options holds the key/value pairs returned by the DHCPv6 server.
type DHCP6Options map[string]any

// DHCP6Config is an eagerly-populated snapshot of DHCPv6 configuration.
type DHCP6Config struct {
	Options DHCP6Options
}

// NewDHCP6Config fetches all DHCPv6 properties from D-Bus and returns a
// populated snapshot.
func NewDHCP6Config(ctx context.Context, conn Conn, objectPath dbus.ObjectPath) (*DHCP6Config, error) {
	b := dbusext.NewBase(conn, objectPath)

	raw, err := b.GetVariantMap(ctx, dhcp6ConfigPropertyOptions)
	if err != nil {
		return nil, err
	}

	opts := make(DHCP6Options, len(raw))
	for k, v := range raw {
		opts[k] = v.Value()
	}

	return &DHCP6Config{Options: opts}, nil
}
