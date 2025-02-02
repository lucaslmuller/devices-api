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

	if err == nil && cached != nil {
		device = *cached
		return
	}

	device, err = s.repository.GetByID(ctx, ID)

	if device != nil {
		s.redis.Set(ctx, device.CacheKey().GetByID(), device)
	}

	return
}

func (s *DeviceService) GetByBrand(ctx context.Context, brandID string) (devices []model.Device, err error) {
	dummy := model.Device{
		Brand: brandID,
	}

	cached, err := cache.Get[[]model.Device](ctx, dummy.CacheKey().GetByBrand(), s.redis)

	if err == nil && cached != nil {
		devices = *cached
		return
	}

	devices, err = s.repository.GetByBrand(ctx, brandID)

	s.redis.Set(ctx, dummy.CacheKey().GetByBrand(), devices)

	return
}

func (s *DeviceService) GetByState(ctx context.Context, state string) (devices []model.Device, err error) {
	dummy := model.Device{
		State: state,
	}

	cached, err := cache.Get[[]model.Device](ctx, dummy.CacheKey().GetByState(), s.redis)

	if err == nil && cached != nil {
		devices = *cached
		return
	}

	devices, err = s.repository.GetByState(ctx, state)

	s.redis.Set(ctx, dummy.CacheKey().GetByState(), devices)

	return
}

func (s *DeviceService) GetAll(ctx context.Context) (devices []model.Device, err error) {
	dummy := model.Device{}
	cached, err := cache.Get[[]model.Device](ctx, dummy.CacheKey().GetAll(), s.redis)

	if err == nil && cached != nil {
		devices = *cached
		return
	}

	devices, err = s.repository.GetAll(ctx)

	s.redis.Set(ctx, dummy.CacheKey().GetAll(), devices)

	return
}
