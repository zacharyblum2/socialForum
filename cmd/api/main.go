package main

import "log"

func main() { 
	// Initialize configuration settings.
	// Declared separately to make code cleaner.
	cfg := config{
		addr: ":8080",
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