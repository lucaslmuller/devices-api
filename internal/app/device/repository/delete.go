package repository

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

func (r *DeviceRepository) Delete(ctx context.Context, ID string) (err error) {
	f := &model.Device{}
	r.db.NewDelete().Model(f).Where("id = ?", ID).Exec(ctx)
	return nil
}
