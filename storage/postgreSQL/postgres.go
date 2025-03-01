package storage

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"movieWishlistAPI/cfg"
)

type PostgreSQL struct {
	DB *sql.DB
}

func NewPostgresDB(config *cfg.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", config.PostgresURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}
	return db, nil
}
