package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func PanicErr(err error, str string) {
	if err != nil {
		panic(str)
	}
}

func FatalErr(err error, str string) {
	if err != nil {
		log.Fatal(str)
	}
}

func ServerErr(str string, statusCode int) error {
	return fiber.NewError(statusCode, str)
}
