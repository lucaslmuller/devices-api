package repository

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

type Device struct {
	Database DatabaseRepository
}

type DatabaseRepository struct {
	Get    IDatabaseGetRepository
	Create IDatabaseCreateRepository
	Update IDatabaseUpdateRepository
	Delete IDatabaseDeleteRepository
}

type IDatabaseGetRepository interface {
	ByID(ctx context.Context, ID string) (device *model.Device, err error)
	ByBrand(ctx context.Context, brandID string) (devices []model.Device, err error)
	ByState(ctx context.Context, state string) (devices []model.Device, err error)
	All(ctx context.Context) (devices []model.Device, err error)
}
type IDatabaseCreateRepository interface {
	NewDevice(ctx context.Context, device *model.Device) (err error)
}
type IDatabaseUpdateRepository interface {
	Device(ctx context.Context, device *model.Device) (err error)
}
type IDatabaseDeleteRepository interface {
	ByID(ctx context.Context, ID string) (err error)
}
