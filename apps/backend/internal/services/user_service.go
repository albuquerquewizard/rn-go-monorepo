package services

import (
	"context"
	"errors"

	"github.com/albuquerquewizard/monorepo/backend/internal/models"
	"github.com/albuquerquewizard/monorepo/backend/internal/repositories"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// UserService defines the interface for user business logic
type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uint) error
	ListUsers(ctx context.Context, offset, limit int) ([]models.User, int64, error)
	AuthenticateUser(ctx context.Context, username, password string) (*models.User, error)
}

// userService implements UserService
type userService struct {
	userRepo repositories.UserRepository
	validate *validator.Validate
}

// NewUserService creates a new user service
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
		validate: validator.New(),
	}
}

// CreateUser creates a new user with validation and password hashing
func (s *userService) CreateUser(ctx context.Context, user *models.User) error {
	// Validate user data
	if err := s.validate.Struct(user); err != nil {
		return err
	}

	// Check if username already exists
	existingUser, err := s.userRepo.GetByUsername(ctx, user.Username)
	if err == nil && existingUser != nil {
		return errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Create user
	return s.userRepo.Create(ctx, user)
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

// GetUserByUsername retrieves a user by username
func (s *userService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return s.userRepo.GetByUsername(ctx, username)
}

// UpdateUser updates an existing user
func (s *userService) UpdateUser(ctx context.Context, user *models.User) error {
	// Validate user data
	if err := s.validate.Struct(user); err != nil {
		return err
	}

	// Check if user exists
	existingUser, err := s.userRepo.GetByID(ctx, user.ID)
	if err != nil {
		return err
	}

	// If password is being updated, hash it
	if user.Password != "" && user.Password != existingUser.Password {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	} else {
		// Keep existing password if not updated
		user.Password = existingUser.Password
	}

	return s.userRepo.Update(ctx, user)
}

// DeleteUser deletes a user by ID
func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	return s.userRepo.Delete(ctx, id)
}

// ListUsers retrieves a paginated list of users
func (s *userService) ListUsers(ctx context.Context, offset, limit int) ([]models.User, int64, error) {
	return s.userRepo.List(ctx, offset, limit)
}

// AuthenticateUser authenticates a user with username and password
func (s *userService) AuthenticateUser(ctx context.Context, username, password string) (*models.User, error) {
	// Get user by username
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("user account is deactivated")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
