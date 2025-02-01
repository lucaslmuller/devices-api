package service

import (
	"context"
	"errors"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
	"github.com/lucaslmuller/technical-test/internal/app/device/repository"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis/cache"
)

type update struct {
	repository *repository.Device
	redis      *cache.Redis
}

func NewUpdate(r *repository.Device, redis *cache.Redis) update {
	return update{r, redis}
}

func (s update) Device(ctx context.Context, device *model.Device) (updatedDevice *model.Device, err error) {
	currentDevice, err := s.repository.Database.Get.ByID(ctx, *device.Id)

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

	err = s.repository.Database.Update.Device(ctx, device)
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
