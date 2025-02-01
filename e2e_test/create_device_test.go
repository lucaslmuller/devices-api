package e2etest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
	"github.com/lucaslmuller/technical-test/internal/utils"
	"github.com/stretchr/testify/suite"
)

type CreateDeviceSuite struct {
	suite.Suite
}

func TestCreateDevice(t *testing.T) {
	suite.Run(t, new(CreateDeviceSuite))
}

func (s *CreateDeviceSuite) TestCreateDeviceInvalidState() {
	device := &dto.CreateDeviceInput{
		Name:  "test",
		Brand: "test",
		State: "INVALID",
	}

	r, _, bodyStr := CreateDevice[utils.ErrorResponse](device)

	expected := utils.ErrorResponse{
		Success: false,
		Error:   "invalid state, value must be one of the following: 'active', 'inactive' or 'in-use'",
	}

	s.Equal(http.StatusBadRequest, r.StatusCode)
	s.JSONEq(StringifyData(expected), bodyStr)
}

func (s *CreateDeviceSuite) TestCreateDeviceSuccess() {
	device := &dto.CreateDeviceInput{
		Name:  "test",
		Brand: "test",
		State: "available",
	}

	r, body, bodyStr := CreateDevice[utils.SuccessResponse[*dto.Output]](device)

	s.NotNil(body)
	s.NotNil(body.Data)

	expected := utils.SuccessResponse[*dto.Output]{
		Success: true,
		Data: &dto.Output{
			Id:           body.Data.Id,
			Name:         device.Name,
			Brand:        device.Brand,
			State:        device.State,
			CreationTime: body.Data.CreationTime,
		},
	}

	s.Equal(http.StatusOK, r.StatusCode)
	s.Contains(r.Header.Get("Content-Type"), "application/json")
	s.JSONEq(StringifyData(expected), bodyStr)
}

func CreateDevice[T any](device *dto.CreateDeviceInput) (r *http.Response, responseData *T, responseStr string) {
	c := http.Client{}
	responseData = new(T)
	requestBody, _ := json.Marshal(device)

	r, _ = c.Post("http://localhost:8080/devices", "application/json", bytes.NewReader(requestBody))

	b, _ := io.ReadAll(r.Body)

	responseStr = string(b)

	json.Unmarshal(b, responseData)

	return
}

func MockDevice(state string) *dto.Output {
	device := &dto.CreateDeviceInput{
		Name:  "mock_name",
		Brand: "mock_test",
		State: state,
	}

	_, createdDevice, _ := CreateDevice[utils.SuccessResponse[*dto.Output]](device)

	return createdDevice.Data
}
