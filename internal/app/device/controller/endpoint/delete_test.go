package endpoint

import (
	"errors"

	"github.com/google/uuid"
	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
)

func (s *EndpointSuite) TestDeleteSuccess() {

	input := &dto.DeleteDeviceInput{
		ID: uuid.NewString(),
	}

	s.mockService.
		EXPECT().
		Delete(s.ctx, input.ID).
		Return(nil).
		Times(1)

	validationErr, err := s.endpoints.Delete(s.ctx, input)

	s.NoError(err)
	s.NoError(validationErr)
}

func (s *EndpointSuite) TestDeleteInvalidID() {
	input := &dto.DeleteDeviceInput{
		ID: "INVALID",
	}

	validationErr, err := s.endpoints.Delete(s.ctx, input)

	s.NoError(err)
	s.Error(validationErr)
	s.Equal(validationErr.Error(), "invalid id - must be a valid uuid")
}

func (s *EndpointSuite) TestDeleteWhileInUseError() {
	input := &dto.DeleteDeviceInput{
		ID: uuid.NewString(),
	}

	s.mockService.
		EXPECT().
		Delete(s.ctx, input.ID).
		Return(errors.New("cannot delete while in use")).
		Times(1)

	validationErr, err := s.endpoints.Delete(s.ctx, input)

	s.Error(err)
	s.NoError(validationErr)
	s.Equal(err.Error(), "cannot delete while in use")
}

func (s *EndpointSuite) TestDeleteInternalError() {
	input := &dto.DeleteDeviceInput{
		ID: uuid.NewString(),
	}

	s.mockService.
		EXPECT().
		Delete(s.ctx, input.ID).
		Return(errors.New("error")).
		Times(1)

	_, err := s.endpoints.Delete(s.ctx, input)

	s.Error(err)
}
