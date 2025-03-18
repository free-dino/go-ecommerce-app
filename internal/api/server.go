package api

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/handlers"
	"go-ecommerce-app/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Database connection failed: %v\n", err)
	}

	log.Printf("Database connection established\n")

	db.AutoMigrate(&domain.User{})

	rh := &rest.RestHandler{
		App: app,
		DB:  db,
	}

	setupRoutes(rh)

	// Add error handling for Listen()
	if err := app.Listen(config.ServerPort); err != nil {
		panic(err) // Or use a logger
	}
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
}
