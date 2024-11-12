package database

import (
	"database/sql"
	"github.com/pressly/goose/v3"
	"log"
)

func MigrateUp(db *sql.DB, migrationsPath string) error {
	if migrationsPath == "" {
		migrationsPath = "./migrations"
	}

	log.Println("Applying migrations from:", migrationsPath)
	if err := goose.Up(db, migrationsPath); err != nil {
		return err
	}
	return nil
}

func MigrateUpTo(db *sql.DB, migrationsPath string, version int64) error {
	if migrationsPath == "" {
		migrationsPath = "./migrations"
	}

	log.Printf("Applying migrations up to version: %d", version)
	if err := goose.UpTo(db, migrationsPath, version); err != nil {
		return err
	}
	return nil
}

func MigrateDown(db *sql.DB, migrationsPath string) error {
	if migrationsPath == "" {
		migrationsPath = "./migrations"
	}

	log.Println("Rolling back migrations from:", migrationsPath)
	if err := goose.Down(db, migrationsPath); err != nil {
		return err
	}
	return nil
}

func MigrateDownTo(db *sql.DB, migrationsPath string, version int64) error {
	if migrationsPath == "" {
		migrationsPath = "./migrations"
	}

	log.Printf("Rolling back migrations down to version: %d", version)
	if err := goose.DownTo(db, migrationsPath, version); err != nil {
		return err
	}
	return nil
}
