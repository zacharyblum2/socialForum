package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

// Return http.Handler instead of chi.Mux
// chi.Mux implements http.Handler so we can return http.Handler instead.
// http.Handler is package agnostic & standardized.
// chi.Mux has more fields & methods, but the context only requires
// a http.Handler so this works.
func (app *application) mount() http.Handler {
	// Utilize package chi for router (implementing Mux).
	// Chi comes nested routing (see below) and middleware support.
	r := chi.NewRouter()

	// Declaring middleware.
	r.Use(middleware.RequestID)  // Adds a unique request ID for tracing.
	r.Use(middleware.RealIP)     // Retrieves the real IP of the client.
	r.Use(middleware.Recoverer)   // Recovers from panics to avoid crashes.
	r.Use(middleware.Logger)      // Logs details of each request.

	// Set a timeout value on the request context (ctx) that will signal
	// through ctx.Done() that the request has timed out.
	r.Use(middleware.Timeout(60 * time.Second))
	
	// Create 'group' based on common v1 prefix.
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	// posts

	// users

	// auth

	return r
}

// `run` starts the HTTP server and returns any startup errors.
func (app *application) run(mux http.Handler) error {
	// Set up the server with the specified address and request handler (mux).
	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: 30 * time.Second, // Max duration to write response.
		ReadTimeout:  10 * time.Second, // Max duration to read request.
		IdleTimeout:  1 * time.Minute,   // Max time to keep idle connections.
	}
	
	// Log the start of the server.
	// IDEAL: Have custom logger injected
	log.Printf("server has started at %s", app.config.addr)
	
	// Start the server and listen for requests; return any errors.
	return server.ListenAndServe()
}

// In `main`, you would initialize `application` and call `app.run()` to start.
