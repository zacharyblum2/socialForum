package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// healthCheckResponse represents the JSON structure for the health check response.
type healthCheckResponse struct {
	Status    int    `json:"status"`
	Timestamp string `json:"timestamp"`
}

// healthCheckHandler is an HTTP handler function for health checks.
// It responds with a JSON message containing the status and current timestamp.
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a response struct with status and timestamp.
	response := healthCheckResponse{
		Status:    http.StatusOK,
		Timestamp: time.Now().Format(time.RFC3339), // Format the current time in UTC.
		// Using time.RFC3339 ensures the timestamp is in a standard,
		// easily parseable format.
		// Why this over a local time?
		// time.Now().Local().Format would use the local time zone,
		// which may not be suitable for APIs.
	}

	// Encode the response struct as JSON and write it to the response writer.
	// The `err` variable captures any error that may occur during encoding.
	w.Header().Set("Content-Type", "application/json")

	// The `err` variable is declared then checked if nil.
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// If encoding fails, respond with an internal server error.
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Log the health check request with a timestamp.
	log.Printf("Health check accessed at %s", response.Timestamp)
}
