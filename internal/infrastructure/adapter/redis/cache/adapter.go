package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

type CacheKeyComponent struct {
	Key   string
	Value string
}

type CacheKey struct {
	entity     *string
	components []CacheKeyComponent
}

func NewKey(entity string) *CacheKey {
	return &CacheKey{
		entity:     &entity,
		components: []CacheKeyComponent{},
	}
}

func (c *CacheKey) AddComponent(key string, value string) *CacheKey {
	c.components = append(c.components, CacheKeyComponent{
		Key:   key,
		Value: value,
	})
	return c
}

func (c *CacheKey) ToString() string {
	keyArr := make([]string, 0)
	keyArr = append(keyArr, *c.entity)

	for _, p := range c.components {
		keyArr = append(keyArr, fmt.Sprintf("%s:%s", p.Key, p.Value))
	}

	return strings.Join(keyArr, "#")
}

func Get[model any](ctx context.Context, key string, c Cache) (*model, error) {
	datab, err := c.GetInBytes(ctx, key)
	if err != nil || datab == nil {
		return nil, err
	}

	var m model
	err = json.Unmarshal(datab, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// This function returns the data from cache in bytes.
func (c *Redis) GetInBytes(ctx context.Context, key string) ([]byte, error) {
	data := c.Client.Get(ctx, key)
	if err := data.Err(); err != nil {
		if err.Error() == "redis: nil" {
			return []byte(nil), nil
		}

		return []byte(nil), err
	}

	return []byte(data.Val()), nil
}

func (c *Redis) Set(ctx context.Context, key string, data any) (err error) {
	datab, err := dataToByte(data)
	if err != nil {
		return
	}

	result := c.Client.Set(ctx, key, datab, time.Duration(0))
	if err := result.Err(); err != nil {
		return err
	}

	return nil
}

func (c *Redis) Delete(ctx context.Context, keys ...string) error {
	result := c.Client.Del(ctx, keys...)
	if err := result.Err(); err != nil {
		return err
	}
	return nil
}

func dataToByte(data any) (datab []byte, err error) {
	if str, isString := data.(string); isString {
		return []byte(str), nil
	}

	return json.Marshal(data)
}
