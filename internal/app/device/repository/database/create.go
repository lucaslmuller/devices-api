package database

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
	"github.com/uptrace/bun"
)

type createDatabase struct {
	db *bun.DB
}

func NewCreateRepository(db *bun.DB) createDatabase {
	return createDatabase{db}
}

func (r createDatabase) NewDevice(ctx context.Context, device *model.Device) (err error) {
	_, err = r.db.NewInsert().Model(device).Exec(ctx)

	if err != nil {
		return
	}
	return
}
