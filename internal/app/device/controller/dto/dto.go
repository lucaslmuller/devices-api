package dto

import (
	"time"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

var ValidStates = map[string]bool{
	"available": true,
	"in-use":    true,
	"inactive":  true,
}

type Output struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Brand        string    `json:"brand"`
	State        string    `json:"state"`
	CreationTime time.Time `json:"creationTime"`
}

func (o *Output) FromModel(m *model.Device) {
	if m == nil {
		return
	}

	o.Id = *m.Id
	o.Name = m.Name
	o.Brand = m.Brand
	o.State = m.State
	o.CreationTime = *m.CreationTime
}
