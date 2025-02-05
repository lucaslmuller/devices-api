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

	if err = s.repository.Update(ctx, device); err != nil {
		return nil, err
	}

	updatedDevice = device

	if err = s.handleUpdateCache(ctx, currentDevice, updatedDevice); err != nil {
		s.logger.Error(err)
	}

	return updatedDevice, nil
}

func (s *DeviceService) handleUpdateCache(ctx context.Context, currentDevice *model.Device, updatedDevice *model.Device) (err error) {
	if err = s.redis.Delete(ctx, currentDevice.CacheKey().GetByBrand()); err != nil {
		return err
	}

	if err = s.redis.Delete(ctx, currentDevice.CacheKey().GetByState()); err != nil {
		return err
	}

	if err = s.redis.Delete(ctx, updatedDevice.CacheKey().GetByBrand()); err != nil {
		return err
	}

	if err = s.redis.Delete(ctx, updatedDevice.CacheKey().GetAll()); err != nil {
		return err
	}

	if err = s.redis.Delete(ctx, updatedDevice.CacheKey().GetByState()); err != nil {
		return err
	}

	if err = s.redis.Delete(ctx, updatedDevice.CacheKey().GetByID()); err != nil {
		return err
	}

	if err = s.redis.Set(ctx, updatedDevice.CacheKey().GetByID(), updatedDevice); err != nil {
		return err
	}

	return nil
}
