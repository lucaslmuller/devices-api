package endpoint

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
)

func (e *DeviceEndpoints) Create(ctx context.Context, input *dto.CreateDeviceInput) (output *dto.Output, validationErr error, internalErr error) {
	if validationErr = input.Validate(); validationErr != nil {
		return
	}

	result, internalErr := e.service.Create(ctx, input.ToModel())

	if internalErr != nil {
		return
	}

	output = &dto.Output{}
	output.FromModel(result)

	return
}
