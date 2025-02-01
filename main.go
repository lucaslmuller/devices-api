package main

import (
	"context"

	"github.com/lucaslmuller/technical-test/bootstrap/cmd"
)

func main() {
	ctx := context.Background()
	cmd.Run(ctx, cmd.SetupConfig())
}
