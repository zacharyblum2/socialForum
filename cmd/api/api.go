package main

import (
	"log"
	"net/http"
	"time"
)

// `application` struct holds app configurations and settings.
type application struct {
	config config // Server configuration settings
}

// `config` struct holds specific configuration details,
// such as the server's address.
type config struct {
	addr string // Server listen address
	// rateLimit  // Potential rate limiting setting
	// dbConfig   // Potential database settings
}

func (app *application) mount() *http.ServeMux{
	// `mux` is an HTTP request multiplexer (router).
	mux := http.NewServeMux()

	// Handle route "/v1/health" with healthCheckHandler.
	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)

	return mux
}

// `run` starts the HTTP server and returns any startup errors.
func (app *application) run(mux *http.ServeMux) error {
	// Set up the server with the specified address and request handler (mux).
	server := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
		WriteTimeout: 30 * time.Second, // Max duration to write a response to the client
		ReadTimeout:  10 * time.Second, // Max duration to read the request from the client
		IdleTimeout:  1 * time.Minute,  // Max time to keep idle connections open
	}
	
	// Log start of server.
	// IDEAL: Have custom logger injected
	log.Printf("server has started at %s", app.config.addr)
	
	// Start the server and listen for requests; return any errors encountered.
	return server.ListenAndServe()
}

// In `main`, you would initialize `application` and call `app.run()` to start.