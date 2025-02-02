package endpoint

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
)

func (e *DeviceEndpoints) Update(ctx context.Context, input *dto.UpdateDeviceInput) (output *dto.Output, validationErr error, internalErr error) {
	if validationErr = input.Validate(); validationErr != nil {
		return
	}

	updatedDevice, internalErr := e.service.Update(ctx, input.ToModel())

	if internalErr != nil {
		return
	}

	output = &dto.Output{}
	output.FromModel(updatedDevice)
	return
}
