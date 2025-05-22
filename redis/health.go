package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync/atomic"
	"time"
)

func (c *Cluster) PingAll(ctx context.Context) bool {
	var alive int32
	_ = c.Client.ForEachShard(ctx, func(ctx context.Context, shard *redis.Client) error {
		if shard.Ping(ctx).Err() == nil {
			atomic.AddInt32(&alive, 1)
		}
		return nil
	})
	c.IsAvailable.Store(alive > 0)
	return alive > 0
}

func (c *Cluster) IsAlive() bool {
	return c.IsAvailable.Load()
}

func (c *Cluster) startHealthCheckLoop(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			c.PingAll(ctx)
		}
	}
}
