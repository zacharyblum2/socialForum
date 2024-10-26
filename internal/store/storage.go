package store

import (
	"context"
	"database/sql"
)

// Storage struct holds the repositories for different entities, in this case,
// posts and users.
type Storage struct {
	Posts interface { // Interface for the Posts repository
		Create(context.Context, *Post) error // Method to create a new post
	}
	Users interface { // Interface for the Users repository
		Create(context.Context, *User) error // Method to create a new user
	}
}

// NewStorage initializes a new Storage instance with the provided database 
// connection. It sets up the Posts and Users repositories.
func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostsStore{db}, // Initialize the Posts repository with the given 
		// database connection
		Users: &UsersStore{db}, // Initialize the Users repository with the given 
		// database connection
	}
}
