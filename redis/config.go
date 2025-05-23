package redis

import "time"

type Config struct {
	Cluster        ClusterConfig
	Retry          RetryConfig
	CircuitBreaker CircuitBreakerConfig
	Password       string
	DialTimeout    time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxRetries     int
	PoolSize       int
	TTL            time.Duration
	DefaultTTL     time.Duration
}

type ClusterConfig struct {
	Enabled bool
	Nodes   []string
}

type RetryConfig struct {
	MaxTries   int
	Backoff    time.Duration
	MaxBackoff time.Duration
}

type CircuitBreakerConfig struct {
	Enabled     bool
	MaxRequests uint32
	Interval    time.Duration
	Timeout     time.Duration
}
