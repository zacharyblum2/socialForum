package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq" // Postgres library for handling SQL operations like arrays
)

// Posts data model
// Can move this to a `model` folder if needed for better organization
type Post struct {
	ID        int64    `json:"id"`        // ID generated in the database
	Content   string   `json:"content"`   // Post content
	Title     string   `json:"title"`     // Post title
	UserID    int64    `json:"user_id"`   // ID of the user who created the post
	CreatedAt string   `json:"created_at"` // Timestamp of post creation, set in the db
	UpdatedAt string   `json:"updated_at"` // Timestamp of post update, set in the db
	Tags      []string `json:"tags"`      // Array of tags associated with the post
}

// PostsStore handles operations related to posts in the database
type PostsStore struct {
	db *sql.DB // Database connection pool
}

// Create inserts a new post into the database and retrieves auto-generated fields.
func (s *PostsStore) Create(ctx context.Context, post *Post) error {
	// SQL query to insert a post record and return auto-generated fields
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	// Execute the query with the provided post values
	// `pq.Array(post.Tags)` converts the Go slice to a format PostgreSQL can use for arrays.
	err := s.db.QueryRowContext(
		ctx,            // Pass in the context for handling cancellation, timeouts, etc.
		query,          // The SQL query to execute
		post.Content,   // First parameter in the query, maps to $1
		post.Title,     // Second parameter in the query, maps to $2
		post.UserID,    // Third parameter in the query, maps to $3
		pq.Array(post.Tags), // Fourth parameter as an array, maps to $4
	).Scan(
		&post.ID,        // Capture generated ID from the query's RETURNING clause
		&post.CreatedAt, // Capture generated creation timestamp
		&post.UpdatedAt, // Capture generated update timestamp
	)

	// If there’s an error during query execution, return it to the caller
	if err != nil {
		return err
	}

	return nil // Return nil on successful execution
}
