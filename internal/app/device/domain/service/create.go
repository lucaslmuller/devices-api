package service

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

func (s *DeviceService) Create(ctx context.Context, device *model.Device) (newDevice *model.Device, err error) {
	err = s.repository.Create(ctx, device)
	newDevice = device

	s.redis.Delete(ctx, newDevice.CacheKey().GetByBrand())
	s.redis.Delete(ctx, newDevice.CacheKey().GetByState())
	s.redis.Delete(ctx, newDevice.CacheKey().GetAll())
	s.redis.Set(ctx, newDevice.CacheKey().GetByID(), newDevice)

	return
}
