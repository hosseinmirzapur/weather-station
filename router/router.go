package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/weatherstation/handlers"
)

func RegisterAppRouter(app *fiber.App) {
	v1 := app.Group("/v1")
	{
		v1.Get("/weather-info", handlers.GetWeatherInfo)
	}
}

/*
 - showing temperature

 - showing humidity

 - showing atmosphere pressure

 - showing weather forecast

 - showing windspeed and direction

 - rainfall measurement

 - UV index monitoring

*/
