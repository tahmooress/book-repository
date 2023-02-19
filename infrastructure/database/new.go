package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //nolint: revive
)

type Database struct {
	db *sql.DB
}

type Config struct {
	DBName   string
	DBHost   string
	DBPort   string
	User     string
	Password string
}

func New(ctx context.Context, cfg Config) (*Database, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.User, cfg.Password, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("connection to postgres failed: %s ,%s", err, connStr)
	}

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("connection to postgres failed: %s ,%s", err, connStr)
	}

	return &Database{
		db: db,
	}, nil
}
