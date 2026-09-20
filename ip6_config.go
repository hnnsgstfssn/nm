package nm

import (
	"context"
	"errors"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	ip6ConfigPropertyAddressData = dbusext.IP6ConfigInterface + ".AddressData"
	ip6ConfigPropertyGateway     = dbusext.IP6ConfigInterface + ".Gateway"
	ip6ConfigPropertyRouteData   = dbusext.IP6ConfigInterface + ".RouteData"
	ip6ConfigPropertyNameservers = dbusext.IP6ConfigInterface + ".Nameservers"
	ip6ConfigPropertyDomains     = dbusext.IP6ConfigInterface + ".Domains"
	ip6ConfigPropertySearches    = dbusext.IP6ConfigInterface + ".Searches"
	ip6ConfigPropertyDNSOptions  = dbusext.IP6ConfigInterface + ".DnsOptions"
	ip6ConfigPropertyDNSPriority = dbusext.IP6ConfigInterface + ".DnsPriority"
)

type IP6AddressData struct {
	Address string
	Prefix  uint32
}

type IP6RouteData struct {
	AdditionalAttributes map[string]string
	Destination          string
	NextHop              string
	Prefix               uint32
	Metric               uint32
}

// IP6Config is an eagerly-populated snapshot of an IPv6 configuration.
type IP6Config struct {
	Path        dbus.ObjectPath
	AddressData []IP6AddressData
	Gateway     string
	RouteData   []IP6RouteData
	Nameservers [][]byte
	Domains     []string
	Searches    []string
	DNSOptions  []string
	DNSPriority int32
}

// NewIP6Config fetches all IPv6 configuration properties from D-Bus and
// returns a populated snapshot.
func NewIP6Config(ctx context.Context, conn Conn, objectPath dbus.ObjectPath) (*IP6Config, error) {
	b := dbusext.NewBase(conn, objectPath)

	addressRaw, err := b.GetVariantMaps(ctx, ip6ConfigPropertyAddressData)
	if err != nil {
		return nil, err
	}
	addressData, err := DecodeIP6AddressData(addressRaw)
	if err != nil {
		return nil, err
	}

	gateway, err := b.GetString(ctx, ip6ConfigPropertyGateway)
	if err != nil {
		return nil, err
	}

	routeRaw, err := b.GetVariantMaps(ctx, ip6ConfigPropertyRouteData)
	if err != nil {
		return nil, err
	}
	routeData, err := DecodeIP6RouteData(routeRaw)
	if err != nil {
		return nil, err
	}

	nameservers, err := b.GetByteSlices(ctx, ip6ConfigPropertyNameservers)
	if err != nil {
		return nil, err
	}
	domains, err := b.GetStrings(ctx, ip6ConfigPropertyDomains)
	if err != nil {
		return nil, err
	}
	searches, err := b.GetStrings(ctx, ip6ConfigPropertySearches)
	if err != nil {
		return nil, err
	}
	dnsOptions, err := b.GetStrings(ctx, ip6ConfigPropertyDNSOptions)
	if err != nil {
		return nil, err
	}
	dnsPriority, err := b.GetInt32(ctx, ip6ConfigPropertyDNSPriority)
	if err != nil {
		return nil, err
	}

	return &IP6Config{
		Path:        objectPath,
		AddressData: addressData,
		Gateway:     gateway,
		RouteData:   routeData,
		Nameservers: nameservers,
		Domains:     domains,
		Searches:    searches,
		DNSOptions:  dnsOptions,
		DNSPriority: dnsPriority,
	}, nil
}

// DecodeIP6AddressData decodes D-Bus variant address data into typed structs.
func DecodeIP6AddressData(addresses []map[string]dbus.Variant) ([]IP6AddressData, error) {
	out := make([]IP6AddressData, 0, len(addresses))
	for _, address := range addresses {
		var d IP6AddressData
		var ok bool
		d.Prefix, ok = address["prefix"].Value().(uint32)
		if !ok {
			return out, errors.New("unexpected variant type for prefix")
		}
		d.Address, ok = address["address"].Value().(string)
		if !ok {
			return out, errors.New("unexpected variant type for address")
		}
		out = append(out, d)
	}
	return out, nil
}

// DecodeIP6RouteData decodes D-Bus variant route data into typed structs.
func DecodeIP6RouteData(routes []map[string]dbus.Variant) ([]IP6RouteData, error) {
	out := make([]IP6RouteData, 0, len(routes))
	for _, route := range routes {
		var d IP6RouteData
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
