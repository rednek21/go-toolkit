package redis

import (
	"math"
	"math/rand"
	"time"
)

func (c *Cluster) retryWithBackoff(op func() error) error {
	var err error
	for attempt := 0; attempt < c.retryCount; attempt++ {
		err = op()
		if err == nil {
			return nil
		}
		backoff := c.calcBackoff(attempt)
		time.Sleep(backoff)
	}
	return err
}

func (c *Cluster) calcBackoff(attempt int) time.Duration {
	exp := math.Pow(2, float64(attempt)) * float64(c.retryBackoff)
	jitter := float64(rand.Intn(100)) / 100 * float64(c.retryBackoff)
	total := time.Duration(exp + jitter)
	if total > c.maxBackoff {
		return c.maxBackoff
	}
	return total
}
