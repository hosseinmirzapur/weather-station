package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hosseinmirzapur/weatherstation/config"
	"github.com/hosseinmirzapur/weatherstation/router"
	"github.com/hosseinmirzapur/weatherstation/utils"
)

func RunServer() error {
	// fiber initialization
	app := fiber.New()

	// using cors middleware
	app.Use(cors.New())

	// only accepting json
	app.Use(func(c *fiber.Ctx) error {
		if c.Is("json") {
			return c.Next()
		}
		return c.SendString("Only JSON allowed!")
	})

	// register routes
	err := router.RegisterAppRouter(app)
	utils.FatalErr(err, "unable to register app routes")

	// returning app instance
	return app.Listen(config.GetServerConfig().Port)
}
