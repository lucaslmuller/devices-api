package database

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
	"github.com/uptrace/bun"
)

type getDatabase struct {
	db *bun.DB
}

func NewGetRepository(db *bun.DB) getDatabase {
	return getDatabase{db}
}

func (r getDatabase) ByID(ctx context.Context, ID string) (device *model.Device, err error) {
	devices := []model.Device{}
	_, err = r.db.NewSelect().Model(&devices).Where("id = ?", ID).Limit(1).Exec(ctx, &devices)

	if len(devices) > 0 {
		device = &devices[0]
	}

	return
}

func (r getDatabase) ByBrand(ctx context.Context, brandID string) (devices []model.Device, err error) {
	_, err = r.db.NewSelect().Model(&devices).Where("brand = ?", brandID).Exec(ctx, &devices)

	return
}

func (r getDatabase) ByState(ctx context.Context, state string) (devices []model.Device, err error) {
	_, err = r.db.NewSelect().Model(&devices).Where("state = ?", state).Exec(ctx, &devices)

	return
}

func (r getDatabase) All(ctx context.Context) (devices []model.Device, err error) {
	_, err = r.db.NewSelect().Model(&devices).Exec(ctx, &devices)

	return
}
