package service

import (
	"context"
	"errors"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

func (s *DeviceService) Update(ctx context.Context, device *model.Device) (updatedDevice *model.Device, err error) {
	currentDevice, err := s.repository.GetByID(ctx, *device.Id)

	if err != nil {
		return nil, err
	}

	if currentDevice == nil {
		return nil, errors.New("device not found")
	}

	device = currentDevice.Merge(device)

	if currentDevice.State == "in-use" {
		if device.Brand != currentDevice.Brand {
			return nil, errors.New("cannot change brand while in use")
		}

		if device.Name != currentDevice.Name {
			return nil, errors.New("cannot change name while in use")
		}
	}

	err = s.repository.Update(ctx, device)
	updatedDevice = device

	s.redis.Delete(ctx, currentDevice.CacheKey().GetByBrand())
	s.redis.Delete(ctx, currentDevice.CacheKey().GetByState())
	s.redis.Delete(ctx, device.CacheKey().GetByBrand())
	s.redis.Delete(ctx, device.CacheKey().GetAll())
	s.redis.Delete(ctx, device.CacheKey().GetByState())
	s.redis.Delete(ctx, device.CacheKey().GetByID())

	s.redis.Set(ctx, updatedDevice.CacheKey().GetByID(), updatedDevice)

	return
}
