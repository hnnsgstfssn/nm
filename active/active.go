package active

import (
	"context"
	"iter"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	activeConnectionPropertyConnection     = dbusext.ActiveConnectionInterface + ".Connection"
	activeConnectionPropertySpecificObject = dbusext.ActiveConnectionInterface + ".SpecificObject"
	activeConnectionPropertyID             = dbusext.ActiveConnectionInterface + ".Id"
	activeConnectionPropertyUUID           = dbusext.ActiveConnectionInterface + ".Uuid"
	activeConnectionPropertyType           = dbusext.ActiveConnectionInterface + ".Type"
	activeConnectionPropertyDevices        = dbusext.ActiveConnectionInterface + ".Devices"
	activeConnectionPropertyState          = dbusext.ActiveConnectionInterface + ".State"
	activeConnectionPropertyStateFlags     = dbusext.ActiveConnectionInterface + ".StateFlags"
	activeConnectionPropertyDefault        = dbusext.ActiveConnectionInterface + ".Default"
	activeConnectionPropertyIP4Config      = dbusext.ActiveConnectionInterface + ".Ip4Config"
	activeConnectionPropertyDhcp4Config    = dbusext.ActiveConnectionInterface + ".Dhcp4Config"
	activeConnectionPropertyDefault6       = dbusext.ActiveConnectionInterface + ".Default6"
	activeConnectionPropertyIP6Config      = dbusext.ActiveConnectionInterface + ".Ip6Config"
	activeConnectionPropertyDhcp6Config    = dbusext.ActiveConnectionInterface + ".Dhcp6Config"
	activeConnectionPropertyVpn            = dbusext.ActiveConnectionInterface + ".Vpn"
	activeConnectionPropertyMaster         = dbusext.ActiveConnectionInterface + ".Master"

	activeConnectionSignalStateChanged = "StateChanged"
)

func GetPropertyConnection(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) (*nm.Connection, error) {
	b := dbusext.NewBase(conn, ac.Path)
	path, err := b.GetObject(ctx, activeConnectionPropertyConnection)
	if err != nil {
		return nil, err
	}
	return nm.NewConnection(path), nil
}

// GetPropertySpecificObject returns the access point this connection was
// activated against, or nil for a connection that has none.
func GetPropertySpecificObject(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) (*nm.AccessPoint, error) {
	b := dbusext.NewBase(conn, ac.Path)
	path, ok, err := b.GetOptionalObject(ctx, activeConnectionPropertySpecificObject)
	if err != nil || !ok {
		return nil, err
	}
	return nm.NewAccessPoint(ctx, conn, path)
}

func GetPropertyID(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) (string, error) {
	b := dbusext.NewBase(conn, ac.Path)
	return b.GetString(ctx, activeConnectionPropertyID)
}

func GetPropertyUUID(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) (string, error) {
	b := dbusext.NewBase(conn, ac.Path)
	return b.GetString(ctx, activeConnectionPropertyUUID)
}

func GetPropertyType(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) (string, error) {
	b := dbusext.NewBase(conn, ac.Path)
	return b.GetString(ctx, activeConnectionPropertyType)
}

func GetPropertyDevices(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) ([]*nm.Device, error) {
	b := dbusext.NewBase(conn, ac.Path)
	paths, err := b.GetObjects(ctx, activeConnectionPropertyDevices)
	if err != nil {
		return nil, err
	}
	devices := make([]*nm.Device, len(paths))
	for i, path := range paths {
		devices[i] = nm.NewDevice(path)
	}
	return devices, nil
}

func GetPropertyState(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) (nm.ActiveConnectionState, error) {
	b := dbusext.NewBase(conn, ac.Path)
	state, err := b.GetUint32(ctx, activeConnectionPropertyState)
	return nm.ActiveConnectionState(state), err
}

func GetPropertyStateFlags(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) (nm.ActivationStateFlag, error) {
	b := dbusext.NewBase(conn, ac.Path)
	flags, err := b.GetUint32(ctx, activeConnectionPropertyStateFlags)
	return nm.ActivationStateFlag(flags), err
}

