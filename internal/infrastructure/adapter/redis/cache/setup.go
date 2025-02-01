package cache

import (
	"github.com/redis/go-redis/v9"
)

/*
Options keeps the settings to setup redis connection.
https://pkg.go.dev/github.com/go-redis/redis/v9@v9.0.0-beta.2#Options
*/
type Options redis.Options

type Configuration struct {
	Options  Options
	Timezone string
}

func Connect(config Configuration) *Redis {
	var v = redis.Options(config.Options)
	r := &Redis{
		Client: redis.NewClient(&v),
	}

	return r
}
