package middleware

import (
	"errors"
	"runtime/debug"
	"strings"

	"github.com/albuquerquewizard/monorepo/backend/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

// Custom error types for better error handling
var (
	ErrValidationFailed = errors.New("validation failed")
	ErrNotFound         = errors.New("resource not found")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrForbidden        = errors.New("forbidden")
	ErrBadRequest       = errors.New("bad request")
	ErrInternalServer   = errors.New("internal server error")
	ErrTimeout          = errors.New("request timeout")
)

// ErrorResponse represents a detailed error response
type ErrorResponse struct {
	Success   bool                   `json:"success"`
	Error     string                 `json:"error"`
	Message   string                 `json:"message,omitempty"`
	Details   map[string]interface{} `json:"details,omitempty"`
	RequestID string                 `json:"request_id,omitempty"`
	Timestamp string                 `json:"timestamp,omitempty"`
}

// GlobalErrorHandler is the main error handler for the application
func GlobalErrorHandler(c *fiber.Ctx, err error) error {
	// Get logger from context or create a new one
	logger := zerolog.Nop()
	if log := c.Locals("logger"); log != nil {
		if l, ok := log.(zerolog.Logger); ok {
			logger = l
		}
	}

	// Get request ID from context
	requestID := c.Get("X-Request-ID")
	if requestID == "" {
		requestID = c.Get("X-Correlation-ID")
	}

	// Default status code
	statusCode := fiber.StatusInternalServerError
	errorMessage := "Internal Server Error"
	userMessage := "Something went wrong. Please try again later."

	// Handle different types of errors
	switch {
	case errors.Is(err, ErrValidationFailed):
		statusCode = fiber.StatusBadRequest
		errorMessage = "Validation Failed"
		userMessage = "The provided data is invalid. Please check your input."

	case errors.Is(err, ErrNotFound):
		statusCode = fiber.StatusNotFound
		errorMessage = "Not Found"
		userMessage = "The requested resource was not found."

	case errors.Is(err, ErrUnauthorized):
		statusCode = fiber.StatusUnauthorized
		errorMessage = "Unauthorized"
		userMessage = "You are not authorized to access this resource."

	case errors.Is(err, ErrForbidden):
		statusCode = fiber.StatusForbidden
		errorMessage = "Forbidden"
		userMessage = "You do not have permission to access this resource."

	case errors.Is(err, ErrBadRequest):
		statusCode = fiber.StatusBadRequest
		errorMessage = "Bad Request"
		userMessage = "The request is malformed or contains invalid data."

	case errors.Is(err, ErrTimeout):
		statusCode = fiber.StatusRequestTimeout
		errorMessage = "Request Timeout"
		userMessage = "The request took too long to process."

	default:
		// Handle Fiber specific errors
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
			errorMessage = fiberErr.Message
			userMessage = fiberErr.Message
		}
	}

	// Log the error with context
	logger.Error().
		Err(err).
		Str("request_id", requestID).
		Str("method", c.Method()).
		Str("path", c.Path()).
		Str("ip", c.IP()).
		Str("user_agent", c.Get("User-Agent")).
		Int("status_code", statusCode).
		Str("error_type", errorMessage).
		Msg("Request error occurred")

	// In development mode, include stack trace
	var details map[string]interface{}
	// Check if we're in development mode by looking for environment variable or context
	if env := c.Locals("env"); env != nil {
		if envStr, ok := env.(string); ok && envStr == "development" {
			details = map[string]interface{}{
				"stack_trace": strings.Split(string(debug.Stack()), "\n"),
				"error":       err.Error(),
			}
		}
	}

	// Create error response
	errorResp := ErrorResponse{
		Success:   false,
		Error:     errorMessage,
		Message:   userMessage,
		Details:   details,
		RequestID: requestID,
	}

	// Add timestamp if available
	if timestamp := c.Locals("timestamp"); timestamp != nil {
		if ts, ok := timestamp.(string); ok {
			errorResp.Timestamp = ts
		}
	}

	// Return JSON response
	return c.Status(statusCode).JSON(errorResp)
}

