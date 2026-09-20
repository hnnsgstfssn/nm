package connection

import (
	"context"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	connectionUpdate        = dbusext.ConnectionInterface + ".Update"
	connectionUpdateUnsaved = dbusext.ConnectionInterface + ".UpdateUnsaved"
	connectionDelete        = dbusext.ConnectionInterface + ".Delete"
	connectionGetSettings   = dbusext.ConnectionInterface + ".GetSettings"
	connectionGetSecrets    = dbusext.ConnectionInterface + ".GetSecrets"
	connectionClearSecrets  = dbusext.ConnectionInterface + ".ClearSecrets"
	connectionSave          = dbusext.ConnectionInterface + ".Save"

	connectionPropertyUnsaved  = dbusext.ConnectionInterface + ".Unsaved"
	connectionPropertyFlags    = dbusext.ConnectionInterface + ".Flags"
	connectionPropertyFilename = dbusext.ConnectionInterface + ".Filename"
)

// Update replaces all connection settings and saves to disk.
func Update(ctx context.Context, conn nm.Conn, c *nm.Connection, settings nm.ConnectionSettings) error {
	b := dbusext.NewBase(conn, c.Path)
	return b.Call(ctx, connectionUpdate, settings)
}

// UpdateUnsaved replaces all connection settings but does not save to disk.
func UpdateUnsaved(ctx context.Context, conn nm.Conn, c *nm.Connection, settings nm.ConnectionSettings) error {
	b := dbusext.NewBase(conn, c.Path)
	return b.Call(ctx, connectionUpdateUnsaved, settings)
}

// Delete removes this connection.
func Delete(ctx context.Context, conn nm.Conn, c *nm.Connection) error {
	b := dbusext.NewBase(conn, c.Path)
	return b.Call(ctx, connectionDelete)
}

// GetSettings returns the settings maps describing this network configuration.
// Secrets are not included; use GetSecrets separately.
func GetSettings(ctx context.Context, conn nm.Conn, c *nm.Connection) (nm.ConnectionSettings, error) {
	b := dbusext.NewBase(conn, c.Path)
	var settings map[string]map[string]dbus.Variant
	if err := b.CallWithReturn(ctx, &settings, connectionGetSettings); err != nil {
		return nil, err
	}
	return decodeSettings(settings), nil
}

// GetSecrets returns the secrets belonging to this network configuration.
func GetSecrets(ctx context.Context, conn nm.Conn, c *nm.Connection, settingName string) (nm.ConnectionSettings, error) {
	b := dbusext.NewBase(conn, c.Path)
	var settings map[string]map[string]dbus.Variant
	if err := b.CallWithReturn(ctx, &settings, connectionGetSecrets, settingName); err != nil {
		return nil, err
	}
	return decodeSettings(settings), nil
}

// ClearSecrets removes secrets from this connection profile.
func ClearSecrets(ctx context.Context, conn nm.Conn, c *nm.Connection) error {
	b := dbusext.NewBase(conn, c.Path)
	return b.Call(ctx, connectionClearSecrets)
}

// Save persists unsaved changes to disk.
func Save(ctx context.Context, conn nm.Conn, c *nm.Connection) error {
	b := dbusext.NewBase(conn, c.Path)
	return b.Call(ctx, connectionSave)
}

func GetPropertyUnsaved(ctx context.Context, conn nm.Conn, c *nm.Connection) (bool, error) {
	b := dbusext.NewBase(conn, c.Path)
	return b.GetBool(ctx, connectionPropertyUnsaved)
}

func GetPropertyFlags(ctx context.Context, conn nm.Conn, c *nm.Connection) (uint32, error) {
	b := dbusext.NewBase(conn, c.Path)
	return b.GetUint32(ctx, connectionPropertyFlags)
}

func GetPropertyFilename(ctx context.Context, conn nm.Conn, c *nm.Connection) (string, error) {
	b := dbusext.NewBase(conn, c.Path)
	return b.GetString(ctx, connectionPropertyFilename)
}

func decodeSettings(input map[string]map[string]dbus.Variant) nm.ConnectionSettings {
	out := nm.ConnectionSettings{}
	for key, data := range input {
		out[key], _ = decode(data).(map[string]any) // nolint:errcheck
	}
	return out
}

func decode(input any) any {
	if variant, ok := input.(dbus.Variant); ok {
		return decode(variant.Value())
	}
	if m, ok := input.(map[string]dbus.Variant); ok {
		return decodeMap(m)
	}
	if a, ok := input.([]dbus.Variant); ok {
		return decodeArray(a)
	}
	if a, ok := input.([]map[string]dbus.Variant); ok {
		return decodeMapArray(a)
	}
	return input
}

func decodeArray(input []dbus.Variant) []any {
	var out []any
	for _, v := range input {
		out = append(out, decode(v))
	}
	return out
}

func decodeMapArray(input []map[string]dbus.Variant) []map[string]any {
	var out []map[string]any
	for _, m := range input {
		out = append(out, decodeMap(m))
	}
	return out
}

func decodeMap(input map[string]dbus.Variant) map[string]any {
	out := map[string]any{}
	for key, v := range input {
		out[key] = decode(v)
	}
	return out
}
