package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/albuquerquewizard/monorepo/backend/internal/app"
	"github.com/albuquerquewizard/monorepo/backend/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Setup logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Set log level
	level, err := zerolog.ParseLevel(cfg.Log.Level)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	logger.Info().Msgf("🚀 Starting %s in %s mode", cfg.App.Name, cfg.App.Env)
	logger.Info().Msgf("📡 Server will be available at http://localhost:%s", cfg.App.Port)

	// Create and initialize application
	application := app.NewApp(cfg, logger)

	// Start server in a goroutine
	go func() {
		if err := application.Start(); err != nil {
			logger.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info().Msg("🛑 Shutting down server...")

	// Gracefully shutdown the server
	if err := application.Shutdown(); err != nil {
		logger.Error().Err(err).Msg("Error during server shutdown")
	}

	logger.Info().Msg("✅ Server stopped")
}
