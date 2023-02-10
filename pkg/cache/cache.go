package cache

import (
	"context"
	"time"
)

type Key string
type Value string

type Cache interface {
	Get(context.Context, Key) (string, error)
	Put(context.Context, Key, Value, time.Duration) error
	Delete(context.Context, Key) error
	Exist(context.Context, Key) (bool, error)
	Expire(context.Context, Key, time.Duration) error
}
