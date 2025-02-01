package endpoint

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain"
)

type delete struct {
	service *domain.DeviceServices
}

func NewDeletionEndpoints(service *domain.DeviceServices) delete {
	return delete{service}
}

func (u delete) Device(ctx context.Context, input *dto.DeleteDeviceInput) (validationErr error, internalErr error) {
	if validationErr = input.Validate(); validationErr != nil {
		return
	}

	internalErr = u.service.Delete.Device(ctx, input.ID)

	if internalErr != nil {
		return
	}

	return
}
