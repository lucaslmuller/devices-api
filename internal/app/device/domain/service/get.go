package service

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis/cache"
)

func (s *DeviceService) GetByID(ctx context.Context, ID string) (device *model.Device, err error) {
	dummy := model.Device{
		Id: &ID,
	}
	cached, err := cache.Get[*model.Device](ctx, dummy.CacheKey().GetByID(), s.redis)

	if err != nil {
		s.logger.Error(err)
	}

	if err == nil && cached != nil {
		device = *cached
		return
	}

	device, err = s.repository.GetByID(ctx, ID)

	if err != nil {
		return nil, err
	}

	if device != nil {
		if err = s.redis.Set(ctx, device.CacheKey().GetByID(), device); err != nil {
			s.logger.Error(err)
		}
	}

	return device, nil
}

func (s *DeviceService) GetByBrand(ctx context.Context, brandID string) (devices []model.Device, err error) {
	dummy := model.Device{
		Brand: brandID,
	}

	cached, err := cache.Get[[]model.Device](ctx, dummy.CacheKey().GetByBrand(), s.redis)

	if err != nil {
		s.logger.Error(err)
	}

	if err == nil && cached != nil {
		devices = *cached
		return devices, nil
	}

	devices, err = s.repository.GetByBrand(ctx, brandID)

	if err != nil {
		return nil, err
	}

	if err = s.redis.Set(ctx, dummy.CacheKey().GetByBrand(), devices); err != nil {
		s.logger.Error(err)
	}

	return devices, nil
}

func (s *DeviceService) GetByState(ctx context.Context, state string) (devices []model.Device, err error) {
	dummy := model.Device{
		State: state,
	}

	cached, err := cache.Get[[]model.Device](ctx, dummy.CacheKey().GetByState(), s.redis)

	if err != nil {
		s.logger.Error(err)
	}

	if err == nil && cached != nil {
		devices = *cached
		return devices, nil
	}

	devices, err = s.repository.GetByState(ctx, state)

	if err != nil {
		return nil, err
	}

	if err = s.redis.Set(ctx, dummy.CacheKey().GetByState(), devices); err != nil {
		s.logger.Error(err)
	}

	return devices, nil
}

func (s *DeviceService) GetAll(ctx context.Context) (devices []model.Device, err error) {
	dummy := model.Device{}
	cached, err := cache.Get[[]model.Device](ctx, dummy.CacheKey().GetAll(), s.redis)

	if err != nil {
		s.logger.Error(err)
	}

	if err == nil && cached != nil {
		devices = *cached
		return
	}

	devices, err = s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	if err = s.redis.Set(ctx, dummy.CacheKey().GetAll(), devices); err != nil {
		s.logger.Error(err)
	}

	return devices, nil
}
