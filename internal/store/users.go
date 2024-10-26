package store

import (
	"context"
	"database/sql"
)

// User represents the data model for a user in the application.
// It includes fields for user ID, username, email, password, and
// the timestamp of when the user was created.
type User struct { 
	ID        int64  `json:"id"`          // Unique identifier for the user
	Username  string `json:"username"`    // Username of the user
	Email     string `json:"email"`       // Email address of the user
	Password  string `json:"password"`    // Password for user authentication
	CreatedAt string `json:"created_at"`  // Timestamp of user creation
}

// UsersStore holds the database connection for user-related operations.
type UsersStore struct {
	db *sql.DB // Database connection instance
}

// Create inserts a new user into the database.
// It executes an SQL query to insert the user's details and 
// retrieves the auto-generated ID and creation timestamp.
func (s *UsersStore) Create(ctx context.Context, user *User) error {
	// SQL query to insert a user and return auto-generated fields.
	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	// Execute the query with the user's details and scan the returned
	// values into the user struct.
	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username, // Correctly using Username for insertion
		user.Email,
		user.Password,
	).Scan(
		&user.ID,        // Scan the auto-generated ID into user.ID
		&user.CreatedAt, // Scan the created_at timestamp into user.CreatedAt
	)

	// If an error occurs during query execution, return the error.
	if err != nil {
		return err
	}

	return nil // Return nil if the user was created successfully
}
