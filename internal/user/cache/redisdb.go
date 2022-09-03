package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"rest-api-test/internal/user"
	"time"
)

type cache struct {
	cache *redis.Client
}

func NewCache(c *redis.Client) user.Cache {
	return &cache{cache: c}
}

func (c *cache) SetJson(ctx context.Context) {
	obj := map[string]string{
		"some-key": "some-value",
	}

	marshal, err := json.Marshal(obj)
	if err != nil {
		return
	}
	c.cache.Set(ctx, "json-obj", marshal, 10*time.Minute)
}

func (c *cache) GetJson(ctx context.Context) {
	res, err := c.cache.Get(ctx, "json-obj").Result()
	if err != nil {
		return
	}
	fmt.Println(json.Marshal(res))
}
