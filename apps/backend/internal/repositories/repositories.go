package repositories

import (
	"gorm.io/gorm"
)

// Repositories holds all repository instances
type Repositories struct {
	User UserRepository
	// Add more repositories here as you create them
}

// NewRepositories creates a new Repositories instance with all repositories
func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: NewUserRepository(db),
		// Add more repositories here as you create them
	}
}
