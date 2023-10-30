package config

import (
	"fmt"
	"os"
)

type serverConfig struct {
	Port string
}

func GetServerConfig() *serverConfig {
	return &serverConfig{
		Port: fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
	}
}
