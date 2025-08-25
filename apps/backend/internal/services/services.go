package services

import (
	"github.com/albuquerquewizard/monorepo/backend/internal/repositories"
)

// Services holds all service instances
type Services struct {
	User UserService
	// Add more services here as you create them
}

// NewServices creates a new Services instance with all services
func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		User: NewUserService(repos.User),
		// Add more services here as you create them
	}
}
