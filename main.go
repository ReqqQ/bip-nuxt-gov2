package main

import (
	"cookieApi/backend/controller"
	"cookieApi/backend/server"
)

func main() {
	controller.GetControllers()
	server.RunServer()
}
