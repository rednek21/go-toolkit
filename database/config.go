package database

import "time"

type Config struct {
	MinConns     int
	MaxConns     int
	MaxIdleConns time.Duration
	ConnLifetime time.Duration
}
