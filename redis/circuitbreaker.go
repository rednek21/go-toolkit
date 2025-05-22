package redis

import (
	"github.com/sony/gobreaker"
)

func NewBreaker(cfg CircuitBreakerConfig) *gobreaker.CircuitBreaker {
	settings := gobreaker.Settings{
		Name:        "RedisCluster",
		MaxRequests: cfg.MaxRequests,
		Interval:    cfg.Interval,
		Timeout:     cfg.Timeout,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			totalReqs := counts.Requests
			totalFailures := counts.TotalFailures
			return totalReqs >= 10 && float64(totalFailures)/float64(totalReqs) >= 0.5
		},
		OnStateChange: func(name string, from, to gobreaker.State) {},
	}
	return gobreaker.NewCircuitBreaker(settings)
}
