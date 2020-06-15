package main

import (
	"github.com/rapando/finance-api/api/controllers"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func main() {
	godotenv.Load()
	server.Init()
	server.Run()

}
