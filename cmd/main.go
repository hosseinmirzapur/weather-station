package main

import (
	"github.com/hosseinmirzapur/weatherstation/server"
	"github.com/hosseinmirzapur/weatherstation/utils"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()
	utils.PanicErr(err, "unable to find or load .env file")

	// run server
	err = server.RunServer()
	utils.PanicErr(err, "unable to run the server")

}
