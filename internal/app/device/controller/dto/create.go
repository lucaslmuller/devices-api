package dto

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

type CreateDeviceInput struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
	State string `json:"state"`
}

func (i *CreateDeviceInput) Validate() error {
	if i.Name == "" {
		return errors.New("name is required")
	}

	if i.Brand == "" {
		return errors.New("brand is required")
	}

	if i.State == "" {
		return errors.New("state is required")
	}

	if !ValidStates[i.State] {
		return errors.New("invalid state, value must be one of the following: 'active', 'inactive' or 'in-use'")
	}

	return nil
}

func (i *CreateDeviceInput) ToModel() (device *model.Device) {
	now := time.Now()
	id := uuid.NewString()
	device = &model.Device{
		Id:           &id,
		Name:         i.Name,
		Brand:        i.Brand,
		State:        i.State,
		CreationTime: &now,
	}

	return
}

type CreateDeviceOutput struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Brand        string    `json:"brand"`
	State        string    `json:"state"`
	CreationTime time.Time `json:"creationTime"`
}

func (o *CreateDeviceOutput) FromModel(m *model.Device) {
	o.Id = *m.Id
	o.Name = m.Name
	o.Brand = m.Brand
	o.State = m.State
	o.CreationTime = *m.CreationTime
}
