package app

import (
	"strings"
	"time"

	"github.com/albuquerquewizard/monorepo/backend/internal/config"
	"github.com/albuquerquewizard/monorepo/backend/internal/controllers"
	"github.com/albuquerquewizard/monorepo/backend/internal/middleware"
	"github.com/albuquerquewizard/monorepo/backend/internal/models"
	"github.com/albuquerquewizard/monorepo/backend/internal/repositories"
	"github.com/albuquerquewizard/monorepo/backend/internal/routes"
	"github.com/albuquerquewizard/monorepo/backend/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/rs/zerolog"
)

// App holds the application dependencies
type App struct {
	FiberApp    *fiber.App
	Config      *config.Config
	Database    *config.Database
	Repos       *repositories.Repositories
	Services    *services.Services
	Controllers *controllers.Controllers
}

// NewApp creates a new application instance
func NewApp(cfg *config.Config, logger zerolog.Logger) *App {
	// Initialize database
	database := config.NewDatabase(cfg)

	// Run database migrations
	if err := models.AutoMigrate(database.DB); err != nil {
		logger.Fatal().Err(err).Msg("Failed to run database migrations")
	}

	// Seed initial data
	if err := models.SeedData(database.DB); err != nil {
		logger.Warn().Err(err).Msg("Failed to seed initial data")
	}

	// Initialize repositories
	repos := repositories.NewRepositories(database.DB)

	// Initialize services
	svcs := services.NewServices(repos)

	// Initialize controllers
	ctrls := controllers.NewControllers(svcs)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.GlobalErrorHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	// Setup middleware
	setupMiddleware(app, cfg, logger)

	// Setup custom error handlers
	app.Use(func(c *fiber.Ctx) error {
		// Add timestamp to context for error handlers
		c.Locals("timestamp", time.Now().UTC().Format(time.RFC3339))
		return c.Next()
	})

	// Setup routes
	routes.SetupRoutes(app, ctrls)

	// Setup 404 handler
	app.Use(middleware.NotFoundHandler)

	// Setup 405 handler
	app.Use(middleware.MethodNotAllowedHandler)

	// Setup health check
	app.Get("/health", func(c *fiber.Ctx) error {
		// Check database health
		if err := database.HealthCheck(); err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"status": "unhealthy",
				"error":  "Database connection failed",
			})
		}

		return c.JSON(fiber.Map{
			"status":    "healthy",
			"timestamp": time.Now().UTC(),
			"version":   "1.0.0",
		})
	})

	return &App{
		FiberApp:    app,
		Config:      cfg,
		Database:    database,
		Repos:       repos,
		Services:    svcs,
		Controllers: ctrls,
	}
}

// setupMiddleware configures all middleware for the application
func setupMiddleware(app *fiber.App, cfg *config.Config, appLogger zerolog.Logger) {
	// Custom panic recovery middleware with logging
	app.Use(middleware.PanicRecoveryMiddleware(appLogger))

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(cfg.CORS.AllowedOrigins, ","),
		AllowMethods:     strings.Join(cfg.CORS.AllowedMethods, ","),
		AllowHeaders:     strings.Join(cfg.CORS.AllowedHeaders, ","),
		AllowCredentials: true,
	}))

	// Logger middleware
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))

	// Timeout middleware for all routes
	app.Use(timeout.New(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusRequestTimeout).JSON(fiber.Map{
			"error": "Request timeout",
		})
	}, 30*time.Second))
}

// Start starts the application
func (a *App) Start() error {
	return a.FiberApp.Listen(":" + a.Config.App.Port)
}

// Shutdown gracefully shuts down the application
func (a *App) Shutdown() error {
	if err := a.Database.Close(); err != nil {
		return err
	}
	return a.FiberApp.Shutdown()
}
