package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, cfg *Config) (*Database, error) {
	dsn, err := getDSNFromEnv()
	if err != nil {
		return nil, err
	}

	if cfg.MinConns <= 0 || cfg.MaxConns <= 0 || cfg.MaxIdleConns <= 0 || cfg.ConnLifetime <= 0 {
		return nil, fmt.Errorf("invalid database configuration: %v", cfg)
	}

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	poolConfig.MinConns = int32(cfg.MinConns)
	poolConfig.MaxConns = int32(cfg.MaxConns)
	poolConfig.MaxConnIdleTime = cfg.MaxIdleConns
	poolConfig.MaxConnLifetime = cfg.ConnLifetime

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	return &Database{
		Pool: pool,
	}, nil
}

func (d *Database) Ping(ctx context.Context) error {
	if err := d.Pool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}
	return nil
}

func (d *Database) Close() {
	d.Pool.Close()
}

func (d *Database) RunInTransaction(ctx context.Context, fn func(tx pgx.Tx) error) error {
	tx, err := d.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	if err := fn(tx); err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
