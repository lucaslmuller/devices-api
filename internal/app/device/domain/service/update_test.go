package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

func (s *ServiceSuite) TestUpdateError() {
	id := uuid.NewString()
	device := &model.Device{
		Id:    &id,
		Name:  "test",
		Brand: "test",
		State: "INVALID",
	}

	s.mockRepository.
		EXPECT().
		GetByID(s.ctx, id).
		Return(device, nil).
		Times(1)

	s.mockRepository.
		EXPECT().
		Update(s.ctx, device).
		Return(errors.New("error")).
		Times(1)

	_, err := s.service.Update(s.ctx, device)
	s.Error(err)
}

func (s *ServiceSuite) TestUpdateSuccess() {
	id := uuid.NewString()
	currentDevice := &model.Device{
		Id:    &id,
		Name:  "test",
		Brand: "test",
		State: "available",
	}

	updatedDevice := &model.Device{
		Id:    &id,
		Name:  "test",
		Brand: "updated_brand",
		State: "available",
	}

	s.mockRepository.
		EXPECT().
		GetByID(s.ctx, id).
		Return(currentDevice, nil).
		Times(1)

	s.mockRepository.
		EXPECT().
		Update(s.ctx, updatedDevice).
		Return(nil).
		Times(1)

	s.cache.EXPECT().
		Delete(s.ctx, currentDevice.CacheKey().GetByBrand()).
		Return(nil).
		Times(1)

	s.cache.EXPECT().
		Delete(s.ctx, currentDevice.CacheKey().GetByState()).
		Return(nil).
		Times(1)

	s.cache.EXPECT().
		Delete(s.ctx, updatedDevice.CacheKey().GetByBrand()).
		Return(nil).
		Times(1)

	s.cache.EXPECT().
		Delete(s.ctx, updatedDevice.CacheKey().GetAll()).
		Return(nil).
		Times(1)

	s.cache.EXPECT().
		Delete(s.ctx, updatedDevice.CacheKey().GetByState()).
		Return(nil).
		Times(1)

	s.cache.EXPECT().
		Delete(s.ctx, updatedDevice.CacheKey().GetByID()).
		Return(nil).
		Times(1)

	s.cache.EXPECT().
		Set(s.ctx, updatedDevice.CacheKey().GetByID(), updatedDevice).
		Return(nil).
		Times(1)

	result, err := s.service.Update(s.ctx, updatedDevice)
	s.NoError(err)

	s.Equal(result, updatedDevice)
}
