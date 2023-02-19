package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/tahmooress/book-repository/pkg/util"
)

const defaultTTL = time.Minute * 5

type Cache struct {
	client *redis.Client
	ttl    time.Duration
}

type Config struct {
	Host     string
	Password string
	TTL      string
}

func New(ctx context.Context, cfg Config) (*Cache, error) {
	cleint := redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Password,
		DB:       0,
	})

	_, err := cleint.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("redis: New() >> %w", err)
	}

	return &Cache{
		client: cleint,
		ttl:    util.ParseDurationWithDefault(cfg.TTL, defaultTTL),
	}, nil
}
