package tools

import (
	"github.com/lucaslmuller/technical-test/internal/infrastructure"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/postgresql"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis"
)

func ConnectTools(res *infrastructure.Resources) {
	postgresql.Connect(res)
	redis.ConnectRedis(res)
}
