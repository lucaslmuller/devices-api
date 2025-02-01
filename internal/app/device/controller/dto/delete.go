package dto

import (
	"errors"

	"github.com/google/uuid"
)

type DeleteDeviceInput struct {
	ID string
}

func (i *DeleteDeviceInput) Validate() error {
	if i.ID == "" {
		return errors.New("id is required")
	}

	err := uuid.Validate(i.ID)
	if err != nil {
		return errors.New("invalid id - must be a valid uuid")
	}

	return nil
}
