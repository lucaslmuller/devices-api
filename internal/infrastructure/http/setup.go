package http

import (
	"net/http"
	"time"

	"github.com/lucaslmuller/technical-test/internal/infrastructure"
)

const SERVER_CLOSE_ERROR = "http: Server closed"

type Server struct {
	*infrastructure.Resources
	HTTP *http.Server
}

func StartServer(res *infrastructure.Resources) {
	res.HTTP = &http.Server{
		Addr:              res.Config.Server.HTTP.ListenAddr,
		Handler:           res.Router,
		ReadHeaderTimeout: 30 * time.Second,
	}

	res.Logger.Info("running HTTP server")

	err := res.HTTP.ListenAndServe()

	if err != nil && err.Error() != SERVER_CLOSE_ERROR {
		res.Logger.Fatal("failed to listen and serve HTTP server")
	}
}
