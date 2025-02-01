package database

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
	"github.com/uptrace/bun"
)

type deleteDatabase struct {
	db *bun.DB
}

func NewDeleteRepository(db *bun.DB) deleteDatabase {
	return deleteDatabase{db}
}

func (r deleteDatabase) ByID(ctx context.Context, ID string) (err error) {
	f := &model.Device{}
	r.db.NewDelete().Model(f).Where("id = ?", ID).Exec(ctx)
	return nil
}
