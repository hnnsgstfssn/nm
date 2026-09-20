package manager

import (
	"context"
	"errors"
	"fmt"
	"iter"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	networkManagerReload                          = dbusext.NetworkManagerInterface + ".Reload"
	networkManagerGetDevices                      = dbusext.NetworkManagerInterface + ".GetDevices"
	networkManagerGetAllDevices                   = dbusext.NetworkManagerInterface + ".GetAllDevices"
	networkManagerGetDeviceByIPIface              = dbusext.NetworkManagerInterface + ".GetDeviceByIpIface"
	networkManagerActivateConnection              = dbusext.NetworkManagerInterface + ".ActivateConnection"
	networkManagerAddAndActivateConnection        = dbusext.NetworkManagerInterface + ".AddAndActivateConnection"
	networkManagerDeactivateConnection            = dbusext.NetworkManagerInterface + ".DeactivateConnection"
	networkManagerSleep                           = dbusext.NetworkManagerInterface + ".Sleep"
	networkManagerEnable                          = dbusext.NetworkManagerInterface + ".Enable"
	networkManagerCheckConnectivity               = dbusext.NetworkManagerInterface + ".CheckConnectivity"
	networkManagerState                           = dbusext.NetworkManagerInterface + ".state"
	networkManagerCheckpointCreate                = dbusext.NetworkManagerInterface + ".CheckpointCreate"
	networkManagerCheckpointDestroy               = dbusext.NetworkManagerInterface + ".CheckpointDestroy"
	networkManagerCheckpointRollback              = dbusext.NetworkManagerInterface + ".CheckpointRollback"
	networkManagerCheckpointAdjustRollbackTimeout = dbusext.NetworkManagerInterface + ".CheckpointAdjustRollbackTimeout"

	networkManagerPropertyDevices                    = dbusext.NetworkManagerInterface + ".Devices"
	networkManagerPropertyAllDevices                 = dbusext.NetworkManagerInterface + ".AllDevices"
	networkManagerPropertyCheckpoints                = dbusext.NetworkManagerInterface + ".Checkpoints"
	networkManagerPropertyNetworkingEnabled          = dbusext.NetworkManagerInterface + ".NetworkingEnabled"
	networkManagerPropertyWirelessEnabled            = dbusext.NetworkManagerInterface + ".WirelessEnabled"
	networkManagerPropertyWirelessHardwareEnabled    = dbusext.NetworkManagerInterface + ".WirelessHardwareEnabled"
	networkManagerPropertyWwanEnabled                = dbusext.NetworkManagerInterface + ".WwanEnabled"
	networkManagerPropertyWwanHardwareEnabled        = dbusext.NetworkManagerInterface + ".WwanHardwareEnabled"
	networkManagerPropertyWimaxEnabled               = dbusext.NetworkManagerInterface + ".WimaxEnabled"
	networkManagerPropertyWimaxHardwareEnabled       = dbusext.NetworkManagerInterface + ".WimaxHardwareEnabled"
	networkManagerPropertyActiveConnections          = dbusext.NetworkManagerInterface + ".ActiveConnections"
	networkManagerPropertyPrimaryConnection          = dbusext.NetworkManagerInterface + ".PrimaryConnection"
	networkManagerPropertyPrimaryConnectionType      = dbusext.NetworkManagerInterface + ".PrimaryConnectionType"
	networkManagerPropertyMetered                    = dbusext.NetworkManagerInterface + ".Metered"
	networkManagerPropertyActivatingConnection       = dbusext.NetworkManagerInterface + ".ActivatingConnection"
	networkManagerPropertyStartup                    = dbusext.NetworkManagerInterface + ".Startup"
	networkManagerPropertyVersion                    = dbusext.NetworkManagerInterface + ".Version"
	networkManagerPropertyCapabilities               = dbusext.NetworkManagerInterface + ".Capabilities"
	networkManagerPropertyState                      = dbusext.NetworkManagerInterface + ".State"
	networkManagerPropertyConnectivity               = dbusext.NetworkManagerInterface + ".Connectivity"
	networkManagerPropertyConnectivityCheckAvailable = dbusext.NetworkManagerInterface + ".ConnectivityCheckAvailable"
	networkManagerPropertyConnectivityCheckEnabled   = dbusext.NetworkManagerInterface + ".ConnectivityCheckEnabled"
)

