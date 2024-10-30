package db

import (
	"context"
	"database/sql"
	"time"
)

// Repeating parameters in order to separate internal from API.
func New(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	
	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)
	
	// If takes more than 5 seconds to connect, we timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	// Verifies connection is still alive, starts one if necessary.
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}