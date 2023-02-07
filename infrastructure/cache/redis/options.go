package redis

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/go-redis/redis/v8"
)

type Options struct {
	opts redis.Options
	ex   time.Duration
}

// Option is a function adapter to change config of the CloudWatch struct
type Option func(*Options)

func WithCredential(username string, password string) Option {
	return func(o *Options) {
		o.opts.Username, o.opts.Password = username, password
	}
}

func WithDB(db int) Option {
	return func(o *Options) {
		o.opts.DB = db
	}
}

func WithNetwork(network string) Option {
	return func(o *Options) {
		o.opts.Network = network
	}
}

func WithDialer(dialer func(ctx context.Context, network, addr string) (net.Conn, error)) Option {
	return func(o *Options) {
		o.opts.Dialer = dialer
	}
}

func WithOnConnect(onConnect func(ctx context.Context, cn *redis.Conn) error) Option {
	return func(o *Options) {
		o.opts.OnConnect = onConnect
	}
}

func WithMaxRetries(m int) Option {
	return func(o *Options) {
		o.opts.MaxRetries = m
	}
}

func WithMinRetryBackoff(d time.Duration) Option {
	return func(o *Options) {
		o.opts.MinRetryBackoff = d
	}
}

func WithMaxRetryBackoff(d time.Duration) Option {
	return func(o *Options) {
		o.opts.MaxRetryBackoff = d
	}
}

func WithDialTimeout(d time.Duration) Option {
	return func(o *Options) {
		o.opts.DialTimeout = d
	}
}

func WithReadTimeout(d time.Duration) Option {
	return func(o *Options) {
		o.opts.ReadTimeout = d
	}
}

func WithWriteTimeout(d time.Duration) Option {
	return func(o *Options) {
		o.opts.WriteTimeout = d
	}
}

func WithPoolFIFO(v bool) Option {
	return func(o *Options) {
		o.opts.PoolFIFO = v
	}
}

func WithPoolSize(size int) Option {
	return func(o *Options) {
		o.opts.PoolSize = size
	}
}

func WithMinIdleConns(c int) Option {
	return func(o *Options) {
		o.opts.MinIdleConns = c
	}
}

func WithMaxConnAge(d time.Duration) Option {
	return func(o *Options) {
		o.opts.MaxConnAge = d
	}
}

func WithPoolTimeout(d time.Duration) Option {
	return func(o *Options) {
		o.opts.PoolTimeout = d
	}
}

func WithIdleTimeout(d time.Duration) Option {
	return func(o *Options) {
		o.opts.IdleTimeout = d
	}
}

func WithIdleCheckFrequency(d time.Duration) Option {
	return func(o *Options) {
		o.opts.IdleCheckFrequency = d
	}
}

func WithTLSConfig(cfg *tls.Config) Option {
	return func(o *Options) {
		o.opts.TLSConfig = cfg
	}
}

func WithDefaultEx(d time.Duration) Option {
	return func(o *Options) {
		o.ex = d
	}
}