func newBase(conn nm.Conn) dbusext.Base {
	return dbusext.NewBase(conn, dbusext.NetworkManagerObjectPath)
}

// Reload reloads NetworkManager configuration.
func Reload(ctx context.Context, conn nm.Conn, flags uint32) error {
	b := newBase(conn)
	return b.Call(ctx, networkManagerReload, flags)
}

// GetDevices returns the list of realized network devices.
func GetDevices(ctx context.Context, conn nm.Conn) ([]*nm.Device, error) {
	b := newBase(conn)
	var paths []dbus.ObjectPath
	if err := b.CallWithReturn(ctx, &paths, networkManagerGetDevices); err != nil {
		return nil, err
	}
	devices := make([]*nm.Device, len(paths))
	for i, path := range paths {
		devices[i] = nm.NewDevice(path)
	}
	return devices, nil
}

// GetAllDevices returns both realized and un-realized network devices.
func GetAllDevices(ctx context.Context, conn nm.Conn) ([]*nm.Device, error) {
	b := newBase(conn)
	var paths []dbus.ObjectPath
	if err := b.CallWithReturn(ctx, &paths, networkManagerGetAllDevices); err != nil {
		return nil, err
	}
	devices := make([]*nm.Device, len(paths))
	for i, path := range paths {
		devices[i] = nm.NewDevice(path)
	}
	return devices, nil
}

// GetDeviceByIPIface returns the device referenced by its IP interface name.
func GetDeviceByIPIface(ctx context.Context, conn nm.Conn, interfaceID string) (*nm.Device, error) {
	b := newBase(conn)
	var path dbus.ObjectPath
	if err := b.CallWithReturn(ctx, &path, networkManagerGetDeviceByIPIface, interfaceID); err != nil {
		return nil, err
	}
	return nm.NewDevice(path), nil
}

// devicePath is the path to pass for an optional device argument.
// NetworkManager reads [nm.NoObject] as "pick a device yourself"; the empty
// path is not a valid object path, so godbus rejects the message before it
// reaches the daemon.
func devicePath(device *nm.Device) dbus.ObjectPath {
	if device == nil {
		return nm.NoObject
	}
	return device.Path
}

// ActivateConnection activates a connection using the supplied device. A nil
// device lets NetworkManager choose one; a nil specificObject leaves the
// choice of access point or other sub-object to it as well.
func ActivateConnection(ctx context.Context, conn nm.Conn, connection *nm.Connection, device *nm.Device, specificObject *dbus.Object) (*nm.ActiveConnection, error) {
	b := newBase(conn)

	specificObjectPath := nm.NoObject
	if specificObject != nil {
		specificObjectPath = specificObject.Path()
	}

	var connectionPath dbus.ObjectPath
	if err := b.CallWithReturn(ctx, &connectionPath, networkManagerActivateConnection, connection.Path, devicePath(device), specificObjectPath); err != nil {
		return nil, err
	}
	return nm.NewActiveConnection(connectionPath), nil
}

// AddAndActivateConnection adds a new connection using the given details
// as a template, then activates it. A nil device lets NetworkManager choose
// one.
func AddAndActivateConnection(ctx context.Context, conn nm.Conn, connection map[string]map[string]any, device *nm.Device) (*nm.ActiveConnection, error) {
	b := newBase(conn)

	var profilePath, activePath dbus.ObjectPath
	if err := b.CallWithReturn2(ctx, &profilePath, &activePath, networkManagerAddAndActivateConnection, connection, devicePath(device), nm.NoObject); err != nil {
		return nil, err
	}
	return nm.NewActiveConnection(activePath), nil
}

