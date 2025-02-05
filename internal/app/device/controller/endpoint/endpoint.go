package endpoint

import "github.com/lucaslmuller/technical-test/internal/app/device/domain/service"

type DeviceEndpoints struct {
	service service.IService
}

func NewDeviceEndpoints(service service.IService) *DeviceEndpoints {
	return &DeviceEndpoints{service}
}