func GetPropertyDefault(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) (bool, error) {
	b := dbusext.NewBase(conn, ac.Path)
	return b.GetBool(ctx, activeConnectionPropertyDefault)
}

func GetPropertyIP4Config(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) (*nm.IP4Config, error) {
	b := dbusext.NewBase(conn, ac.Path)
	path, ok, err := b.GetOptionalObject(ctx, activeConnectionPropertyIP4Config)
	if err != nil || !ok {
		return nil, err
	}
	return nm.NewIP4Config(ctx, conn, path)
}

func GetPropertyDHCP4Config(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) (*nm.DHCP4Config, error) {
	b := dbusext.NewBase(conn, ac.Path)
	path, ok, err := b.GetOptionalObject(ctx, activeConnectionPropertyDhcp4Config)
	if err != nil || !ok {
		return nil, err
	}
	return nm.NewDHCP4Config(ctx, conn, path)
}

func GetPropertyDefault6(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) (bool, error) {
	b := dbusext.NewBase(conn, ac.Path)
	return b.GetBool(ctx, activeConnectionPropertyDefault6)
}

func GetPropertyIP6Config(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) (*nm.IP6Config, error) {
	b := dbusext.NewBase(conn, ac.Path)
	path, ok, err := b.GetOptionalObject(ctx, activeConnectionPropertyIP6Config)
	if err != nil || !ok {
		return nil, err
	}
	return nm.NewIP6Config(ctx, conn, path)
}

func GetPropertyDHCP6Config(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) (*nm.DHCP6Config, error) {
	b := dbusext.NewBase(conn, ac.Path)
	path, ok, err := b.GetOptionalObject(ctx, activeConnectionPropertyDhcp6Config)
	if err != nil || !ok {
		return nil, err
	}
	return nm.NewDHCP6Config(ctx, conn, path)
}

func GetPropertyVPN(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) (bool, error) {
	b := dbusext.NewBase(conn, ac.Path)
	return b.GetBool(ctx, activeConnectionPropertyVpn)
}

func GetPropertyMaster(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) (*nm.Device, error) {
	b := dbusext.NewBase(conn, ac.Path)
	path, ok, err := b.GetOptionalObject(ctx, activeConnectionPropertyMaster)
	if err != nil || !ok {
		return nil, err
	}
	return nm.NewDevice(path), nil
}

// StateChanges returns an iterator over ac's activation state transitions.
// Iteration ends when ctx is done.
//
// The subscription is in place before this returns, so a caller can read the
// current state afterwards without a transition slipping through the gap.
// The match rule is removed when the range loop ends or when ctx is done,
// whichever comes first.
func StateChanges(ctx context.Context, conn nm.Conn, ac *nm.ActiveConnection) (iter.Seq[nm.StateChange], error) {
	signals, err := dbusext.Signals(ctx, conn,
		dbus.WithMatchInterface(dbusext.ActiveConnectionInterface),
		dbus.WithMatchMember(activeConnectionSignalStateChanged),
		dbus.WithMatchObjectPath(ac.Path),
	)
	if err != nil {
		return nil, err
	}

	return func(yield func(nm.StateChange) bool) {
		for signal := range signals {
			change, ok := stateChangeFrom(signal)
			if !ok {
				continue
			}
			if !yield(change) {
				return
			}
		}
	}, nil
}

// stateChangeFrom reads a StateChanged signal, reporting whether it was one.
// The body is (state, reason), and its length is checked because any process
// on the bus can emit a signal with this name.
func stateChangeFrom(signal *dbus.Signal) (nm.StateChange, bool) {
	if signal.Name != dbusext.ActiveConnectionInterface+"."+activeConnectionSignalStateChanged || len(signal.Body) < 2 {
		return nm.StateChange{}, false
	}
	state, ok1 := signal.Body[0].(uint32)
	reason, ok2 := signal.Body[1].(uint32)
	if !ok1 || !ok2 {
		return nm.StateChange{}, false
	}
	return nm.StateChange{
		Path:   signal.Path,
		State:  nm.ActiveConnectionState(state),
		Reason: nm.ActiveConnectionStateReason(reason),
	}, true
}