// ActivateWirelessConnection activates an access point on a network device.
func ActivateWirelessConnection(ctx context.Context, conn nm.Conn, connection *nm.Connection, device *nm.Device, accessPoint *nm.AccessPoint) (*nm.ActiveConnection, error) {
	b := newBase(conn)
	var opath dbus.ObjectPath
	if err := b.CallWithReturn(ctx, &opath, networkManagerActivateConnection, connection.Path, device.Path, accessPoint.Path); err != nil {
		return nil, err
	}
	return nm.NewActiveConnection(opath), nil
}

// AddAndActivateWirelessConnection adds a new connection profile and
// activates it to the given access point.
func AddAndActivateWirelessConnection(ctx context.Context, conn nm.Conn, connection map[string]map[string]any, device *nm.Device, accessPoint *nm.AccessPoint) (*nm.ActiveConnection, error) {
	b := newBase(conn)
	var opath1, opath2 dbus.ObjectPath
	if err := b.CallWithReturn2(ctx, &opath1, &opath2, networkManagerAddAndActivateConnection, connection, device.Path, accessPoint.Path); err != nil {
		return nil, err
	}
	return nm.NewActiveConnection(opath2), nil
}

// DeactivateConnection deactivates an active connection.
func DeactivateConnection(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) error {
	b := newBase(conn)
	return b.Call(ctx, networkManagerDeactivateConnection, ac.Path)
}

// Sleep controls the sleep state of the NM daemon.
func Sleep(ctx context.Context, conn nm.Conn, sleepNWake bool) error {
	b := newBase(conn)
	return b.Call(ctx, networkManagerSleep, sleepNWake)
}

// Enable controls whether overall networking is enabled or disabled.
func Enable(ctx context.Context, conn nm.Conn, enableNDisable bool) error {
	b := newBase(conn)
	return b.Call(ctx, networkManagerEnable, enableNDisable)
}

// CheckConnectivity re-checks the network connectivity state.
func CheckConnectivity(ctx context.Context, conn nm.Conn) error {
	b := newBase(conn)
	return b.Call(ctx, networkManagerCheckConnectivity)
}

// State returns the overall networking state through NetworkManager's
// legacy state() method. [GetPropertyState] reads the same thing from the
// State property and is what new code should use; this exists for a daemon
// old enough not to export it.
func State(ctx context.Context, conn nm.Conn) (nm.State, error) {
	b := newBase(conn)
	var state uint32
	if err := b.CallWithReturn(ctx, &state, networkManagerState); err != nil {
		return nm.StateUnknown, err
	}
	return nm.State(state), nil
}

// CheckpointCreate creates a checkpoint of the current networking config.
func CheckpointCreate(ctx context.Context, conn nm.Conn, devices []*nm.Device, rollbackTimeout uint32, flags nm.CheckpointCreateFlags) (*nm.Checkpoint, error) {
	b := newBase(conn)
	paths := make([]dbus.ObjectPath, len(devices))
	for i, dev := range devices {
		paths[i] = dev.Path
	}

	var checkpointPath dbus.ObjectPath
	if err := b.CallWithReturn(ctx, &checkpointPath, networkManagerCheckpointCreate, paths, rollbackTimeout, flags); err != nil {
		return nil, err
	}
	return nm.NewCheckpoint(ctx, conn, checkpointPath)
}

// CheckpointDestroy destroys a previously created checkpoint.
func CheckpointDestroy(ctx context.Context, conn nm.Conn, checkpoint *nm.Checkpoint) error {
	b := newBase(conn)
	if checkpoint == nil {
		return b.Call(ctx, networkManagerCheckpointDestroy)
	}
	return b.Call(ctx, networkManagerCheckpointDestroy, checkpoint.Path)
}

// CheckpointRollback rolls back a checkpoint before the timeout, reporting
// the result per device.
func CheckpointRollback(ctx context.Context, conn nm.Conn, checkpoint *nm.Checkpoint) (map[dbus.ObjectPath]nm.RollbackResult, error) {
	b := newBase(conn)
	var raw map[dbus.ObjectPath]uint32
	if err := b.CallWithReturn(ctx, &raw, networkManagerCheckpointRollback, checkpoint.Path); err != nil {
		return nil, err
	}
	results := make(map[dbus.ObjectPath]nm.RollbackResult, len(raw))
	for path, result := range raw {
		results[path] = nm.RollbackResult(result)
	}
	return results, nil
}

