package service

import (
	"github.com/lucaslmuller/technical-test/internal/app/device/repository"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis/cache"
	"go.uber.org/zap"
)

type DeviceService struct {
	repository repository.IRepository
	redis      cache.Cache
	logger     *zap.SugaredLogger
}

func NewService(r repository.IRepository, c cache.Cache, l *zap.SugaredLogger) *DeviceService {
	return &DeviceService{
		repository: r,
		redis:      c,
		logger:     l,
	}
}
