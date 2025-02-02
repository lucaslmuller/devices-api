package cmd

import (
	"github.com/lucaslmuller/technical-test/internal/infrastructure"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/postgresql"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis"
)

func connectTools(res *infrastructure.Resources) {
	postgresql.Connect(res)
	redis.ConnectRedis(res)
}
