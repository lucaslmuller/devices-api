package service

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
	"github.com/lucaslmuller/technical-test/internal/app/device/repository"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis/cache"
)

type create struct {
	repository *repository.Device
	redis      *cache.Redis
}

func NewCreate(r *repository.Device, redis *cache.Redis) create {
	return create{r, redis}
}

func (s create) Device(ctx context.Context, device *model.Device) (newDevice *model.Device, err error) {
	err = s.repository.Database.Create.NewDevice(ctx, device)
	newDevice = device

	s.redis.Delete(ctx, newDevice.CacheKey().GetByBrand())
	s.redis.Delete(ctx, newDevice.CacheKey().GetByState())
	s.redis.Delete(ctx, newDevice.CacheKey().GetAll())
	s.redis.Set(ctx, newDevice.CacheKey().GetByID(), newDevice)

	return
}
