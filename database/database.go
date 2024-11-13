package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Init(cfg Config) (*sql.DB, error) {
	dsn := toDSN(cfg.GetHost(), cfg.GetPort(), cfg.GetUsername(), cfg.GetPassword(), cfg.GetDBName(), cfg.GetSSLMode())

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(cfg.GetMaxIdleConns())
	conn.SetMaxOpenConns(cfg.GetMaxOpenConns())

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
