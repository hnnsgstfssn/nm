package device

import (
	"context"

	"github.com/hnnsgstfssn/nm"
	"github.com/hnnsgstfssn/nm/internal/dbusext"
)

const (
	deviceStatisticsPropertyRefreshRateMs = dbusext.DeviceStatisticsInterface + ".RefreshRateMs"
	deviceStatisticsPropertyTxBytes       = dbusext.DeviceStatisticsInterface + ".TxBytes"
	deviceStatisticsPropertyRxBytes       = dbusext.DeviceStatisticsInterface + ".RxBytes"
)

func StatisticsGetPropertyRefreshRateMs(ctx context.Context, conn nm.Conn, dev *nm.Device) (uint32, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetUint32(ctx, deviceStatisticsPropertyRefreshRateMs)
}

func StatisticsSetPropertyRefreshRateMs(ctx context.Context, conn nm.Conn, dev *nm.Device, rate uint32) error {
	b := dbusext.NewBase(conn, dev.Path)
	return b.SetProperty(ctx, deviceStatisticsPropertyRefreshRateMs, rate)
}

func StatisticsGetPropertyTxBytes(ctx context.Context, conn nm.Conn, dev *nm.Device) (uint64, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetUint64(ctx, deviceStatisticsPropertyTxBytes)
}

func StatisticsGetPropertyRxBytes(ctx context.Context, conn nm.Conn, dev *nm.Device) (uint64, error) {
	b := dbusext.NewBase(conn, dev.Path)
	return b.GetUint64(ctx, deviceStatisticsPropertyRxBytes)
}