// CheckpointAdjustRollbackTimeout resets the rollback timeout.
func CheckpointAdjustRollbackTimeout(ctx context.Context, conn nm.Conn, checkpoint *nm.Checkpoint, addTimeout uint32) error {
	b := newBase(conn)
	return b.Call(ctx, networkManagerCheckpointAdjustRollbackTimeout, checkpoint.Path, addTimeout)
}

/* PROPERTIES */

// GetPropertyDevices returns the list of realized network devices.
func GetPropertyDevices(ctx context.Context, conn nm.Conn) ([]*nm.Device, error) {
	b := newBase(conn)
	paths, err := b.GetObjects(ctx, networkManagerPropertyDevices)
	if err != nil {
		return nil, err
	}
	devices := make([]*nm.Device, len(paths))
	for i, path := range paths {
		devices[i] = nm.NewDevice(path)
	}
	return devices, nil
}

// GetPropertyAllDevices returns both realized and un-realized devices.
func GetPropertyAllDevices(ctx context.Context, conn nm.Conn) ([]*nm.Device, error) {
	b := newBase(conn)
	paths, err := b.GetObjects(ctx, networkManagerPropertyAllDevices)
	if err != nil {
		return nil, err
	}
	devices := make([]*nm.Device, len(paths))
	for i, path := range paths {
		devices[i] = nm.NewDevice(path)
	}
	return devices, nil
}

// GetPropertyCheckpoints returns the list of active checkpoints.
func GetPropertyCheckpoints(ctx context.Context, conn nm.Conn) ([]*nm.Checkpoint, error) {
	b := newBase(conn)
	paths, err := b.GetObjects(ctx, networkManagerPropertyCheckpoints)
	if err != nil {
		return nil, err
	}
	checkpoints := make([]*nm.Checkpoint, len(paths))
	for i, path := range paths {
		checkpoints[i], err = nm.NewCheckpoint(ctx, conn, path)
		if err != nil {
			return checkpoints, err
		}
	}
	return checkpoints, nil
}

// GetPropertyNetworkingEnabled reports whether overall networking is enabled.
func GetPropertyNetworkingEnabled(ctx context.Context, conn nm.Conn) (bool, error) {
	b := newBase(conn)
	return b.GetBool(ctx, networkManagerPropertyNetworkingEnabled)
}

// GetPropertyWirelessEnabled reports whether wireless is enabled.
func GetPropertyWirelessEnabled(ctx context.Context, conn nm.Conn) (bool, error) {
	b := newBase(conn)
	return b.GetBool(ctx, networkManagerPropertyWirelessEnabled)
}

// SetPropertyWirelessEnabled sets whether wireless is enabled.
func SetPropertyWirelessEnabled(ctx context.Context, conn nm.Conn, enabled bool) error {
	b := newBase(conn)
	return b.SetProperty(ctx, networkManagerPropertyWirelessEnabled, enabled)
}

// GetPropertyWirelessHardwareEnabled reports whether wireless HW is enabled.
func GetPropertyWirelessHardwareEnabled(ctx context.Context, conn nm.Conn) (bool, error) {
	b := newBase(conn)
	return b.GetBool(ctx, networkManagerPropertyWirelessHardwareEnabled)
}

// GetPropertyWwanEnabled reports whether mobile broadband is enabled.
func GetPropertyWwanEnabled(ctx context.Context, conn nm.Conn) (bool, error) {
	b := newBase(conn)
	return b.GetBool(ctx, networkManagerPropertyWwanEnabled)
}

// GetPropertyWwanHardwareEnabled reports whether mobile broadband HW is enabled.
func GetPropertyWwanHardwareEnabled(ctx context.Context, conn nm.Conn) (bool, error) {
	b := newBase(conn)
	return b.GetBool(ctx, networkManagerPropertyWwanHardwareEnabled)
}

// GetPropertyWimaxEnabled reports whether WiMAX devices are enabled.
func GetPropertyWimaxEnabled(ctx context.Context, conn nm.Conn) (bool, error) {
	b := newBase(conn)
	return b.GetBool(ctx, networkManagerPropertyWimaxEnabled)
}

