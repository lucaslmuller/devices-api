package shutdown

import (
	"github.com/lucaslmuller/technical-test/internal/infrastructure"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/postgresql"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis"
)

func closeConnections(res *infrastructure.Resources) {
	postgresql.Disconnect(res)
	redis.CloseConnection(res)

	res.Logger.Info("connections closed")
}
