package service

import (
	"context"
	"errors"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
	"github.com/lucaslmuller/technical-test/internal/app/device/repository"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis/cache"
)

type delete struct {
	repository *repository.Device
	redis      *cache.Redis
}

func NewDelete(r *repository.Device, redis *cache.Redis) delete {
	return delete{r, redis}
}

func (s delete) Device(ctx context.Context, ID string) (err error) {
	currentDevice, err := s.repository.Database.Get.ByID(ctx, ID)

	if err != nil {
		return err
	}

	if currentDevice == nil {
		return errors.New("device not found")
	}

	if currentDevice.State == "in-use" {
		return errors.New("cannot delete device in use")
	}

	err = s.repository.Database.Delete.ByID(ctx, ID)

	m := model.Device{Id: &ID}
	s.redis.Delete(ctx, m.CacheKey().GetAll())
	s.redis.Delete(ctx, m.CacheKey().GetByID())
	s.redis.Delete(ctx, m.CacheKey().GetByBrand())
	s.redis.Delete(ctx, m.CacheKey().GetByState())

	return
}
