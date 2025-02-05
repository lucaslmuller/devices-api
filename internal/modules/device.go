package modules

import (
	"github.com/lucaslmuller/technical-test/internal/app/device/controller"
	"github.com/lucaslmuller/technical-test/internal/app/device/controller/handler"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain/service"
	"github.com/lucaslmuller/technical-test/internal/app/device/repository"
	"github.com/lucaslmuller/technical-test/internal/infrastructure"
)

type deviceModule struct {
}

func setupDeviceModule(res *infrastructure.Resources) *deviceModule {
	module := &deviceModule{}

	// * Setup repositories
	repository := repository.NewRepository(res)

	// * Setup services
	s := service.NewService(repository, res.Redis, res.Logger)

	// * Setup API (handler and usecases)
	api := handler.SetupAPI(s)

	// * Setup routes
	controller.SetupRoutes(res.Router, api)

	return module
}
