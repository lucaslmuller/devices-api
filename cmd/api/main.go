package main

import (
	"context"

	"github.com/lucaslmuller/technical-test/config"
	"github.com/lucaslmuller/technical-test/internal/infrastructure"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/http"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/shutdown"
	"github.com/lucaslmuller/technical-test/internal/modules"
	"github.com/lucaslmuller/technical-test/tools"
)

func main() {
	ctx := context.Background()
	config := config.SetupConfig()
	resources := infrastructure.NewResources(ctx, config)

	tools.ConnectTools(resources)
	http.NewRouter(resources)
	modules.SetupModules(resources)

	// Start HTTP server
	go http.StartServer(resources)

	// Listen shutdown in background to handle it gracefully
	shutdown.Gracefully(resources)
}
