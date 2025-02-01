package service

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
	"github.com/lucaslmuller/technical-test/internal/app/device/repository"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis/cache"
)

type get struct {
	repository *repository.Device
	redis      *cache.Redis
}

func NewGet(r *repository.Device, redis *cache.Redis) get {
	return get{r, redis}
}

func (s get) ByID(ctx context.Context, ID string) (device *model.Device, err error) {
	dummy := model.Device{
		Id: &ID,
	}
	cached, err := cache.Get[*model.Device](ctx, dummy.CacheKey().GetByID(), s.redis)

	if err == nil && cached != nil {
		device = *cached
		return
	}

	device, err = s.repository.Database.Get.ByID(ctx, ID)

	if device != nil {
		s.redis.Set(ctx, device.CacheKey().GetByID(), device)
	}

	return
}

func (s get) ByBrand(ctx context.Context, brandID string) (devices []model.Device, err error) {
	dummy := model.Device{
		Brand: brandID,
	}

	cached, err := cache.Get[[]model.Device](ctx, dummy.CacheKey().GetByBrand(), s.redis)

	if err == nil && cached != nil {
		devices = *cached
		return
	}

	devices, err = s.repository.Database.Get.ByBrand(ctx, brandID)

	s.redis.Set(ctx, dummy.CacheKey().GetByBrand(), devices)

	return
}

func (s get) ByState(ctx context.Context, state string) (devices []model.Device, err error) {
	dummy := model.Device{
		State: state,
	}

	cached, err := cache.Get[[]model.Device](ctx, dummy.CacheKey().GetByState(), s.redis)

	if err == nil && cached != nil {
		devices = *cached
		return
	}

	devices, err = s.repository.Database.Get.ByState(ctx, state)

	s.redis.Set(ctx, dummy.CacheKey().GetByState(), devices)

	return
}

func (s get) All(ctx context.Context) (devices []model.Device, err error) {
	dummy := model.Device{}
	cached, err := cache.Get[[]model.Device](ctx, dummy.CacheKey().GetAll(), s.redis)

	if err == nil && cached != nil {
		devices = *cached
		return
	}

	devices, err = s.repository.Database.Get.All(ctx)

	s.redis.Set(ctx, dummy.CacheKey().GetAll(), devices)

	return
}
