package infrastructure

import (
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis/cache"
)

var EmptyCfg = App{}

// Configuration contains all settings for this project
type Configuration struct {
	Server   Server
	Postgres *Postgres
	Redis    *Redis
}

func NewConfiguration() *Configuration {
	return &Configuration{
		Server: Server{
			HTTP: HTTP{
				ListenAddr: ":8080",
				Network:    "tcp",
			},
		},
		Postgres: &Postgres{},
		Redis:    &Redis{},
	}
}

type HTTP struct {
	Network    string
	ListenAddr string
}

type Server struct {
	HTTP
}

type App struct {
	BaseURL string
	Token   *string
}

type Redis struct {
	*cache.Configuration
}

type Settings struct {
	MaxOpenConns         int
	MaxIdleConns         int
	MaxConnLifeTimeInMin int
}

type Postgres struct {
	Host       *string
	TimeoutSec *int
	Port       int
	DBName     string
	Username   string
	Password   string
	Settings   Settings
}
