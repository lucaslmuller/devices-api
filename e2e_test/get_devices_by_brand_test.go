package e2etest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
	"github.com/lucaslmuller/technical-test/internal/utils"
	"github.com/stretchr/testify/suite"
)

type GetDevicesByBrandSuite struct {
	suite.Suite
}

func TestGetDevicesByBrand(t *testing.T) {
	suite.Run(t, new(GetDevicesByBrandSuite))
}

func (s *GetDevicesByBrandSuite) TestGetDevicesByBrandEmpty() {
	r, data, _ := getByBrand[utils.SuccessResponse[[]dto.Output]]()

	s.Equal(http.StatusOK, r.StatusCode)
	s.Len(data.Data, 0)
	s.JSONEq(StringifyData(data.Data), "[]")
}

func (s *GetDevicesByBrandSuite) TestGetDevicesByBrandSuccess() {
	devicesList := CreateMockDevices()

	r, data, _ := getByBrand[utils.SuccessResponse[[]dto.Output]]()

	s.Equal(http.StatusOK, r.StatusCode)
	s.Len(data.Data, len(devicesList))
	s.JSONEq(StringifyData(data.Data), StringifyData(devicesList))

	DeleteMockDevices(devicesList)
}

func getByBrand[T any]() (r *http.Response, responseData *T, responseStr string) {
	c := http.Client{}

	responseData = new(T)

	r, _ = c.Get(fmt.Sprintf("http://localhost:8080/devices/brand/%s", "available"))
	b, _ := io.ReadAll(r.Body)

	responseStr = string(b)
	json.Unmarshal(b, responseData)

	return
}
