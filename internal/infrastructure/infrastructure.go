package infrastructure

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucaslmuller/technical-test/internal/infrastructure/adapter/redis/cache"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

type Resources struct {
	Ctx    context.Context
	Config *Configuration
	Router *chi.Mux
	HTTP   *http.Server
	Tools
}

type Tools struct {
	Logger     *zap.SugaredLogger
	Redis      *cache.Redis
	PostgreSQL *bun.DB
}

func NewResources(ctx context.Context, config *Configuration) *Resources {
	res := &Resources{
		Ctx:    ctx,
		Config: config,
		Tools: Tools{
			Redis:      &cache.Redis{},
			PostgreSQL: &bun.DB{},
		},
	}

	res.newLogger()

	return res
}

func (res *Resources) newLogger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	res.Logger = sugar
}
