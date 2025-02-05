package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

func (s *ServiceSuite) TestDeleteError() {
	id := uuid.NewString()
	device := &model.Device{
		Id:    &id,
		Name:  "test",
		Brand: "test",
		State: "INVALID",
	}

	s.mockRepository.
		EXPECT().
		Create(s.ctx, device).
		Return(errors.New("error")).
		Times(1)

	_, err := s.service.Create(s.ctx, device)
	s.Error(err)
}

func (s *ServiceSuite) TestDeleteSuccess() {
	id := uuid.NewString()
	device := &model.Device{
		Id:    &id,
		Name:  "test",
		Brand: "test",
		State: "available",
	}

	s.mockRepository.
		EXPECT().
		Create(s.ctx, device).
		Return(nil).
		Times(1)

	s.cache.EXPECT().
		Delete(s.ctx, device.CacheKey().GetByBrand()).
		Return(nil).
		Times(1)

	s.cache.EXPECT().
		Delete(s.ctx, device.CacheKey().GetByState()).
		Return(nil).
		Times(1)

	s.cache.EXPECT().
		Delete(s.ctx, device.CacheKey().GetAll()).
		Return(nil).
		Times(1)

	s.cache.EXPECT().
		Set(s.ctx, device.CacheKey().GetByID(), device).
		Return(nil).
		Times(1)

	createdDevice, err := s.service.Create(s.ctx, device)
	s.NoError(err)

	s.Equal(createdDevice, device)
}
