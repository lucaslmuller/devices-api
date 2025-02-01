package database

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
	"github.com/uptrace/bun"
)

type updateDatabase struct {
	db *bun.DB
}

func NewUpdateRepository(db *bun.DB) updateDatabase {
	return updateDatabase{db}
}

func (r updateDatabase) Device(ctx context.Context, device *model.Device) (err error) {
	_, err = r.db.NewUpdate().Model(device).WherePK().Exec(ctx)

	if err != nil {
		return
	}
	return
}
