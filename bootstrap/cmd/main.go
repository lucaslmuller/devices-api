package cmd

import (
	"context"

	"github.com/lucaslmuller/technical-test/internal/infrastructure"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/http"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/shutdown"
	"github.com/lucaslmuller/technical-test/internal/modules"
)

func Run(ctx context.Context, config *infrastructure.Configuration) {
	resources := infrastructure.NewResources(ctx, config)

	connectTools(resources)
	http.NewRouter(resources)
	modules.SetupModules(resources)

	// Start HTTP server
	go http.StartServer(resources)

	// Listen shutdown in background to handle it gracefully
	shutdown.Gracefully(resources)
}
