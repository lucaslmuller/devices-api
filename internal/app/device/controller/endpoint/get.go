package endpoint

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain"
)

type get struct {
	service *domain.DeviceServices
}

func NewGetEndpoints(service *domain.DeviceServices) get {
	return get{service}
}

func (u get) ByID(ctx context.Context, input *dto.GetDeviceByIDInput) (output *dto.Output, validationErr error, internalErr error) {
	if validationErr = input.Validate(); validationErr != nil {
		return
	}

	result, internalErr := u.service.Get.ByID(ctx, input.ID)

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

func (u get) ByState(ctx context.Context, input *dto.GetDeviceByStateInput) (output []dto.Output, validationErr error, internalErr error) {
	if validationErr = input.Validate(); validationErr != nil {
		return
	}

	result, internalErr := u.service.Get.ByState(ctx, input.State)

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
func (u get) ByBrand(ctx context.Context, input *dto.GetDeviceByBrandInput) (output []dto.Output, validationErr error, internalErr error) {
	if validationErr = input.Validate(); validationErr != nil {
		return
	}

	result, internalErr := u.service.Get.ByBrand(ctx, input.Brand)

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
func (u get) All(ctx context.Context) (output []dto.Output, err error) {
	result, err := u.service.Get.All(ctx)

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
