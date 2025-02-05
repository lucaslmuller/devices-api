package endpoint

import (
	"errors"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
	"github.com/stretchr/testify/mock"
)

func (s *EndpointSuite) TestCreateSuccess() {
	input := &dto.CreateDeviceInput{
		Name:  "test",
		Brand: "test",
		State: "available",
	}

	d := input.ToModel()
	s.mockService.
		EXPECT().
		Create(s.ctx, mock.MatchedBy(func(device *model.Device) bool {
			return device.Name == d.Name && device.Brand == d.Brand && device.State == d.State
		})).
		Return(d, nil).
		Times(1)

	o, validationErr, err := s.endpoints.Create(s.ctx, input)

	expectedOutput := &dto.Output{}
	expectedOutput.FromModel(d)

	s.NoError(err)
	s.NoError(validationErr)
	s.Equal(o, expectedOutput)
}

func (s *EndpointSuite) TestCreateInvalidState() {
	input := &dto.CreateDeviceInput{
		Name:  "test",
		Brand: "test",
		State: "invalid",
	}

	_, validationErr, _ := s.endpoints.Create(s.ctx, input)

	s.Error(validationErr)
	s.Equal("invalid state, value must be one of the following: 'active', 'inactive' or 'in-use'", validationErr.Error())
}

func (s *EndpointSuite) TestCreateInternalError() {
	input := &dto.CreateDeviceInput{
		Name:  "test",
		Brand: "test",
		State: "available",
	}

	d := input.ToModel()
	s.mockService.
		EXPECT().
		Create(s.ctx, mock.MatchedBy(func(device *model.Device) bool {
			return device.Name == d.Name && device.Brand == d.Brand && device.State == d.State
		})).
		Return(nil, errors.New("error")).
		Times(1)

	_, _, err := s.endpoints.Create(s.ctx, input)

	s.Error(err)
}
