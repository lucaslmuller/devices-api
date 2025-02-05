package endpoint

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
)

type IEndpoint interface {
	Create(ctx context.Context, input *dto.CreateDeviceInput) (output *dto.Output, validationErr error, internalErr error)
	Update(ctx context.Context, input *dto.UpdateDeviceInput) (output *dto.Output, validationErr error, internalErr error)
	Delete(ctx context.Context, input *dto.DeleteDeviceInput) (validationErr error, internalErr error)
	GetByID(ctx context.Context, input *dto.GetDeviceByIDInput) (output *dto.Output, validationErr error, internalErr error)
	GetByState(ctx context.Context, input *dto.GetDeviceByStateInput) (output []dto.Output, validationErr error, internalErr error)
	GetByBrand(ctx context.Context, input *dto.GetDeviceByBrandInput) (output []dto.Output, validationErr error, internalErr error)
	GetAll(ctx context.Context) (output []dto.Output, err error)
}
