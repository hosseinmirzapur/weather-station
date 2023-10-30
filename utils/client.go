package utils

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hosseinmirzapur/weatherstation/config"
	"github.com/hosseinmirzapur/weatherstation/types"
)

func CallWeatherApi(lat, lon float64) (*types.ApiResponse, error) {
	endpoint := fmt.Sprintf(
		"%s?lat=%f&lon=%f&appid=%s",
		config.GetWeatherConfig().Endpoint,
		lat,
		lon,
		config.GetWeatherConfig().Api_Key,
	)

	agent := fiber.Get(endpoint)

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return nil, ServerErr("unable to call weather api", statusCode)
	}

	var data types.ApiResponse

	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, ServerErr("unable to unmarshal weather api data", 400)
	}

	return &data, nil
}

func ValidateCoords(lon, lat float64) error {
	if lon < -180 || lon > 180 {
		return errors.New("invalid longitude")
	}

	if lat < -90 || lat > 90 {
		return errors.New("invalid latitude")
	}

	return nil

}
