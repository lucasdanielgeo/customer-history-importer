package db

import (
	"database/sql"
	"fmt"

	"github.com/lucasdanielgeo/customer-history-importer/internal/infra/env"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", env.GetPostgresConnectionString())
	if err != nil {
		// %w?
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