// GetPropertyWimaxHardwareEnabled reports whether WiMAX HW is enabled.
func GetPropertyWimaxHardwareEnabled(ctx context.Context, conn nm.Conn) (bool, error) {
	b := newBase(conn)
	return b.GetBool(ctx, networkManagerPropertyWimaxHardwareEnabled)
}

// GetPropertyActiveConnections returns active connection objects.
func GetPropertyActiveConnections(ctx context.Context, conn nm.Conn) ([]*nm.ActiveConnection, error) {
	b := newBase(conn)
	paths, err := b.GetObjects(ctx, networkManagerPropertyActiveConnections)
	if err != nil {
		return nil, err
	}
	ac := make([]*nm.ActiveConnection, len(paths))
	for i, path := range paths {
		ac[i] = nm.NewActiveConnection(path)
	}
	return ac, nil
}

// GetPropertyPrimaryConnection returns the active connection carrying the
// default route, or nil when there is none.
func GetPropertyPrimaryConnection(ctx context.Context, conn nm.Conn) (*nm.ActiveConnection, error) {
	b := newBase(conn)
	path, ok, err := b.GetOptionalObject(ctx, networkManagerPropertyPrimaryConnection)
	if err != nil || !ok {
		return nil, err
	}
	return nm.NewActiveConnection(path), nil
}

// GetPropertyPrimaryConnectionType returns the type of the primary connection.
func GetPropertyPrimaryConnectionType(ctx context.Context, conn nm.Conn) (string, error) {
	b := newBase(conn)
	return b.GetString(ctx, networkManagerPropertyPrimaryConnectionType)
}

// GetPropertyMetered reports whether the connectivity is metered.
func GetPropertyMetered(ctx context.Context, conn nm.Conn) (nm.Metered, error) {
	b := newBase(conn)
	metered, err := b.GetUint32(ctx, networkManagerPropertyMetered)
	return nm.Metered(metered), err
}

// GetPropertyActivatingConnection returns the connection being activated, or
// nil when none is.
func GetPropertyActivatingConnection(ctx context.Context, conn nm.Conn) (*nm.ActiveConnection, error) {
	b := newBase(conn)
	path, ok, err := b.GetOptionalObject(ctx, networkManagerPropertyActivatingConnection)
	if err != nil || !ok {
		return nil, err
	}
	return nm.NewActiveConnection(path), nil
}

// GetPropertyStartup reports whether NM is still starting up.
func GetPropertyStartup(ctx context.Context, conn nm.Conn) (bool, error) {
	b := newBase(conn)
	return b.GetBool(ctx, networkManagerPropertyStartup)
}

// GetPropertyVersion returns the NetworkManager version string.
func GetPropertyVersion(ctx context.Context, conn nm.Conn) (string, error) {
	b := newBase(conn)
	return b.GetString(ctx, networkManagerPropertyVersion)
}

// GetPropertyCapabilities returns the current set of NM capabilities.
func GetPropertyCapabilities(ctx context.Context, conn nm.Conn) ([]nm.Capability, error) {
	b := newBase(conn)
	raw, err := b.GetUint32s(ctx, networkManagerPropertyCapabilities)
	if err != nil {
		return nil, err
	}
	caps := make([]nm.Capability, len(raw))
	for i, c := range raw {
		caps[i] = nm.Capability(c)
	}
	return caps, nil
}

// GetPropertyState returns the overall state of the NM daemon.
func GetPropertyState(ctx context.Context, conn nm.Conn) (nm.State, error) {
	b := newBase(conn)
	state, err := b.GetUint32(ctx, networkManagerPropertyState)
	return nm.State(state), err
}

// GetPropertyConnectivity returns the result of the last connectivity check.
func GetPropertyConnectivity(ctx context.Context, conn nm.Conn) (nm.Connectivity, error) {
	b := newBase(conn)
	connectivity, err := b.GetUint32(ctx, networkManagerPropertyConnectivity)
	return nm.Connectivity(connectivity), err
}

