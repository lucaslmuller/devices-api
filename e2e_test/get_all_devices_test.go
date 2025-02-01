package e2etest

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
	"github.com/lucaslmuller/technical-test/internal/utils"
	"github.com/stretchr/testify/suite"
)

type GetAllDevicesSuite struct {
	suite.Suite
}

func TestGetAllDevices(t *testing.T) {
	suite.Run(t, new(GetAllDevicesSuite))
}

func (s *GetAllDevicesSuite) TestGetAllDevicesEmpty() {
	r, data, _ := getAll[utils.SuccessResponse[[]dto.Output]]()

	s.Equal(http.StatusOK, r.StatusCode)
	s.Len(data.Data, 0)
	s.JSONEq(StringifyData(data.Data), "[]")
}

func (s *GetAllDevicesSuite) TestGetAllDevicesSuccess() {
	devicesList := CreateMockDevices()

	r, data, _ := getAll[utils.SuccessResponse[[]dto.Output]]()

	s.T().Log(data)
	s.Equal(http.StatusOK, r.StatusCode)
	s.Len(data.Data, len(devicesList))
	s.JSONEq(StringifyData(data.Data), StringifyData(devicesList))

	DeleteMockDevices(devicesList)
}

func CreateMockDevices() []dto.Output {
	result := []dto.Output{}
	for i := 0; i < 3; i++ {
		_, data, _ := CreateDevice[utils.SuccessResponse[*dto.Output]](&dto.CreateDeviceInput{
			Name:  "test",
			Brand: "test",
			State: "available",
		})

		result = append(result, *data.Data)
	}

	return result
}

func DeleteMockDevices(devices []dto.Output) {
	for _, d := range devices {
		DeleteDevice[any](d.Id)
	}
}

func getAll[T any]() (r *http.Response, responseData *T, responseStr string) {
	c := http.Client{}

	responseData = new(T)

	r, _ = c.Get("http://localhost:8080/devices")
	b, _ := io.ReadAll(r.Body)

	responseStr = string(b)
	json.Unmarshal(b, responseData)

	return
}
