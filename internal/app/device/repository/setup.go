package repository

import (
	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
	"github.com/lucaslmuller/technical-test/internal/app/device/repository/database"
	"github.com/lucaslmuller/technical-test/internal/infrastructure"
)

func Setup(res *infrastructure.Resources) *Device {
	CreateTable(res)

	return &Device{
		Database: DatabaseRepository{
			Get:    database.NewGetRepository(res.PostgreSQL),
			Create: database.NewCreateRepository(res.PostgreSQL),
			Update: database.NewUpdateRepository(res.PostgreSQL),
			Delete: database.NewDeleteRepository(res.PostgreSQL),
		},
	}
}

func CreateTable(res *infrastructure.Resources) {
	_, err := res.PostgreSQL.NewCreateTable().IfNotExists().Model((*model.Device)(nil)).Exec(res.Ctx)
	if err != nil {
		res.Logger.Fatal("error creating favorite table")
	}
}
