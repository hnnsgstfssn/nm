package settings

import (
	"context"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	settingsListConnections      = dbusext.SettingsInterface + ".ListConnections"
	settingsGetConnectionByUUID  = dbusext.SettingsInterface + ".GetConnectionByUuid"
	settingsAddConnection        = dbusext.SettingsInterface + ".AddConnection"
	settingsAddConnectionUnsaved = dbusext.SettingsInterface + ".AddConnectionUnsaved"
	settingsReloadConnections    = dbusext.SettingsInterface + ".ReloadConnections"
	settingsSaveHostname         = dbusext.SettingsInterface + ".SaveHostname"

	settingsPropertyHostname  = dbusext.SettingsInterface + ".Hostname"
	settingsPropertyCanModify = dbusext.SettingsInterface + ".CanModify"
)

func newBase(conn nm.Conn) dbusext.Base {
	return dbusext.NewBase(conn, dbusext.SettingsObjectPath)
}

// ListConnections returns all saved network connections.
func ListConnections(ctx context.Context, conn nm.Conn) ([]*nm.Connection, error) {
	b := newBase(conn)
	var paths []dbus.ObjectPath
	if err := b.CallWithReturn(ctx, &paths, settingsListConnections); err != nil {
		return nil, err
	}
	connections := make([]*nm.Connection, len(paths))
	for i, path := range paths {
		connections[i] = nm.NewConnection(path)
	}
	return connections, nil
}

// ReloadConnections tells NetworkManager to reload config files from disk.
func ReloadConnections(ctx context.Context, conn nm.Conn) error {
	b := newBase(conn)
	return b.Call(ctx, settingsReloadConnections)
}

// GetConnectionByUUID returns the connection for the given UUID.
func GetConnectionByUUID(ctx context.Context, conn nm.Conn, uuid string) (*nm.Connection, error) {
	b := newBase(conn)
	var path dbus.ObjectPath
	if err := b.CallWithReturn(ctx, &path, settingsGetConnectionByUUID, uuid); err != nil {
		return nil, err
	}
	return nm.NewConnection(path), nil
}

// AddConnection adds a new connection and saves it to disk.
func AddConnection(ctx context.Context, conn nm.Conn, s nm.ConnectionSettings) (*nm.Connection, error) {
	b := newBase(conn)
	var path dbus.ObjectPath
	if err := b.CallWithReturn(ctx, &path, settingsAddConnection, s); err != nil {
		return nil, err
	}
	return nm.NewConnection(path), nil
}

// AddConnectionUnsaved adds a new connection without saving to disk.
func AddConnectionUnsaved(ctx context.Context, conn nm.Conn, s nm.ConnectionSettings) (*nm.Connection, error) {
	b := newBase(conn)
	var path dbus.ObjectPath
	if err := b.CallWithReturn(ctx, &path, settingsAddConnectionUnsaved, s); err != nil {
		return nil, err
	}
	return nm.NewConnection(path), nil
}

// SaveHostname persists the hostname to configuration.
func SaveHostname(ctx context.Context, conn nm.Conn, hostname string) error {
	b := newBase(conn)
	return b.Call(ctx, settingsSaveHostname, hostname)
}

func GetPropertyHostname(ctx context.Context, conn nm.Conn) (string, error) {
	b := newBase(conn)
	return b.GetString(ctx, settingsPropertyHostname)
}

func GetPropertyCanModify(ctx context.Context, conn nm.Conn) (bool, error) {
	b := newBase(conn)
	return b.GetBool(ctx, settingsPropertyCanModify)
}
