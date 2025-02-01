package handler

import (
	"github.com/lucaslmuller/technical-test/internal/app/device/controller"
	"github.com/lucaslmuller/technical-test/internal/app/device/controller/endpoint"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain"
)

type api struct {
	Endpoints *controller.DeviceEndpoints
}

func SetupAPI(services *domain.DeviceServices) *api {
	return &api{
		Endpoints: &controller.DeviceEndpoints{
			Get:    endpoint.NewGetEndpoints(services),
			Create: endpoint.NewCreationEndpoints(services),
			Update: endpoint.NewUpdateEndpoints(services),
			Delete: endpoint.NewDeletionEndpoints(services),
		},
	}
}
