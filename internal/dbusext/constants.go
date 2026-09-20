package dbusext

import "github.com/godbus/dbus/v5"

// NoObject is what NetworkManager reports for an object property whose
// relationship does not currently exist: a device with no active connection,
// a wireless device that has not associated. The D-Bus specification has no
// null object path, so "/" carries that meaning by convention.
const NoObject dbus.ObjectPath = "/"

// D-Bus interface and object path constants for NetworkManager.
const (
	NetworkManagerInterface  = "org.freedesktop.NetworkManager"
	NetworkManagerObjectPath = "/org/freedesktop/NetworkManager"

	DeviceInterface           = NetworkManagerInterface + ".Device"
	DeviceWirelessInterface   = DeviceInterface + ".Wireless"
	DeviceBridgeInterface     = DeviceInterface + ".Bridge"
	DeviceWiredInterface      = DeviceInterface + ".Wired"
	DeviceGenericInterface    = DeviceInterface + ".Generic"
	DeviceIPTunnelInterface   = DeviceInterface + ".IPTunnel"
	DeviceStatisticsInterface = DeviceInterface + ".Statistics"

	SettingsInterface  = NetworkManagerInterface + ".Settings"
	SettingsObjectPath = NetworkManagerObjectPath + "/Settings"

	ConnectionInterface = SettingsInterface + ".Connection"

	ActiveConnectionInterface = NetworkManagerInterface + ".Connection.Active"

	AccessPointInterface = NetworkManagerInterface + ".AccessPoint"

	CheckpointInterface = NetworkManagerInterface + ".Checkpoint"

	IP4ConfigInterface = NetworkManagerInterface + ".IP4Config"
	IP6ConfigInterface = NetworkManagerInterface + ".IP6Config"

	DHCP4ConfigInterface = NetworkManagerInterface + ".DHCP4Config"
	DHCP6ConfigInterface = NetworkManagerInterface + ".DHCP6Config"

	DNSManagerInterface  = NetworkManagerInterface + ".DnsManager"
	DNSManagerObjectPath = "/org/freedesktop/NetworkManager/DnsManager"

	VpnConnectionInterface = NetworkManagerInterface + ".VPN.Connection"
)
