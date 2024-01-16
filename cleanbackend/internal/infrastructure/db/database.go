package db

import (
	"time"

	"github.com/ProlificLabs/captrivia/internal/infrastructure/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// NewDatabase creates a new database connection based on the given configuration.
func NewDatabase(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	// Good practice to do a ping to ensure connection is alive
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Setting reasonable defaults for the database connection pool.
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(15 * time.Minute)

	return db, nil
}