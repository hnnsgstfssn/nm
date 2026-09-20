package nm

import (
	"context"

	"github.com/godbus/dbus/v5"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	checkpointPropertyDevices         = dbusext.CheckpointInterface + ".Devices"
	checkpointPropertyCreated         = dbusext.CheckpointInterface + ".Created"
	checkpointPropertyRollbackTimeout = dbusext.CheckpointInterface + ".RollbackTimeout"
)

// Checkpoint is an eagerly-populated snapshot of a NetworkManager checkpoint.
type Checkpoint struct {
	Path            dbus.ObjectPath
	Devices         []*Device
	Created         int64
	RollbackTimeout uint32
}

// NewCheckpoint fetches all checkpoint properties from D-Bus and returns a
// populated snapshot.
func NewCheckpoint(ctx context.Context, conn Conn, objectPath dbus.ObjectPath) (*Checkpoint, error) {
	b := dbusext.NewBase(conn, objectPath)

	devicesPaths, err := b.GetObjects(ctx, checkpointPropertyDevices)
	if err != nil {
		return nil, err
	}
	devices := make([]*Device, len(devicesPaths))
	for i, path := range devicesPaths {
		devices[i] = NewDevice(path)
	}

	created, err := b.GetInt64(ctx, checkpointPropertyCreated)
	if err != nil {
		return nil, err
	}
	rollbackTimeout, err := b.GetUint32(ctx, checkpointPropertyRollbackTimeout)
	if err != nil {
		return nil, err
	}

	return &Checkpoint{
		Path:            objectPath,
		Devices:         devices,
		Created:         created,
		RollbackTimeout: rollbackTimeout,
	}, nil
}
