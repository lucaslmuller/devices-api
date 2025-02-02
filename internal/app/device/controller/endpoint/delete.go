package endpoint

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
)

func (e DeviceEndpoints) Delete(ctx context.Context, input *dto.DeleteDeviceInput) (validationErr error, internalErr error) {
	if validationErr = input.Validate(); validationErr != nil {
		return
	}

	internalErr = e.service.Delete(ctx, input.ID)

	if internalErr != nil {
		return
	}

	return
}
