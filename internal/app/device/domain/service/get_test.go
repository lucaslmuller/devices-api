package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

func (s *ServiceSuite) TestGetError() {
	s.cache.EXPECT().
		GetInBytes(s.ctx, (&model.Device{}).CacheKey().GetAll()).
		Return([]byte(nil), nil).
		Times(1)

	s.mockRepository.
		EXPECT().
		GetAll(s.ctx).
		Return(nil, errors.New("error")).
		Times(1)

	_, err := s.service.GetAll(s.ctx)
	s.Error(err)
}

func (s *ServiceSuite) TestGetSuccess() {
	id := uuid.NewString()
	device := &model.Device{
		Id:    &id,
		Name:  "test",
		Brand: "test",
		State: "available",
	}

	s.mockRepository.
		EXPECT().
		GetAll(s.ctx).
		Return([]model.Device{*device}, nil).
		Times(1)

	s.cache.EXPECT().
		GetInBytes(s.ctx, device.CacheKey().GetAll()).
		Return([]byte(nil), nil).Times(1)

	s.cache.EXPECT().
		Set(s.ctx, device.CacheKey().GetAll(), []model.Device{*device}).
		Return(nil).
		Times(1)

	devices, err := s.service.GetAll(s.ctx)

	s.NoError(err)
	s.GreaterOrEqual(len(devices), 1)
}
