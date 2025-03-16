package api

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/handlers"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	rh := &rest.RestHandler{
		App: app,
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
