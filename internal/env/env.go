package env

import (
	"os"
	"strconv"
	"strings"
)

// GetString retrieves the value of the specified environment variable as a string.
// If the variable is not set, it returns the provided fallback value.
func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key) // Check if the environment variable exists.

	if !ok {
		return fallback 
	}

	return val 
}

// GetInt retrieves the value of the specified environment variable as an int.
// If the variable is not set or cannot be converted, it returns the provided 
// fallback value.
func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key) 

	if !ok {
		return fallback 
	}

	// Attempt to convert the string value to an int.
	valAsInt, err := strconv.Atoi(val) 
	if err != nil {
		return fallback 
	}

	return valAsInt 
}

// GetBool retrieves the value of the specified environment variable as a bool.
// It recognizes common representations of true/false, returning the fallback 
// value if unrecognized.
func GetBool(key string, fallback bool) bool {
	val, ok := os.LookupEnv(key) 

	if !ok {
		return fallback
	}

	// Check for common representations of true/false in a case-insensitive manner.
	switch strings.ToLower(val) {
	case "1", "true": 
		return true
	case "0", "false":
		return false
	}
	
	return fallback
}
