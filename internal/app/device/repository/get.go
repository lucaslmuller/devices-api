package repository

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

func (r *DeviceRepository) GetByID(ctx context.Context, ID string) (device *model.Device, err error) {
	devices := []model.Device{}
	_, err = r.db.NewSelect().Model(&devices).Where("id = ?", ID).Limit(1).Exec(ctx, &devices)

	if len(devices) > 0 {
		device = &devices[0]
	}

	return
}

func (r *DeviceRepository) GetByBrand(ctx context.Context, brandID string) (devices []model.Device, err error) {
	_, err = r.db.NewSelect().Model(&devices).Where("brand = ?", brandID).Exec(ctx, &devices)

	return
}

func (r *DeviceRepository) GetByState(ctx context.Context, state string) (devices []model.Device, err error) {
	_, err = r.db.NewSelect().Model(&devices).Where("state = ?", state).Exec(ctx, &devices)

	return
}

func (r *DeviceRepository) GetAll(ctx context.Context) (devices []model.Device, err error) {
	_, err = r.db.NewSelect().Model(&devices).Exec(ctx, &devices)

	return
}
