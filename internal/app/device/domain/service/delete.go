package service

import (
	"context"
	"errors"

	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

func (s *DeviceService) Delete(ctx context.Context, ID string) (err error) {
	currentDevice, err := s.repository.GetByID(ctx, ID)

	if err != nil {
		return err
	}

	if currentDevice == nil {
		return errors.New("device not found")
	}

	if currentDevice.State == "in-use" {
		return errors.New("cannot delete device in use")
	}

	err = s.repository.Delete(ctx, ID)

	m := model.Device{Id: &ID}
	s.redis.Delete(ctx, m.CacheKey().GetAll())
	s.redis.Delete(ctx, m.CacheKey().GetByID())
	s.redis.Delete(ctx, m.CacheKey().GetByBrand())
	s.redis.Delete(ctx, m.CacheKey().GetByState())

	return
}
