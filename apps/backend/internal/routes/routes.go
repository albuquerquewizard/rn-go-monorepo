package routes

import (
	"github.com/albuquerquewizard/monorepo/backend/internal/controllers"
	"github.com/albuquerquewizard/monorepo/backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(app *fiber.App, controllers *controllers.Controllers) {
	api := app.Group("/api")

	// Health check endpoint
	api.Get("/health", func(c *fiber.Ctx) error {
		return utils.SuccessResponse(c, "Welcome to Fiber API", fiber.Map{
			"status":  "success",
			"version": "1.0.0",
		})
	})

	// User routes (public for practice projects)
	users := api.Group("/users")
	users.Post("/", controllers.User.CreateUser)
	users.Get("/", controllers.User.ListUsers)
	users.Get("/:id", controllers.User.GetUser)
	users.Put("/:id", controllers.User.UpdateUser)
	users.Delete("/:id", controllers.User.DeleteUser)

	// TODO: Add more API routes here
}
