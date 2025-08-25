package controllers

import (
	"strconv"

	"github.com/albuquerquewizard/monorepo/backend/internal/models"
	"github.com/albuquerquewizard/monorepo/backend/internal/services"
	"github.com/albuquerquewizard/monorepo/backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// UserController handles HTTP requests for user operations
type UserController struct {
	userService services.UserService
}

// NewUserController creates a new user controller
func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// CreateUser handles POST /api/users
func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	// Validate required fields
	if user.Username == "" || user.Password == "" {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Username and password are required")
	}

	// Create user
	if err := c.userService.CreateUser(ctx.Context(), &user); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	// Don't return password in response
	user.Password = ""

	return utils.SuccessResponse(ctx, "User created successfully", user)
}

// GetUser handles GET /api/users/:id
func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid user ID")
	}

	user, err := c.userService.GetUserByID(ctx.Context(), uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			return utils.ErrorResponse(ctx, fiber.StatusNotFound, "User not found")
		}
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(ctx, "User retrieved successfully", user)
}

// UpdateUser handles PUT /api/users/:id
func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid user ID")
	}

	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	user.ID = uint(id)

	if err := c.userService.UpdateUser(ctx.Context(), &user); err != nil {
		if err.Error() == "user not found" {
			return utils.ErrorResponse(ctx, fiber.StatusNotFound, "User not found")
		}
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	// Don't return password in response
	user.Password = ""

	return utils.SuccessResponse(ctx, "User updated successfully", user)
}

// DeleteUser handles DELETE /api/users/:id
func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid user ID")
	}

	if err := c.userService.DeleteUser(ctx.Context(), uint(id)); err != nil {
		if err.Error() == "user not found" {
			return utils.ErrorResponse(ctx, fiber.StatusNotFound, "User not found")
		}
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(ctx, "User deleted successfully", nil)
}

// ListUsers handles GET /api/users
func (c *UserController) ListUsers(ctx *fiber.Ctx) error {
	offset, _ := strconv.Atoi(ctx.Query("offset", "0"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	// Validate pagination parameters
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	users, total, err := c.userService.ListUsers(ctx.Context(), offset, limit)
	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	response := fiber.Map{
		"users":  users,
		"total":  total,
		"offset": offset,
		"limit":  limit,
	}

	return utils.SuccessResponse(ctx, "Users retrieved successfully", response)
}
