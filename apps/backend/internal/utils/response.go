package utils

import (
	"github.com/gofiber/fiber/v2"
)

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Meta    interface{} `json:"meta,omitempty"` // For pagination, filters, etc.
}

// SuccessResponse returns a success response
func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SuccessResponseWithMeta returns a success response with metadata
func SuccessResponseWithMeta(c *fiber.Ctx, message string, data interface{}, meta interface{}) error {
	return c.JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

// ErrorResponse returns an error response
func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(Response{
		Success: false,
		Error:   message,
	})
}

// ValidationErrorResponse returns a validation error response
func ValidationErrorResponse(c *fiber.Ctx, errors map[string]string) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Success: false,
		Error:   "Validation failed",
		Data:    errors,
	})
}

// PaginatedResponse returns a paginated response with metadata
func PaginatedResponse(c *fiber.Ctx, message string, data interface{}, total, page, limit int) error {
	meta := fiber.Map{
		"pagination": fiber.Map{
			"total": total,
			"page":  page,
			"limit": limit,
			"pages": (total + limit - 1) / limit, // Calculate total pages
		},
	}

	return SuccessResponseWithMeta(c, message, data, meta)
}
