package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/weatherstation/utils"
)

func GetWeatherInfo(c *fiber.Ctx) error {
	lat := c.Query("lat")
	lon := c.Query("lon")

	if lat == "" || lon == "" {
		return c.Status(400).JSON(&fiber.Map{
			"message": "declaring lon and lat as url params is necessary",
		})
	}

	lat64, err := utils.StringToF64(lat)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "bad request params",
		})
	}

	lon64, err := utils.StringToF64(lon)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "bad request params",
		})
	}

	err = utils.ValidateCoords(*lon64, *lat64)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	data, err := utils.CallWeatherApi(*lat64, *lon64)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "no data recieved from weather api",
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"data": data,
	})
}
