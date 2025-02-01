package endpoint

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain"
)

type create struct {
	service *domain.DeviceServices
}

func NewCreationEndpoints(service *domain.DeviceServices) create {
	return create{service}
}

func (u create) Device(ctx context.Context, input *dto.CreateDeviceInput) (output *dto.Output, validationErr error, internalErr error) {
	if validationErr = input.Validate(); validationErr != nil {
		return
	}

	result, internalErr := u.service.Create.Device(ctx, input.ToModel())

	if internalErr != nil {
		return
	}

	output = &dto.Output{}
	output.FromModel(result)

	return
}
