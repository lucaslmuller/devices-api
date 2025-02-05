package repository

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

type IRepository interface {
	Create(ctx context.Context, device *model.Device) (err error)
	GetByID(ctx context.Context, ID string) (device *model.Device, err error)
	GetAll(ctx context.Context) (devices []model.Device, err error)
	GetByState(ctx context.Context, state string) (devices []model.Device, err error)
	GetByBrand(ctx context.Context, brandID string) (devices []model.Device, err error)
	Update(ctx context.Context, device *model.Device) (err error)
	Delete(ctx context.Context, ID string) (err error)
}
