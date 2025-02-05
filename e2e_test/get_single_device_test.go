package e2etest

import (
	"fmt"
	"io"
	"net/http"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
	"github.com/lucaslmuller/technical-test/internal/utils"
)

func (s *e2eSuite) TestGetSingleDeviceInvalidUUID() {
	r, body := getDevice("1234")

	expected := utils.ErrorResponse{
		Success: false,
		Error:   "invalid id - must be a valid uuid",
	}

	s.Equal(http.StatusBadRequest, r.StatusCode)
	s.JSONEq(StringifyData(expected), body)
}

func (s *e2eSuite) TestGetSingleDeviceNotFound() {
	r, body := getDevice("0d3ec017-290e-4f6e-b2f6-a97209aa4abb")

	expected := utils.ErrorResponse{
		Success: false,
		Error:   "Device not found",
	}

	s.Equal(http.StatusNotFound, r.StatusCode)
	s.Contains(r.Header.Get("Content-Type"), "application/json")
	s.JSONEq(StringifyData(expected), body)
}

func (s *e2eSuite) TestGetSingleDeviceSuccess() {
	_, device, _ := CreateDevice[utils.SuccessResponse[*dto.Output]](&dto.CreateDeviceInput{
		Name:  "test",
		Brand: "test",
		State: "available",
	})

	r, body := getDevice(device.Data.Id)

	expected := utils.SuccessResponse[*dto.Output]{
		Success: true,
		Data:    device.Data,
	}

	s.Equal(http.StatusOK, r.StatusCode)
	s.Contains(r.Header.Get("Content-Type"), "application/json")
	s.JSONEq(StringifyData(expected), body)

	DeleteDevice[any](device.Data.Id)
}

func getDevice(id string) (r *http.Response, response string) {
	c := http.Client{}
	r, _ = c.Get(fmt.Sprintf("http://localhost:8080/devices/%s", id))
	b, _ := io.ReadAll(r.Body)
	response = string(b)
	return
}