// GetPropertyConnectivityCheckAvailable reports whether connectivity checking
// has been configured.
func GetPropertyConnectivityCheckAvailable(ctx context.Context, conn nm.Conn) (bool, error) {
	b := newBase(conn)
	return b.GetBool(ctx, networkManagerPropertyConnectivityCheckAvailable)
}

// GetPropertyConnectivityCheckEnabled reports whether connectivity checking
// is enabled.
func GetPropertyConnectivityCheckEnabled(ctx context.Context, conn nm.Conn) (bool, error) {
	b := newBase(conn)
	return b.GetBool(ctx, networkManagerPropertyConnectivityCheckEnabled)
}

// Signals returns an iterator over every signal NetworkManager emits under
// its object namespace, for the cases the typed subscriptions elsewhere in
// these packages do not cover. Iteration ends when ctx is done.
func Signals(ctx context.Context, conn nm.Conn) (iter.Seq[*dbus.Signal], error) {
	return dbusext.Signals(ctx, conn, dbus.WithMatchPathNamespace(dbusext.NetworkManagerObjectPath))
}

// snapshot collects property reads, keeping the values that came back and the
// errors from the reads that did not.
type snapshot struct {
	values map[string]any
	errs   []error
}

// read adds one property to s. A failed read leaves its key absent rather
// than storing a zero value that reads as real state, and stops once ctx is
// done so that a cancelled snapshot does not make twenty more doomed calls.
func read[T any](ctx context.Context, conn nm.Conn, s *snapshot, key string, get func(context.Context, nm.Conn) (T, error)) {
	if ctx.Err() != nil {
		return
	}
	v, err := get(ctx, conn)
	if err != nil {
		s.errs = append(s.errs, fmt.Errorf("%s: %w", key, err))
		return
	}
	s.values[key] = v
}

// Snapshot reads every manager property. A property that cannot be read is
// absent from the map and its error is joined into the returned error, so a
// partial snapshot is still worth printing.
func Snapshot(ctx context.Context, conn nm.Conn) (map[string]any, error) {
	s := &snapshot{values: map[string]any{}}

	read(ctx, conn, s, "Devices", GetPropertyDevices)
	read(ctx, conn, s, "AllDevices", GetPropertyAllDevices)
	read(ctx, conn, s, "Checkpoints", GetPropertyCheckpoints)
	read(ctx, conn, s, "NetworkingEnabled", GetPropertyNetworkingEnabled)
	read(ctx, conn, s, "WirelessEnabled", GetPropertyWirelessEnabled)
	read(ctx, conn, s, "WirelessHardwareEnabled", GetPropertyWirelessHardwareEnabled)
	read(ctx, conn, s, "WwanEnabled", GetPropertyWwanEnabled)
	read(ctx, conn, s, "WwanHardwareEnabled", GetPropertyWwanHardwareEnabled)
	read(ctx, conn, s, "WimaxEnabled", GetPropertyWimaxEnabled)
	read(ctx, conn, s, "WimaxHardwareEnabled", GetPropertyWimaxHardwareEnabled)
	read(ctx, conn, s, "ActiveConnections", GetPropertyActiveConnections)
	read(ctx, conn, s, "PrimaryConnection", GetPropertyPrimaryConnection)
	read(ctx, conn, s, "PrimaryConnectionType", GetPropertyPrimaryConnectionType)
	read(ctx, conn, s, "Metered", GetPropertyMetered)
	read(ctx, conn, s, "ActivatingConnection", GetPropertyActivatingConnection)
	read(ctx, conn, s, "Startup", GetPropertyStartup)
	read(ctx, conn, s, "Version", GetPropertyVersion)
	read(ctx, conn, s, "Capabilities", GetPropertyCapabilities)
	read(ctx, conn, s, "State", GetPropertyState)
	read(ctx, conn, s, "Connectivity", GetPropertyConnectivity)
	read(ctx, conn, s, "ConnectivityCheckAvailable", GetPropertyConnectivityCheckAvailable)
	read(ctx, conn, s, "ConnectivityCheckEnabled", GetPropertyConnectivityCheckEnabled)

	return s.values, errors.Join(s.errs...)
}
