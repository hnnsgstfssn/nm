package nm

import (
	"context"
	"errors"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	dnsManagerPropertyMode          = dbusext.DNSManagerInterface + ".Mode"
	dnsManagerPropertyRcManager     = dbusext.DNSManagerInterface + ".RcManager"
	dnsManagerPropertyConfiguration = dbusext.DNSManagerInterface + ".Configuration"
)

// DNSConfigurationData holds a single DNS configuration entry.
type DNSConfigurationData struct {
	Interface   string
	Nameservers []string
	Priority    int32
	Vpn         bool
}

// DNSManager is an eagerly-populated snapshot of NetworkManager's DNS
// manager state.
type DNSManager struct {
	Path          dbus.ObjectPath
	Mode          string
	RcManager     string
	Configuration []DNSConfigurationData
}

// NewDNSManager fetches all DNS manager properties and returns a snapshot.
func NewDNSManager(ctx context.Context, conn Conn) (*DNSManager, error) {
	b := dbusext.NewBase(conn, dbusext.DNSManagerObjectPath)

	mode, err := b.GetString(ctx, dnsManagerPropertyMode)
	if err != nil {
		return nil, err
	}
	rcManager, err := b.GetString(ctx, dnsManagerPropertyRcManager)
	if err != nil {
		return nil, err
	}
	configurations, err := b.GetVariantMaps(ctx, dnsManagerPropertyConfiguration)
	if err != nil {
		return nil, err
	}
	cfgs, err := decodeDNSConfigurations(configurations)
	if err != nil {
		return nil, err
	}

	return &DNSManager{
		Path:          dbusext.DNSManagerObjectPath,
		Mode:          mode,
		RcManager:     rcManager,
		Configuration: cfgs,
	}, nil
}

func decodeDNSConfigurations(configurations []map[string]dbus.Variant) ([]DNSConfigurationData, error) {
	ret := make([]DNSConfigurationData, len(configurations))
	for i, conf := range configurations {
		if v, ok := conf["nameservers"]; ok {
			servers, ok := v.Value().([]string)
			if !ok {
				return nil, errors.New("unexpected variant type for nameservers")
			}
			ret[i].Nameservers = servers
		}
		if v, ok := conf["priority"]; ok {
			priority, ok := v.Value().(int32)
			if !ok {
				return nil, errors.New("unexpected variant type for priority")
			}
			ret[i].Priority = priority
		}
		if v, ok := conf["interface"]; ok {
			iface, ok := v.Value().(string)
			if !ok {
				return nil, errors.New("unexpected variant type for interface")
			}
			ret[i].Interface = iface
		}
		if v, ok := conf["vpn"]; ok {
			vpn, ok := v.Value().(bool)
			if !ok {
				return nil, errors.New("unexpected variant type for vpn")
			}
			ret[i].Vpn = vpn
		}
	}
	return ret, nil
}
