package endpoint

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain/model"
	"github.com/stretchr/testify/mock"
)

func (s *EndpointSuite) TestUpdateSuccess() {
	name := "test"

	input := &dto.UpdateDeviceInput{
		ID:   uuid.NewString(),
		Name: &name,
	}

	now := time.Now()
	currentDevice := &model.Device{
		Id:           &input.ID,
		Name:         "test",
		Brand:        "test",
		State:        "available",
		CreationTime: &now,
	}

	deviceModel := currentDevice.Merge(input.ToModel())

	s.mockService.
		EXPECT().
		Update(s.ctx, mock.MatchedBy(func(x *model.Device) bool {
			return x.Id == &input.ID && x.Name == *input.Name
		})).
		Return(deviceModel, nil).
		Times(1)

	o, validationErr, err := s.endpoints.Update(s.ctx, input)

	expectedOutput := &dto.Output{}
	expectedOutput.FromModel(deviceModel)

	s.NoError(err)
	s.NoError(validationErr)
	s.Equal(o, expectedOutput)
}

func (s *EndpointSuite) TestUpdateInvalidID() {
	name := "test"

	input := &dto.UpdateDeviceInput{
		ID:   "INVALID",
		Name: &name,
	}

	_, validationErr, _ := s.endpoints.Update(s.ctx, input)

	s.Error(validationErr)
	s.Equal(validationErr.Error(), "invalid id - must be a valid uuid")
}

func (s *EndpointSuite) TestUpdateInvalidState() {
	state := "INVALID"

	input := &dto.UpdateDeviceInput{
		ID:    uuid.NewString(),
		State: &state,
	}

	_, validationErr, _ := s.endpoints.Update(s.ctx, input)

	s.Error(validationErr)
	s.Equal(validationErr.Error(), "invalid state, value must be one of the following: 'active', 'inactive' or 'in-use'")
}

func (s *EndpointSuite) TestUpdateBrandWhileInUseError() {
	brand := "anything"

	input := &dto.UpdateDeviceInput{
		ID:    uuid.NewString(),
		Brand: &brand,
	}

	s.mockService.
		EXPECT().
		Update(s.ctx, mock.MatchedBy(func(d *model.Device) bool {
			return true
		})).
		Return(nil, errors.New("cannot change brand while in use")).
		Times(1)

	_, validationErr, err := s.endpoints.Update(s.ctx, input)

	s.Error(err)
	s.NoError(validationErr)
	s.Equal(err.Error(), "cannot change brand while in use")
}

func (s *EndpointSuite) TestUpdateNameWhileInUseError() {
	name := "anything"

	input := &dto.UpdateDeviceInput{
		ID:    uuid.NewString(),
		Brand: &name,
	}

	s.mockService.
		EXPECT().
		Update(s.ctx, mock.MatchedBy(func(d *model.Device) bool {
			return true
		})).
		Return(nil, errors.New("cannot change name while in use")).
		Times(1)

	_, validationErr, err := s.endpoints.Update(s.ctx, input)

	s.Error(err)
	s.NoError(validationErr)
	s.Equal(err.Error(), "cannot change name while in use")
}

func (s *EndpointSuite) TestUpdateInternalError() {
	name := "test"

	input := &dto.UpdateDeviceInput{
		ID:   uuid.NewString(),
		Name: &name,
	}

	s.mockService.
		EXPECT().
		Update(s.ctx, mock.MatchedBy(func(x *model.Device) bool {
			return true
		})).
		Return(nil, errors.New("error")).
		Times(1)

	_, _, err := s.endpoints.Update(s.ctx, input)

	s.Error(err)
}
