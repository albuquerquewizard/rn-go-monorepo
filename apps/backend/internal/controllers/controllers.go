package controllers

import (
	"github.com/albuquerquewizard/monorepo/backend/internal/services"
)

// Controllers holds all controller instances
type Controllers struct {
	User *UserController
	// Add more controllers here as you create them
}

// NewControllers creates a new Controllers instance with all controllers
func NewControllers(services *services.Services) *Controllers {
	return &Controllers{
		User: NewUserController(services.User),
		// Add more controllers here as you create them
	}
}
