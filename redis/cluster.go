package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sony/gobreaker"
	"sync/atomic"
	"time"
)

type Cluster struct {
	Client       *redis.ClusterClient
	IsAvailable  atomic.Bool
	cb           *gobreaker.CircuitBreaker
	retryCount   int
	retryBackoff time.Duration
	maxBackoff   time.Duration
	defaultTTL   time.Duration
}

func InitRedisCluster(ctx context.Context, cfg Config) (*Cluster, error) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        cfg.Cluster.Nodes,
		Password:     cfg.Password,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		PoolSize:     cfg.PoolSize,
		MaxRetries:   cfg.MaxRetries,
	})

	cluster := &Cluster{
		Client:       rdb,
		retryCount:   cfg.Retry.MaxTries,
		retryBackoff: cfg.Retry.Backoff,
		maxBackoff:   cfg.Retry.MaxBackoff,
		defaultTTL:   cfg.DefaultTTL,
	}

	if cfg.CircuitBreaker.Enabled {
		cb := NewBreaker(cfg.CircuitBreaker)
		cluster.cb = cb
	}

	go cluster.startHealthCheckLoop(ctx, 5*time.Second)

	return cluster, nil
}
