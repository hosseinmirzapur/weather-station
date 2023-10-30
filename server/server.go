package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hosseinmirzapur/weatherstation/config"
	"github.com/hosseinmirzapur/weatherstation/router"
)

func RunServer() error {
	// fiber initialization
	app := fiber.New()

	// using cors middleware
	app.Use(cors.New())

	// register routes
	router.RegisterAppRouter(app)

	// returning app instance
	return app.Listen(config.GetServerConfig().Port)
}
