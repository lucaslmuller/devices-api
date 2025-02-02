package repository

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

func (r *DeviceRepository) Update(ctx context.Context, device *model.Device) (err error) {
	_, err = r.db.NewUpdate().Model(device).WherePK().Exec(ctx)

	if err != nil {
		return
	}
	return
}