// ValidationErrorHandler handles validation errors specifically
func ValidationErrorHandler(c *fiber.Ctx, validationErrors map[string]string) error {
	logger := zerolog.Nop()
	if log := c.Locals("logger"); log != nil {
		if l, ok := log.(zerolog.Logger); ok {
			logger = l
		}
	}

	requestID := c.Get("X-Request-ID")
	if requestID == "" {
		requestID = c.Get("X-Correlation-ID")
	}

	// Log validation errors
	logger.Warn().
		Str("request_id", requestID).
		Str("method", c.Method()).
		Str("path", c.Path()).
		Interface("validation_errors", validationErrors).
		Msg("Validation failed")

	// Use the existing utility function for consistency
	return utils.ValidationErrorResponse(c, validationErrors)
}

// NotFoundHandler handles 404 errors
func NotFoundHandler(c *fiber.Ctx) error {
	logger := zerolog.Nop()
	if log := c.Locals("logger"); log != nil {
		if l, ok := log.(zerolog.Logger); ok {
			logger = l
		}
	}

	requestID := c.Get("X-Request-ID")
	if requestID == "" {
		requestID = c.Get("X-Correlation-ID")
	}

	// Log 404 errors
	logger.Info().
		Str("request_id", requestID).
		Str("method", c.Method()).
		Str("path", c.Path()).
		Str("ip", c.IP()).
		Str("user_agent", c.Get("User-Agent")).
		Msg("Route not found")

	// Return 404 response
	errorResp := ErrorResponse{
		Success:   false,
		Error:     "Not Found",
		Message:   "The requested route was not found",
		RequestID: requestID,
	}

	// Add timestamp if available
	if timestamp := c.Locals("timestamp"); timestamp != nil {
		if ts, ok := timestamp.(string); ok {
			errorResp.Timestamp = ts
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(errorResp)
}

// MethodNotAllowedHandler handles 405 errors
func MethodNotAllowedHandler(c *fiber.Ctx) error {
	logger := zerolog.Nop()
	if log := c.Locals("logger"); log != nil {
		if l, ok := log.(zerolog.Logger); ok {
			logger = l
		}
	}

	requestID := c.Get("X-Request-ID")
	if requestID == "" {
		requestID = c.Get("X-Correlation-ID")
	}

	// Log method not allowed errors
	logger.Warn().
		Str("request_id", requestID).
		Str("method", c.Method()).
		Str("path", c.Path()).
		Str("ip", c.IP()).
		Str("user_agent", c.Get("User-Agent")).
		Msg("Method not allowed")

	// Return 405 response
	errorResp := ErrorResponse{
		Success:   false,
		Error:     "Method Not Allowed",
		Message:   "The HTTP method is not allowed for this endpoint",
		RequestID: requestID,
	}

	// Add timestamp if available
	if timestamp := c.Locals("timestamp"); timestamp != nil {
		if ts, ok := timestamp.(string); ok {
			errorResp.Timestamp = ts
		}
	}

	return c.Status(fiber.StatusMethodNotAllowed).JSON(errorResp)
}

// PanicRecoveryMiddleware recovers from panics and logs them
func PanicRecoveryMiddleware(logger zerolog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				requestID := c.Get("X-Request-ID")
				if requestID == "" {
					requestID = c.Get("X-Correlation-ID")
				}

				// Log the panic
				logger.Error().
					Interface("panic", r).
					Str("request_id", requestID).
					Str("method", c.Method()).
					Str("path", c.Path()).
					Str("stack_trace", string(debug.Stack())).
					Msg("Panic recovered")

				// Return 500 error
				errorResp := ErrorResponse{
					Success:   false,
					Error:     "Internal Server Error",
					Message:   "Something went wrong. Please try again later.",
					RequestID: requestID,
				}

				// Add timestamp if available
				if timestamp := c.Locals("timestamp"); timestamp != nil {
					if ts, ok := timestamp.(string); ok {
						errorResp.Timestamp = ts
					}
				}

				_ = c.Status(fiber.StatusInternalServerError).JSON(errorResp)
			}
		}()

		return c.Next()
	}
}
