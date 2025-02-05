package endpoint

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
)

func (s *EndpointSuite) TestGetAllSuccess() {
	s.mockService.
		EXPECT().
		GetAll(s.ctx).
		Return([]model.Device{}, nil).
		Times(1)

	res, err := s.endpoints.GetAll(s.ctx)

	s.NoError(err)
	s.Equal(res, []dto.Output{})
}

func (s *EndpointSuite) TestGetAllInternalError() {
	s.mockService.
		EXPECT().
		GetAll(s.ctx).
		Return(nil, errors.New("error")).
		Times(1)

	_, err := s.endpoints.GetAll(s.ctx)

	s.Error(err)
}

func (s *EndpointSuite) TestGetByIDSuccess() {
	now := time.Now()
	input := &dto.GetDeviceByIDInput{
		ID: uuid.NewString(),
	}

	device := &model.Device{
		Id:           &input.ID,
		Name:         "test",
		Brand:        "test",
		State:        "available",
		CreationTime: &now,
	}

	s.mockService.
		EXPECT().
		GetByID(s.ctx, input.ID).
		Return(device, nil).
		Times(1)

	res, validationErr, err := s.endpoints.GetByID(s.ctx, input)

	s.NoError(err)
	s.NoError(validationErr)
	s.Equal(res, &dto.Output{
		Id:           input.ID,
		Name:         "test",
		Brand:        "test",
		State:        "available",
		CreationTime: now,
	})
}

func (s *EndpointSuite) TestGetByIDInternalError() {
	input := &dto.GetDeviceByIDInput{
		ID: uuid.NewString(),
	}

	s.mockService.
		EXPECT().
		GetByID(s.ctx, input.ID).
		Return(nil, errors.New("error")).
		Times(1)

	_, _, err := s.endpoints.GetByID(s.ctx, input)

	s.Error(err)
}

func (s *EndpointSuite) TestGetByBrandSuccess() {
	input := &dto.GetDeviceByBrandInput{
		Brand: "brand",
	}
	s.mockService.
		EXPECT().
		GetByBrand(s.ctx, input.Brand).
		Return([]model.Device{}, nil).
		Times(1)

	res, validationErr, err := s.endpoints.GetByBrand(s.ctx, input)

	s.NoError(err)
	s.NoError(validationErr)
	s.Equal(res, []dto.Output{})
}

func (s *EndpointSuite) TestGetByBrandInternalError() {
	input := &dto.GetDeviceByBrandInput{
		Brand: "brand",
	}

	s.mockService.
		EXPECT().
		GetByBrand(s.ctx, input.Brand).
		Return(nil, errors.New("error")).
		Times(1)

	_, _, err := s.endpoints.GetByBrand(s.ctx, input)

	s.Error(err)
}

func (s *EndpointSuite) TestGetByStateSuccess() {
	input := &dto.GetDeviceByStateInput{
		State: "in-use",
	}
	s.mockService.
		EXPECT().
		GetByState(s.ctx, input.State).
		Return([]model.Device{}, nil).
		Times(1)

	res, validationErr, err := s.endpoints.GetByState(s.ctx, input)

	s.NoError(err)
	s.NoError(validationErr)
	s.Equal(res, []dto.Output{})
}

func (s *EndpointSuite) TestGetByStateInvalidState() {
	input := &dto.GetDeviceByStateInput{
		State: "INVALID",
	}

	_, validationErr, err := s.endpoints.GetByState(s.ctx, input)

	s.NoError(err)
	s.Error(validationErr)
	s.Equal(validationErr.Error(), "invalid state, value must be one of the following: 'active', 'inactive' or 'in-use'")
}

func (s *EndpointSuite) TestGetByStateInternalError() {
	input := &dto.GetDeviceByStateInput{
		State: "available",
	}

	s.mockService.
		EXPECT().
		GetByState(s.ctx, input.State).
		Return(nil, errors.New("error")).
		Times(1)

	_, _, err := s.endpoints.GetByState(s.ctx, input)

	s.Error(err)
}
