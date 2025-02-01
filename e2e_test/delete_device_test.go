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

type DeleteDeviceSuite struct {
	suite.Suite
}

func TestDeleteDevice(t *testing.T) {
	suite.Run(t, new(DeleteDeviceSuite))
}

func (s *DeleteDeviceSuite) TestDeleteDeviceInvalidUUID() {
	id := "invalid"
	r, _, bodyStr := DeleteDevice[utils.ErrorResponse](id)

	expected := utils.ErrorResponse{
		Success: false,
		Error:   "invalid id - must be a valid uuid",
	}

	s.Equal(http.StatusBadRequest, r.StatusCode)
	s.JSONEq(StringifyData(expected), bodyStr)
}

func (s *DeleteDeviceSuite) TestDeleteDeviceErrorInUse() {
	device := MockDevice("in-use")
	id := device.Id

	r, body, bodyStr := DeleteDevice[utils.ErrorResponse](id)

	s.NotNil(body)
	s.NotNil(body.Error)

	expected := utils.ErrorResponse{
		Success: false,
		Error:   "cannot delete device in use",
	}

	s.Equal(http.StatusInternalServerError, r.StatusCode)
	s.Contains(r.Header.Get("Content-Type"), "application/json")
	s.JSONEq(StringifyData(expected), bodyStr)

	// clean up
	state := "inactive"
	deviceToUpdate := &dto.UpdateDeviceInput{
		ID:    id,
		State: &state,
	}
	updateDevice[any](deviceToUpdate)
	DeleteDevice[any](id)
}

func (s *DeleteDeviceSuite) TestDeleteDeviceSuccess() {
	device := MockDevice("available")
	id := device.Id

	r, body, bodyStr := DeleteDevice[utils.SuccessResponse[*dto.Output]](id)

	s.NotNil(body)

	expected := utils.SuccessResponse[*dto.Output]{
		Success: true,
		Data:    nil,
	}

	s.Equal(http.StatusOK, r.StatusCode)
	s.Contains(r.Header.Get("Content-Type"), "application/json")
	s.JSONEq(StringifyData(expected), bodyStr)

	DeleteDevice[any](id)
}

func DeleteDevice[T any](id string) (r *http.Response, responseData *T, responseStr string) {
	c := http.Client{}

	responseData = new(T)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/devices/%s", id), bytes.NewReader([]byte("{}")))
	req.Header.Set("Content-Type", "application/json")

	r, _ = c.Do(req)

	b, _ := io.ReadAll(r.Body)

	responseStr = string(b)
	json.Unmarshal(b, responseData)

	return
}
