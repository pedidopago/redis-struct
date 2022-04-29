package json

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	redistruct "github.com/pedidopago/redis-struct"
)

func New(cl redis.Cmdable) redistruct.Client {
	return &client{cl: cl}

}

type client struct {
	cl redis.Cmdable
}

func (c *client) Get(ctx context.Context, key string, target interface{}) error {
	r := c.cl.Get(ctx, key)
	if e := r.Err(); e != nil {
		if e == redis.Nil {
			return redistruct.ErrNotFound
		}
		return e
	}
	return json.Unmarshal([]byte(r.Val()), target)
}

func (c *client) Set(ctx context.Context, key string, val interface{}, ttl time.Duration) error {
	b, e := json.Marshal(val)
	if e != nil {
		return e
	}
	return c.cl.Set(ctx, key, string(b), ttl).Err()
}

func (c *client) Del(ctx context.Context, key string) error {
	return c.cl.Del(ctx, key).Err()
}
