package main

import (
	"context"

	"github.com/lucaslmuller/technical-test/cmd"
)

func main() {
	ctx := context.Background()
	cmd.Run(ctx, cmd.SetupConfig())
}
