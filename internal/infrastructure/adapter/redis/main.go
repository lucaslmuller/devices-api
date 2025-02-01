package redis

import (
	"context"
	"time"

	"github.com/lucaslmuller/technical-test/internal/infrastructure"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis/cache"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis(res *infrastructure.Resources) {
	cfg := res.Config.Redis
	if cfg != nil && cfg.Configuration != nil {
		res.Redis = cache.Connect(*cfg.Configuration)

		err := TestConnection(res.Ctx, res.Redis.Client)
		if err != nil {
			res.Logger.Error("failed to connect to Redis")
			panic(err)
		}

		res.Logger.Info("Redis connected")
		return
	}

	res.Logger.Error("missing cache Redis configuration")
	panic("missing cache Redis configuration")
}

func CloseConnection(res *infrastructure.Resources) {
	res.Logger.Info("closing connection with Redis")

	err := res.Redis.Client.Close()
	if err != nil {
		res.Logger.Error("error when closing redis connection")
		return

	}
}

func TestConnection(ctx context.Context, client *redis.Client) (err error) {
	for i := 1; i <= 3; i++ {
		err = client.Ping(ctx).Err()
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		return nil
	}

	return
}
