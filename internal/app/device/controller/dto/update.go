package dto

import (
	"errors"

	"github.com/google/uuid"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

type UpdateDeviceInput struct {
	ID    string
	Name  *string `json:"name"`
	Brand *string `json:"brand"`
	State *string `json:"state"`
}

func (i *UpdateDeviceInput) Validate() error {
	if i.ID == "" {
		return errors.New("id is required")
	}

	err := uuid.Validate(i.ID)
	if err != nil {
		return errors.New("invalid id - must be a valid uuid")
	}

	if i.State != nil && !ValidStates[*i.State] {
		return errors.New("invalid state, value must be one of the following: 'active', 'inactive' or 'in-use'")
	}

	return nil
}

func (i *UpdateDeviceInput) ToModel() (device *model.Device) {
	device = &model.Device{
		Id: &i.ID,
	}

	if i.Name != nil {
		device.Name = *i.Name
	}
	if i.Brand != nil {
		device.Brand = *i.Brand
	}
	if i.State != nil {
		device.State = *i.State
	}

	return
}
