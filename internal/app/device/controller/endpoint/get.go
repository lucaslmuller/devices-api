package endpoint

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
)

func (e *DeviceEndpoints) GetByID(ctx context.Context, input *dto.GetDeviceByIDInput) (output *dto.Output, validationErr error, internalErr error) {
	if validationErr = input.Validate(); validationErr != nil {
		return
	}

	result, internalErr := e.service.GetByID(ctx, input.ID)

	if internalErr != nil {
		return
	}

	if result == nil {
		return nil, nil, nil
	}

	output = &dto.Output{}
	output.FromModel(result)
	return

}

func (e *DeviceEndpoints) GetByState(ctx context.Context, input *dto.GetDeviceByStateInput) (output []dto.Output, validationErr error, internalErr error) {
	if validationErr = input.Validate(); validationErr != nil {
		return
	}

	result, internalErr := e.service.GetByState(ctx, input.State)

	if internalErr != nil {
		return
	}

	output = make([]dto.Output, 0)

	for _, item := range result {
		o := &dto.Output{}
		o.FromModel(&item)
		output = append(output, *o)
	}

	return

}
func (e *DeviceEndpoints) GetByBrand(ctx context.Context, input *dto.GetDeviceByBrandInput) (output []dto.Output, validationErr error, internalErr error) {
	if validationErr = input.Validate(); validationErr != nil {
		return
	}

	result, internalErr := e.service.GetByBrand(ctx, input.Brand)

	if internalErr != nil {
		return
	}

	output = make([]dto.Output, 0)

	for _, item := range result {
		o := &dto.Output{}
		o.FromModel(&item)
		output = append(output, *o)
	}

	return

}
func (e *DeviceEndpoints) GetAll(ctx context.Context) (output []dto.Output, err error) {
	result, err := e.service.GetAll(ctx)

	if err != nil {
		return
	}

	output = make([]dto.Output, 0)

	for _, item := range result {
		o := &dto.Output{}
		o.FromModel(&item)
		output = append(output, *o)
	}

	return output, err

}
