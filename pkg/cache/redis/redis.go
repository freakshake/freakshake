package redis

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mehdieidi/freakshake/pkg/cache"
)

type redisCache struct {
	client     *redis.Client
	expiration time.Duration
}

func New(host, port string, opts ...Option) cache.Cache {
	addr := net.JoinHostPort(host, port)

	var options Options

	redisOptions, err := redis.ParseURL(addr)
	if err != nil {
		options.opts.Addr = addr
	} else {
		options.opts = *redisOptions
	}

	for _, opt := range opts {
		opt(&options)
	}

	return &redisCache{
		client:     redis.NewClient(&options.opts),
		expiration: options.ex,
	}
}

func (r *redisCache) Get(ctx context.Context, k cache.Key) (string, error) {
	v, err := r.client.Get(ctx, string(k)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", cache.ErrKeyNotFound
		}
		return "", err
	}
	return v, nil
}

func (r *redisCache) Put(ctx context.Context, k cache.Key, v cache.Value, d time.Duration) error {
	if d == 0 {
		d = r.expiration
	}

	return r.client.Set(ctx, string(k), v, d).Err()
}

func (r *redisCache) Delete(ctx context.Context, k cache.Key) error {
	v, err := r.client.Del(ctx, string(k)).Result()
	if err != nil {
		return err
	}

	if v == 0 {
		return cache.ErrKeyNotFound
	}

	return nil
}

func (r *redisCache) Exist(ctx context.Context, k cache.Key) (bool, error) {
	v, err := r.client.Exists(ctx, string(k)).Result()
	return v != 0, err
}

func (r *redisCache) Expire(ctx context.Context, k cache.Key, d time.Duration) error {
	v, err := r.client.Expire(ctx, string(k), d).Result()
	if err != nil {
		return err
	}

	if !v {
		return cache.ErrKeyNotFound
	}

	return nil
}
