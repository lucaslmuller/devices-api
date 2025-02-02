package handler

import (
	"github.com/lucaslmuller/technical-test/internal/app/device/controller/endpoint"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain/service"
)

type api struct {
	Endpoints *endpoint.DeviceEndpoints
}

func SetupAPI(s *service.DeviceService) *api {
	endpoints := endpoint.NewDeviceEndpoints(s)
	return &api{
		Endpoints: endpoints,
	}
}
