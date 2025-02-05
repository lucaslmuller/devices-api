package service

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

func (s *DeviceService) Create(ctx context.Context, device *model.Device) (newDevice *model.Device, err error) {
	err = s.repository.Create(ctx, device)
	if err != nil {
		return nil, err
	}

	if err = s.handleCreateCache(ctx, device); err != nil {
		s.logger.Error(err)
	}

	return device, nil
}

func (s *DeviceService) handleCreateCache(ctx context.Context, device *model.Device) (err error) {
	if err = s.redis.Delete(ctx, device.CacheKey().GetByBrand()); err != nil {
		return err
	}

	if err = s.redis.Delete(ctx, device.CacheKey().GetByState()); err != nil {
		return err
	}

	if err = s.redis.Delete(ctx, device.CacheKey().GetAll()); err != nil {
		return err
	}

	if err = s.redis.Set(ctx, device.CacheKey().GetByID(), device); err != nil {
		return err
	}

	return nil
}
