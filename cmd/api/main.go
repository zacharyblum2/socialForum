package main

import (
	"log"

	"github.com/zacharyblum2/socialForum/internal/env"
)

func main() { 
	// Initialize configuration settings.
	// Declared separately to make code cleaner.
	cfg := config{
		addr: env.GetString("ADDR", ":8080"), // Get address from env
	}

	// Initialize application with configuration.
	app := &application{
		config: cfg,
	}
	
	// Initialize mux.
	// `mux` is an HTTP request multiplexer (router).
	mux := app.mount()

	// Attempt to start the server.
	// If app.run(mux) returns an error, log.Fatal will log and exit.
	log.Fatal(app.run(mux))
}