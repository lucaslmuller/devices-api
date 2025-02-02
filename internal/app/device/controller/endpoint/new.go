package endpoint

import (
	"github.com/lucaslmuller/technical-test/internal/app/device/domain/service"
)

type DeviceEndpoints struct {
	service *service.DeviceService
}

func NewDeviceEndpoints(service *service.DeviceService) *DeviceEndpoints {
	return &DeviceEndpoints{service}
}
