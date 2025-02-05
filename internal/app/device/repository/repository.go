package repository

import (
	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
	"github.com/lucaslmuller/technical-test/internal/infrastructure"
	"github.com/uptrace/bun"
)

type DeviceRepository struct {
	db *bun.DB
}

func NewRepository(res *infrastructure.Resources) *DeviceRepository {
	CreateTable(res)

	return &DeviceRepository{db: res.PostgreSQL}
}

func CreateTable(res *infrastructure.Resources) {
	_, err := res.PostgreSQL.NewCreateTable().IfNotExists().Model((*model.Device)(nil)).Exec(res.Ctx)
	if err != nil {
		res.Logger.Fatal("error creating favorite table")
	}
}
