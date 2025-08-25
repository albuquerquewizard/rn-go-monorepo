package models

import (
	"log"

	"github.com/albuquerquewizard/monorepo/backend/internal/config"
	"gorm.io/gorm"
)

// Models returns all models for auto migration
func Models() []interface{} {
	return []interface{}{
		&User{},
		// Add more models here as you create them for different practice projects:
		// &Product{},
		// &Order{},
		// &Category{},
		// &Post{},
		// &Comment{},
	}
}

// AutoMigrate runs database migrations for all models
func AutoMigrate(db *gorm.DB) error {
	log.Println("🔄 Running database migrations...")

	if err := db.AutoMigrate(Models()...); err != nil {
		return err
	}

	log.Println("✅ Database migrations completed successfully")
	return nil
}

// SeedData seeds the database with initial data
func SeedData(db *gorm.DB) error {
	log.Println("🌱 Seeding database with initial data...")

	// Load configuration to get app information
	cfg := config.LoadConfig()
	log.Printf("📋 App: %s (%s)", cfg.App.Name, cfg.App.Env)

	// No initial data needed for practice projects
	// You can add sample data here for different project types:
	// - E-commerce: sample products, categories
	// - Blog: sample posts, users
	// - Task manager: sample tasks, projects

	log.Println("✅ Database seeding completed")
	return nil
}
