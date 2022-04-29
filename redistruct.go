package redistruct

import (
	"context"
	"fmt"
	"time"
)

// Common errors
var (
	ErrNotFound = fmt.Errorf("not found")
)

type Client interface {
	Get(ctx context.Context, key string, target interface{}) error
	Set(ctx context.Context, key string, val interface{}, ttl time.Duration) error
	Del(ctx context.Context, key string) error
}
