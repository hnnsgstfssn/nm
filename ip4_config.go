package nm

import (
	"context"
	"errors"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	ip4ConfigPropertyAddressData    = dbusext.IP4ConfigInterface + ".AddressData"
	ip4ConfigPropertyGateway        = dbusext.IP4ConfigInterface + ".Gateway"
	ip4ConfigPropertyRouteData      = dbusext.IP4ConfigInterface + ".RouteData"
	ip4ConfigPropertyNameserverData = dbusext.IP4ConfigInterface + ".NameserverData"
	ip4ConfigPropertyDomains        = dbusext.IP4ConfigInterface + ".Domains"
	ip4ConfigPropertySearches       = dbusext.IP4ConfigInterface + ".Searches"
	ip4ConfigPropertyDNSOptions     = dbusext.IP4ConfigInterface + ".DnsOptions"
	ip4ConfigPropertyDNSPriority    = dbusext.IP4ConfigInterface + ".DnsPriority"
	ip4ConfigPropertyWinsServerData = dbusext.IP4ConfigInterface + ".WinsServerData"
)

type IP4AddressData struct {
	Address string
	Prefix  uint32
}

type IP4RouteData struct {
	AdditionalAttributes map[string]string
	Destination          string
	NextHop              string
	Prefix               uint32
	Metric               uint32
}

type IP4NameserverData struct {
	Address string
}

// IP4Config is an eagerly-populated snapshot of an IPv4 configuration.
type IP4Config struct {
	Path           dbus.ObjectPath
	Gateway        string
	AddressData    []IP4AddressData
	RouteData      []IP4RouteData
	NameserverData []IP4NameserverData
	Domains        []string
	Searches       []string
	DNSOptions     []string
	WinsServerData []string
	DNSPriority    int32
}

// NewIP4Config fetches all IPv4 configuration properties from D-Bus and
// returns a populated snapshot.
func NewIP4Config(ctx context.Context, conn Conn, objectPath dbus.ObjectPath) (*IP4Config, error) {
	b := dbusext.NewBase(conn, objectPath)

	addressRaw, err := b.GetVariantMaps(ctx, ip4ConfigPropertyAddressData)
	if err != nil {
		return nil, err
	}
	addressData, err := DecodeIP4AddressData(addressRaw)
	if err != nil {
		return nil, err
	}

	gateway, err := b.GetString(ctx, ip4ConfigPropertyGateway)
	if err != nil {
		return nil, err
	}

	routeRaw, err := b.GetVariantMaps(ctx, ip4ConfigPropertyRouteData)
	if err != nil {
		return nil, err
	}
	routeData, err := DecodeIP4RouteData(routeRaw)
	if err != nil {
		return nil, err
	}

	nsRaw, err := b.GetVariantMaps(ctx, ip4ConfigPropertyNameserverData)
	if err != nil {
		return nil, err
	}
	nsData, err := DecodeIP4NameserverData(nsRaw)
	if err != nil {
		return nil, err
	}

	domains, err := b.GetStrings(ctx, ip4ConfigPropertyDomains)
	if err != nil {
		return nil, err
	}
	searches, err := b.GetStrings(ctx, ip4ConfigPropertySearches)
	if err != nil {
		return nil, err
	}
	dnsOptions, err := b.GetStrings(ctx, ip4ConfigPropertyDNSOptions)
	if err != nil {
		return nil, err
	}
	dnsPriority, err := b.GetInt32(ctx, ip4ConfigPropertyDNSPriority)
	if err != nil {
		return nil, err
	}
	winsServerData, err := b.GetStrings(ctx, ip4ConfigPropertyWinsServerData)
	if err != nil {
		return nil, err
	}

	return &IP4Config{
		Path:           objectPath,
		AddressData:    addressData,
		Gateway:        gateway,
		RouteData:      routeData,
		NameserverData: nsData,
		Domains:        domains,
		Searches:       searches,
		DNSOptions:     dnsOptions,
		DNSPriority:    dnsPriority,
		WinsServerData: winsServerData,
	}, nil
}

// DecodeIP4AddressData decodes D-Bus variant address data into typed structs.
func DecodeIP4AddressData(addresses []map[string]dbus.Variant) ([]IP4AddressData, error) {
	out := make([]IP4AddressData, 0, len(addresses))
	for _, address := range addresses {
		var d IP4AddressData
		for name, attr := range address {
			switch name {
			case "address":
				v, ok := attr.Value().(string)
				if !ok {
					return out, errors.New("unexpected variant type for address")
				}
				d.Address = v
			case "prefix":
				v, ok := attr.Value().(uint32)
				if !ok {
					return out, errors.New("unexpected variant type for prefix")
				}
				d.Prefix = v
			}
		}
		out = append(out, d)
	}
	return out, nil
}

// DecodeIP4RouteData decodes D-Bus variant route data into typed structs.
func DecodeIP4RouteData(routes []map[string]dbus.Variant) ([]IP4RouteData, error) {
	out := make([]IP4RouteData, 0, len(routes))
	for _, route := range routes {
		var d IP4RouteData
		for name, attr := range route {
			switch name {
			case "dest":
				v, ok := attr.Value().(string)
				if !ok {
					return out, errors.New("unexpected variant type for dest")
				}
				d.Destination = v
			case "prefix":
				v, ok := attr.Value().(uint32)
				if !ok {
					return out, errors.New("unexpected variant type for prefix")
				}
				d.Prefix = v
			case "next-hop":
				v, ok := attr.Value().(string)
				if !ok {
					return out, errors.New("unexpected variant type for next-hop")
				}
				d.NextHop = v
			case "metric":
				v, ok := attr.Value().(uint32)
				if !ok {
					return out, errors.New("unexpected variant type for metric")
				}
				d.Metric = v
			default:
				if d.AdditionalAttributes == nil {
					d.AdditionalAttributes = make(map[string]string)
				}
				d.AdditionalAttributes[name] = attr.String()
			}
		}
		out = append(out, d)
	}
	return out, nil
}

// DecodeIP4NameserverData decodes D-Bus variant nameserver data into typed structs.
func DecodeIP4NameserverData(nameserverData []map[string]dbus.Variant) ([]IP4NameserverData, error) {
	out := make([]IP4NameserverData, 0, len(nameserverData))
	for _, ns := range nameserverData {
		addr, ok := ns["address"].Value().(string)
		if !ok {
			return out, errors.New("unexpected variant type for address")
		}
		out = append(out, IP4NameserverData{Address: addr})
	}
	return out, nil
}
