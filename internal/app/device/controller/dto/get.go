package dto

import (
	"errors"

	"github.com/google/uuid"
)

type GetDeviceByIDInput struct {
	ID string
}

func (i *GetDeviceByIDInput) Validate() error {
	if i.ID == "" {
		return errors.New("id is required")
	}

	err := uuid.Validate(i.ID)
	if err != nil {
		return errors.New("invalid id - must be a valid uuid")
	}
	return nil
}

type GetDeviceByBrandInput struct {
	Brand string
}

func (i *GetDeviceByBrandInput) Validate() error {
	if i.Brand == "" {
		return errors.New("brand is required")
	}
	return nil
}

type GetDeviceByStateInput struct {
	State string
}

func (i *GetDeviceByStateInput) Validate() error {
	if i.State == "" {
		return errors.New("state is required")
	}

	if !ValidStates[i.State] {
		return errors.New("invalid state, value must be one of the following: 'active', 'inactive' or 'in-use'")
	}
	return nil
}
