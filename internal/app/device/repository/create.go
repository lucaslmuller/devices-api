package repository

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

func (r *DeviceRepository) Create(ctx context.Context, device *model.Device) (err error) {
	_, err = r.db.NewInsert().Model(device).Exec(ctx)

	if err != nil {
		return
	}
	return
}
