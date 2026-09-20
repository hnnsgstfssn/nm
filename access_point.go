package nm

import (
	"context"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	accessPointPropertyFlags      = dbusext.AccessPointInterface + ".Flags"
	accessPointPropertyWpaFlags   = dbusext.AccessPointInterface + ".WpaFlags"
	accessPointPropertyRsnFlags   = dbusext.AccessPointInterface + ".RsnFlags"
	accessPointPropertySsid       = dbusext.AccessPointInterface + ".Ssid"
	accessPointPropertyFrequency  = dbusext.AccessPointInterface + ".Frequency"
	accessPointPropertyHwAddress  = dbusext.AccessPointInterface + ".HwAddress"
	accessPointPropertyMode       = dbusext.AccessPointInterface + ".Mode"
	accessPointPropertyMaxBitrate = dbusext.AccessPointInterface + ".MaxBitrate"
	accessPointPropertyStrength   = dbusext.AccessPointInterface + ".Strength"
	accessPointPropertyLastSeen   = dbusext.AccessPointInterface + ".LastSeen"
)

// AccessPoint is an eagerly-populated snapshot of a Wi-Fi access point's
// properties. The D-Bus connection is not retained.
type AccessPoint struct {
	Path       dbus.ObjectPath
	SSID       string
	HWAddress  string
	Flags      APFlags
	WPAFlags   APSecurity
	RSNFlags   APSecurity
	Frequency  uint32
	Mode       WifiMode
	MaxBitrate uint32
	LastSeen   int32
	Strength   uint8
}

// NewAccessPoint fetches all access point properties from D-Bus and returns
// a populated snapshot.
func NewAccessPoint(ctx context.Context, conn Conn, objectPath dbus.ObjectPath) (*AccessPoint, error) {
	b := dbusext.NewBase(conn, objectPath)

	flags, err := b.GetUint32(ctx, accessPointPropertyFlags)
	if err != nil {
		return nil, err
	}
	wpaFlags, err := b.GetUint32(ctx, accessPointPropertyWpaFlags)
	if err != nil {
		return nil, err
	}
	rsnFlags, err := b.GetUint32(ctx, accessPointPropertyRsnFlags)
	if err != nil {
		return nil, err
	}
	ssidBytes, err := b.GetBytes(ctx, accessPointPropertySsid)
	if err != nil {
		return nil, err
	}
	frequency, err := b.GetUint32(ctx, accessPointPropertyFrequency)
	if err != nil {
		return nil, err
	}
	hwAddress, err := b.GetString(ctx, accessPointPropertyHwAddress)
	if err != nil {
		return nil, err
	}
	mode, err := b.GetUint32(ctx, accessPointPropertyMode)
	if err != nil {
		return nil, err
	}
	maxBitrate, err := b.GetUint32(ctx, accessPointPropertyMaxBitrate)
	if err != nil {
		return nil, err
	}
	strength, err := b.GetUint8(ctx, accessPointPropertyStrength)
	if err != nil {
		return nil, err
	}
	lastSeen, err := b.GetInt32(ctx, accessPointPropertyLastSeen)
	if err != nil {
		return nil, err
	}

	return &AccessPoint{
		Path:       objectPath,
		Flags:      APFlags(flags),
		WPAFlags:   APSecurity(wpaFlags),
		RSNFlags:   APSecurity(rsnFlags),
		SSID:       string(ssidBytes),
		Frequency:  frequency,
		HWAddress:  hwAddress,
		Mode:       WifiMode(mode),
		MaxBitrate: maxBitrate,
		Strength:   strength,
		LastSeen:   lastSeen,
	}, nil
}

// Secured reports whether joining ap needs a password. The WPA and RSN flags
// cover WPA, WPA2 and WPA3; the privacy flag on its own means WEP.
func (a *AccessPoint) Secured() bool {
	return a.WPAFlags != 0 || a.RSNFlags != 0 || a.Flags&APFlagsPrivacy != 0
}
