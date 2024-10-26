package main

import (
	"log"

	"github.com/zacharyblum2/socialForum/internal/env"
	"github.com/zacharyblum2/socialForum/internal/store"
)

func main() { 
	// Initialize database configuration settings.
	// Declaring settings here makes the code cleaner and more readable.
	dbConfig := dbConfig{
		// Database connection string with default
		addr: env.GetString(
			"DB_ADDR", 
			"postgres://user:adminpassword@localhost/social?sslmode=disabled"), 
		// Limit on maximum open database connections
		maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30), 
		// Limit on maximum idle database connections
		maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30), 
		// Maximum idle time for database connections
		maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15min"), 
	}

	// Initialize application configuration with server address and database settings.
	cfg := config{
		// Retrieve the server address from environment variables or use default ":8080"
		addr: env.GetString("ADDR", ":8080"), 
		// Assign the database configuration created above
		db: dbConfig,
	}

	// Initialize the storage layer (e.g., for database interactions).
	// The argument `nil` will be replaced with a valid `*sql.DB` instance in a real application.
	store := store.NewStorage(nil)

	// Create a new `application` instance with the configuration and storage layer.
	// `app` is the main structure that holds configuration and storage dependencies.
	app := &application{
		config: cfg,
		store:  store,
	}
	
	// Initialize HTTP request multiplexer (router).
	// `mux` will handle routing of HTTP requests to specific handlers in the application.
	mux := app.mount()

	// Attempt to start the HTTP server with the `mux` router.
	// If app.run(mux) returns an error, log.Fatal will log the error and exit the application.
	log.Fatal(app.run(mux))
}
