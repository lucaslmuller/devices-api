package shutdown

import (
	"context"
	"time"

	"github.com/lucaslmuller/technical-test/internal/infrastructure"
)

func closeHTTPServer(res *infrastructure.Resources) {
	res.Logger.Info("closing http server")

	ctxTimeout, cancel := context.WithTimeout(res.Ctx, 30*time.Second)
	defer cancel()

	res.HTTP.SetKeepAlivesEnabled(false)
	err := res.HTTP.Shutdown(ctxTimeout)
	if err != nil {
		res.Logger.Fatal("could not shutdown the HTTP server gracefully")
	}

	res.Logger.Info("HTTP server has been stopped, closing connections")
}
