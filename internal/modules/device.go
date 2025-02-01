package modules

import (
	"github.com/lucaslmuller/technical-test/internal/app/device/controller"
	"github.com/lucaslmuller/technical-test/internal/app/device/controller/handler"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain"
	"github.com/lucaslmuller/technical-test/internal/app/device/repository"
	"github.com/lucaslmuller/technical-test/internal/infrastructure"
)

type deviceModule struct {
}

func setupDeviceModule(res *infrastructure.Resources) *deviceModule {
	module := &deviceModule{}

	// * Setup repositories
	repositories := repository.Setup(res)

	// * Setup services
	services := domain.SetupServices(repositories, res.Redis)

	// * Setup API (handler and usecases)
	api := handler.SetupAPI(services)

	// * Setup routes
	controller.SetupRoutes(res.Router, api)

	return module
}
