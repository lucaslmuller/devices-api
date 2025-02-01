package model

import (
	"time"

	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis/cache"
)

type Device struct {
	Id           *string    `bun:"id,type:uuid,pk"`
	Name         string     `bun:"name"`
	Brand        string     `bun:"brand"`
	State        string     `bun:"state"`
	CreationTime *time.Time `bun:"creation_time"`
}

// Read implements io.Reader.
func (m *Device) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}

func (m *Device) Merge(updated *Device) (device *Device) {
	device = &Device{
		Id:           m.Id,
		Name:         m.Name,
		Brand:        m.Brand,
		State:        m.State,
		CreationTime: m.CreationTime,
	}

	if updated.Name != "" {
		device.Name = updated.Name
	}

	if updated.Brand != "" {
		device.Brand = updated.Brand
	}

	if updated.State != "" {
		device.State = updated.State
	}

	return
}

type CacheKey struct {
	Device *Device
}

func (m *Device) CacheKey() *CacheKey {
	return &CacheKey{
		Device: m,
	}
}

func (m *CacheKey) GetAll() string {
	return cache.NewKey("devices").
		ToString()
}

func (m *CacheKey) GetByID() string {
	return cache.NewKey("devices").
		AddComponent("id", *m.Device.Id).
		ToString()
}

func (m *CacheKey) GetByState() string {
	return cache.NewKey("devices").
		AddComponent("state", m.Device.State).
		ToString()
}

func (m *CacheKey) GetByBrand() string {
	return cache.NewKey("devices").
		AddComponent("brand", m.Device.Brand).
		ToString()
}
