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

	if err = s.repository.Delete(ctx, ID); err != nil {
		return err
	}

	m := &model.Device{Id: &ID}
	if err = s.handleDeleteCache(ctx, m); err != nil {
		s.logger.Error(err)
	}

	return
}

func (s *DeviceService) handleDeleteCache(ctx context.Context, m *model.Device) (err error) {
	if err = s.redis.Delete(ctx, m.CacheKey().GetAll()); err != nil {
		return err
	}

	if err = s.redis.Delete(ctx, m.CacheKey().GetByID()); err != nil {
		return err
	}

	if err = s.redis.Delete(ctx, m.CacheKey().GetByBrand()); err != nil {
		return err
	}

	if err = s.redis.Delete(ctx, m.CacheKey().GetByState()); err != nil {
		return err
	}

	return nil
}
