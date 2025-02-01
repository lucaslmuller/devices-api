package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucaslmuller/technical-test/internal/infrastructure"
)

func NewRouter(res *infrastructure.Resources) {
	res.Logger.Info("creating router")
	res.Router = chi.NewRouter()

	res.Router.Get("/health-check", healthCheck)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
