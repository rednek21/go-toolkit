package redis

import "time"

type Config struct {
	Cluster struct {
		Enabled bool     `yaml:"enabled"`
		Nodes   []string `yaml:"nodes"`
	} `yaml:"cluster"`

	Password       string               `yaml:"password"`
	DialTimeout    time.Duration        `yaml:"dial-timeout"`
	ReadTimeout    time.Duration        `yaml:"read-timeout"`
	WriteTimeout   time.Duration        `yaml:"write-timeout"`
	MaxRetries     int                  `yaml:"max-retries"`
	PoolSize       int                  `yaml:"pool-size"`
	Retry          RetryConfig          `yaml:"retry"`
	CircuitBreaker CircuitBreakerConfig `yaml:"circuit-breaker"`
}

type RetryConfig struct {
	MaxTries   int           `yaml:"max-tries"`
	Backoff    time.Duration `yaml:"backoff"`
	MaxBackoff time.Duration `yaml:"max-backoff"`
}

type CircuitBreakerConfig struct {
	Enabled     bool          `yaml:"enabled"`
	MaxRequests uint32        `yaml:"max-requests"`
	Interval    time.Duration `yaml:"interval"`
	Timeout     time.Duration `yaml:"timeout"`
}
