package controller

import (
	"github.com/go-chi/chi/v5"
)

type controller struct {
	router *chi.Mux
}

func SetupRoutes(router *chi.Mux, api IDeviceAPI) {
	c := controller{
		chi.NewRouter(),
	}

	c.registerRoutes(api)
	router.Mount("/devices", c.router)
}

func (c controller) registerRoutes(api IDeviceAPI) {
	c.router.Get("/state/{state}", api.GetByState)
	c.router.Get("/brand/{brand}", api.GetByBrand)
	c.router.Get("/{id}", api.GetByID)
	c.router.Get("/", api.GetAll)
	c.router.Post("/", api.Create)
	c.router.Put("/{id}", api.Update)
	c.router.Delete("/{id}", api.Delete)
}
