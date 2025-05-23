package redis

import (
	"context"
	"time"
)

func (c *Cluster) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if expiration <= 0 {
		expiration = c.defaultTTL
	}

	op := func() (interface{}, error) {
		return nil, c.Client.Set(ctx, key, value, expiration).Err()
	}

	if c.cb != nil {
		return c.retryWithBackoff(func() error {
			_, err := c.cb.Execute(op)
			return err
		})
	}

	return c.retryWithBackoff(func() error {
		_, err := op()
		return err
	})
}

func (c *Cluster) GetDel(ctx context.Context, key string) (string, error) {
	if c.cb != nil {
		res, err := c.cb.Execute(func() (interface{}, error) {
			return c.Client.GetDel(ctx, key).Result()
		})
		if err != nil {
			return "", err
		}
		return res.(string), nil
	}
	return c.Client.GetDel(ctx, key).Result()
}

func (c *Cluster) Get(ctx context.Context, key string) (string, error) {
	if c.cb != nil {
		res, err := c.cb.Execute(func() (interface{}, error) {
			return c.Client.Get(ctx, key).Result()
		})
		if err != nil {
			return "", err
		}
		return res.(string), nil
	}
	return c.Client.Get(ctx, key).Result()
}

func (c *Cluster) Del(ctx context.Context, keys ...string) (int64, error) {
	op := func() (interface{}, error) {
		return c.Client.Del(ctx, keys...).Result()
	}
	if c.cb != nil {
		res, err := c.cb.Execute(op)
		if err != nil {
			return 0, err
		}
		return res.(int64), nil
	}
	res, err := op()
	return res.(int64), err
}

func (c *Cluster) Exists(ctx context.Context, keys ...string) (int64, error) {
	op := func() (interface{}, error) {
		return c.Client.Exists(ctx, keys...).Result()
	}
	if c.cb != nil {
		res, err := c.cb.Execute(op)
		if err != nil {
			return 0, err
		}
		return res.(int64), nil
	}
	res, err := op()
	return res.(int64), err
}

func (c *Cluster) Expire(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	op := func() (interface{}, error) {
		return c.Client.Expire(ctx, key, ttl).Result()
	}
	if c.cb != nil {
		res, err := c.cb.Execute(op)
		if err != nil {
			return false, err
		}
		return res.(bool), nil
	}
	res, err := op()
	return res.(bool), err
}

func (c *Cluster) TTL(ctx context.Context, key string) (time.Duration, error) {
	op := func() (interface{}, error) {
		return c.Client.TTL(ctx, key).Result()
	}
	if c.cb != nil {
		res, err := c.cb.Execute(op)
		if err != nil {
			return 0, err
		}
		return res.(time.Duration), nil
	}
	res, err := op()
	return res.(time.Duration), err
}

func (c *Cluster) Incr(ctx context.Context, key string) (int64, error) {
	op := func() (interface{}, error) {
		return c.Client.Incr(ctx, key).Result()
	}
	if c.cb != nil {
		res, err := c.cb.Execute(op)
		if err != nil {
			return 0, err
		}
		return res.(int64), nil
	}
	res, err := op()
	return res.(int64), err
}

func (c *Cluster) HGet(ctx context.Context, key, field string) (string, error) {
	op := func() (interface{}, error) {
		return c.Client.HGet(ctx, key, field).Result()
	}
	if c.cb != nil {
		res, err := c.cb.Execute(op)
		if err != nil {
			return "", err
		}
		return res.(string), nil
	}
	res, err := op()
	return res.(string), err
}

func (c *Cluster) HSet(ctx context.Context, key string, values ...interface{}) error {
	op := func() (interface{}, error) {
		return c.Client.HSet(ctx, key, values...).Result()
	}
	if c.cb != nil {
		_, err := c.cb.Execute(op)
		return err
	}
	_, err := op()
	return err
}

func (c *Cluster) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	op := func() (interface{}, error) {
		return c.Client.HDel(ctx, key, fields...).Result()
	}
	if c.cb != nil {
		res, err := c.cb.Execute(op)
		if err != nil {
			return 0, err
		}
		return res.(int64), nil
	}
	res, err := op()
	return res.(int64), err
}
