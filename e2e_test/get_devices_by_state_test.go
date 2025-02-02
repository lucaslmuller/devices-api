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

type GetDevicesByStateSuite struct {
	suite.Suite
}

func TestGetDevicesByState(t *testing.T) {
	suite.Run(t, new(GetDevicesByStateSuite))
}

func (s *GetDevicesByStateSuite) TestGetDevicesByStateSuccess() {
	devicesList := CreateMockDevices()

	r, data, _ := getByState[utils.SuccessResponse[[]dto.Output]]()

	s.Equal(http.StatusOK, r.StatusCode)
	s.GreaterOrEqual(len(data.Data), 0)

	DeleteMockDevices(devicesList)
}

func getByState[T any]() (r *http.Response, responseData *T, responseStr string) {
	c := http.Client{}

	responseData = new(T)

	r, _ = c.Get(fmt.Sprintf("http://localhost:8080/devices/state/%s", "available"))
	b, _ := io.ReadAll(r.Body)

	responseStr = string(b)
	json.Unmarshal(b, responseData)

	return
}
