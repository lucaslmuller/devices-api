package domain

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

type DeviceServices struct {
	Get    IGetService
	Create ICreateService
	Update IUpdateService
	Delete IDeleteService
}

type IGetService interface {
	ByID(ctx context.Context, ID string) (devices *model.Device, err error)
	ByBrand(ctx context.Context, brandID string) (devices []model.Device, err error)
	ByState(ctx context.Context, state string) (devices []model.Device, err error)
	All(ctx context.Context) (devices []model.Device, err error)
}

type ICreateService interface {
	Device(ctx context.Context, device *model.Device) (newDevice *model.Device, err error)
}

type IUpdateService interface {
	Device(ctx context.Context, device *model.Device) (updateDevice *model.Device, err error)
}

type IDeleteService interface {
	Device(ctx context.Context, ID string) (err error)
}
