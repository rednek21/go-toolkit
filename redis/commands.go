package redis

import (
	"context"
	"fmt"
	"time"
)

// Set sets a key with a given value and expiration
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

// SAdd adds one or more members to a set
func (c *Cluster) SAdd(ctx context.Context, key string, members ...interface{}) (int64, error) {
	op := func() (interface{}, error) {
		return c.Client.SAdd(ctx, key, members...).Result()
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

// SMembers retrieves all members of a set
func (c *Cluster) SMembers(ctx context.Context, key string) ([]string, error) {
	op := func() (interface{}, error) {
		return c.Client.SMembers(ctx, key).Result()
	}
	if c.cb != nil {
		res, err := c.cb.Execute(op)
		if err != nil {
			return nil, err
		}
		return res.([]string), nil
	}
	res, err := op()
	return res.([]string), err
}

// SIsMember checks if a member exists in a set
func (c *Cluster) SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	op := func() (interface{}, error) {
		return c.Client.SIsMember(ctx, key, member).Result()
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

// SMIsMember checks if multiple members exist in a set. Returns a slice of booleans
func (c *Cluster) SMIsMember(ctx context.Context, key string, members ...interface{}) ([]bool, error) {
	op := func() (interface{}, error) {
		return c.Client.SMIsMember(ctx, key, members...).Result()
	}
	if c.cb != nil {
		res, err := c.cb.Execute(op)
		if err != nil {
			return nil, err
		}
		return res.([]bool), nil
	}
	res, err := op()
	return res.([]bool), err
}

// GetDel gets the value of a key and deletes it
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

// Get retrieves the value of a key
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

// Del deletes one or more keys
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

// Exists checks if one or more keys exist
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

// Expire sets a time-to-live (TTL) for a key
func (c *Cluster) Expire(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	if ttl <= 0 {
		ttl = c.defaultTTL
	}

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

// TTL gets the remaining time-to-live for a key
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

// Touch updates the TTL of a key if it exists
func (c *Cluster) Touch(ctx context.Context, key string, ttl time.Duration) error {
	if ttl <= 0 {
		ttl = c.defaultTTL
	}
	op := func() (interface{}, error) {
		ok, err := c.Client.Expire(ctx, key, ttl).Result()
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, fmt.Errorf("key %s does not exist or ttl not set", key)
		}
		return nil, nil
	}
	if c.cb != nil {
		_, err := c.cb.Execute(op)
		return err
	}
	_, err := op()
	return err
}

// Persist removes the TTL from a key, making it persistent
func (c *Cluster) Persist(ctx context.Context, key string) (bool, error) {
	op := func() (interface{}, error) {
		return c.Client.Persist(ctx, key).Result()
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

// Incr increments the integer value of a key by one
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

// HGet gets the value of a field in a hash
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

// HSet sets one or more fields in a hash
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

// HDel deletes one or more fields from a hash
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
