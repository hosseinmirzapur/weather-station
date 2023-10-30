package config

import "os"

type weatherConfig struct {
	Api_Key string
}

func GetWeatherConfig() *weatherConfig {
	return &weatherConfig{
		Api_Key: os.Getenv("OPENWEATHER_API_KEY"),
	}
}
