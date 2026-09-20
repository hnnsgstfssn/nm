package device

import (
	"context"

	"github.com/hnnsgstfssn/nm"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	deviceGenericPropertyTypeDescription = dbusext.DeviceGenericInterface + ".TypeDescription"
)

func GenericGetPropertyTypeDescription(ctx context.Context, conn nm.Conn, dev *nm.Device) (string, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetString(ctx, deviceGenericPropertyTypeDescription)
}
