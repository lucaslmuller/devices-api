package e2etest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
	"github.com/lucaslmuller/technical-test/internal/utils"
	"github.com/stretchr/testify/suite"
)

type UpdateDeviceSuite struct {
	suite.Suite
}

func TestUpdateDevice(t *testing.T) {
	suite.Run(t, new(UpdateDeviceSuite))
}

func (s *UpdateDeviceSuite) TestUpdateDeviceInvalidUUID() {
	updatedDevice := &dto.UpdateDeviceInput{
		ID: "invalid",
	}

	r, _, bodyStr := updateDevice[utils.ErrorResponse](updatedDevice)

	expected := utils.ErrorResponse{
		Success: false,
		Error:   "invalid id - must be a valid uuid",
	}

	s.Equal(http.StatusBadRequest, r.StatusCode)
	s.JSONEq(StringifyData(expected), bodyStr)
}

func (s *UpdateDeviceSuite) TestUpdateDeviceNameErrorInUse() {
	device := MockDevice("in-use")

	id := device.Id

	newName := "updated_name"

	updatedDevice := &dto.UpdateDeviceInput{
		ID:   id,
		Name: &newName,
	}

	r, body, bodyStr := updateDevice[utils.SuccessResponse[*dto.Output]](updatedDevice)

	s.NotNil(body)
	s.NotNil(body.Data)

	expected := utils.ErrorResponse{
		Success: false,
		Error:   "cannot change name while in use",
	}

	s.Equal(http.StatusBadRequest, r.StatusCode)
	s.Contains(r.Header.Get("Content-Type"), "application/json")
	s.JSONEq(StringifyData(expected), bodyStr)

	// clean up
	state := "inactive"
	updateDevice[any](&dto.UpdateDeviceInput{
		ID:    id,
		State: &state,
	})
	DeleteDevice[any](id)
}

func (s *UpdateDeviceSuite) TestUpdateDeviceBrandErrorInUse() {
	device := MockDevice("in-use")

	id := device.Id

	newBrand := "updated_brand"

	updatedDevice := &dto.UpdateDeviceInput{
		ID:   id,
		Name: &newBrand,
	}

	r, body, bodyStr := updateDevice[utils.SuccessResponse[*dto.Output]](updatedDevice)

	s.NotNil(body)
	s.NotNil(body.Data)

	expected := utils.ErrorResponse{
		Success: false,
		Error:   "cannot change brand while in use",
	}

	s.Equal(http.StatusBadRequest, r.StatusCode)
	s.Contains(r.Header.Get("Content-Type"), "application/json")
	s.JSONEq(StringifyData(expected), bodyStr)

	// clean up
	state := "inactive"
	updateDevice[any](&dto.UpdateDeviceInput{
		ID:    id,
		State: &state,
	})
	DeleteDevice[any](id)
}

func (s *UpdateDeviceSuite) TestUpdateDeviceNameSuccess() {
	device := MockDevice("inactive")

	id := device.Id

	newName := "updated_name"

	updatedDevice := &dto.UpdateDeviceInput{
		ID:   id,
		Name: &newName,
	}

	r, body, bodyStr := updateDevice[utils.SuccessResponse[*dto.Output]](updatedDevice)

	s.NotNil(body)
	s.NotNil(body.Data)

	expected := utils.SuccessResponse[*dto.Output]{
		Success: true,
		Data:    device,
	}

	expected.Data.Name = newName

	s.Equal(http.StatusOK, r.StatusCode)
	s.Contains(r.Header.Get("Content-Type"), "application/json")
	s.JSONEq(StringifyData(expected), bodyStr)

	// clean up
	DeleteDevice[any](id)
}

func (s *UpdateDeviceSuite) TestUpdateDeviceBrandSuccess() {
	device := MockDevice("inactive")

	id := device.Id

	newBrand := "updated_brand"

	updatedDevice := &dto.UpdateDeviceInput{
		ID:   id,
		Name: &newBrand,
	}

	r, body, bodyStr := updateDevice[utils.SuccessResponse[*dto.Output]](updatedDevice)

	s.NotNil(body)
	s.NotNil(body.Data)

	expected := utils.SuccessResponse[*dto.Output]{
		Success: true,
		Data:    device,
	}

	expected.Data.Brand = newBrand

	s.Equal(http.StatusOK, r.StatusCode)
	s.Contains(r.Header.Get("Content-Type"), "application/json")
	s.JSONEq(StringifyData(expected), bodyStr)

	// clean up
	DeleteDevice[any](id)
}

func updateDevice[T any](device *dto.UpdateDeviceInput) (r *http.Response, responseData *T, responseStr string) {
	c := http.Client{}
	requestBody, _ := json.Marshal(device)
	responseData = new(T)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8080/devices/%s", device.ID), bytes.NewReader(requestBody))

	r, _ = c.Do(req)
	b, _ := io.ReadAll(r.Body)

	responseStr = string(b)

	json.Unmarshal(b, responseData)

	return
}
