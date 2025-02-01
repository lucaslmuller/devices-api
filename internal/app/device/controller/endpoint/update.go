package endpoint

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain"
)

type update struct {
	service *domain.DeviceServices
}

func NewUpdateEndpoints(service *domain.DeviceServices) update {
	return update{service}
}

func (u update) Device(ctx context.Context, input *dto.UpdateDeviceInput) (output *dto.Output, validationErr error, internalErr error) {
	if validationErr = input.Validate(); validationErr != nil {
		return
	}

	updatedDevice, internalErr := u.service.Update.Device(ctx, input.ToModel())

	if internalErr != nil {
		return
	}

	output = &dto.Output{}
	output.FromModel(updatedDevice)
	return
}
