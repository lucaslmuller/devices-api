package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lucaslmuller/technical-test/internal/infrastructure"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis/cache"
)

func SetupConfig() (config *infrastructure.Configuration) {
	config = infrastructure.NewConfiguration()
	// host := "postgres"
	host := os.Getenv("POSTGRES_HOST")
	config.Postgres.Host = &host
	config.Postgres.DBName = os.Getenv("POSTGRES_DB")
	config.Postgres.Username = os.Getenv("POSTGRES_USER")
	config.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")

	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	config.Postgres.Port = port

	config.Redis = &infrastructure.Redis{}
	options := &cache.Options{
		// Addr: "redis:6379",
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		DB:   0,
	}

	cacheConfig := &cache.Configuration{
		Options: *options,
	}

	config.Redis = &infrastructure.Redis{
		Configuration: cacheConfig,
	}

	return
}
