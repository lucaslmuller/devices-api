package service

import (
	"github.com/lucaslmuller/technical-test/internal/app/device/repository"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis/cache"
)

type DeviceService struct {
	repository *repository.DeviceRepository
	redis      *cache.Redis
}

func NewService(r *repository.DeviceRepository, redis *cache.Redis) *DeviceService {
	return &DeviceService{
		repository: r,
		redis:      redis,
	}
}
