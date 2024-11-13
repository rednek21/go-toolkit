package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var db *sql.DB

func Init(cfg *Config) (*sql.DB, error) {
	dsn := toDSN(cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(cfg.MaxIdleConns)
	conn.SetMaxOpenConns(cfg.MaxOpenConns)

	db = conn

	return db, nil
}

func GetDB() (*sql.DB, error) {
	if db == nil {
		return nil, errors.New("database not initialized")
	}
	return db, nil
}

func CloseDB() error {
	if db == nil {
		return nil
	}
	return db.Close()
}

func toDSN(host string, port int, username, password, dbName, sslMode string) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		username, password, host, port, dbName, sslMode)
}
