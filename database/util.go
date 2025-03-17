package database

import (
	"errors"
	"fmt"
	"os"
)

func getDSNFromEnv() (string, error) {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		return "", errors.New("POSTGRES_HOST environment variable not set")
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		return "", errors.New("POSTGRES_PORT environment variable not set")
	}
	username := os.Getenv("POSTGRES_USER")
	if username == "" {
		return "", errors.New("POSTGRES_USER environment variable not set")
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		return "", errors.New("POSTGRES_PASSWORD environment variable not set")
	}
	db := os.Getenv("POSTGRES_DB")
	if db == "" {
		return "", errors.New("POSTGRES_DB environment variable not set")
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, db)
	return dsn, nil
}
