package domain

import (
	"github.com/lucaslmuller/technical-test/internal/app/device/domain/service"
	"github.com/lucaslmuller/technical-test/internal/app/device/repository"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis/cache"
)

func SetupServices(repository *repository.Device, redis *cache.Redis) *DeviceServices {
	return &DeviceServices{
		Get:    service.NewGet(repository, redis),
		Create: service.NewCreate(repository, redis),
		Update: service.NewUpdate(repository, redis),
		Delete: service.NewDelete(repository, redis),
	}
}
