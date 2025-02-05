package cache

import "context"

type Cache interface {
	GetInBytes(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, data any) (err error)
	Delete(ctx context.Context, keys ...string) error
}
