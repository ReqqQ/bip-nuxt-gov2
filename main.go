package main

import (
	"cookieApi/backend/controller"
	"cookieApi/backend/database"
	"cookieApi/backend/server"
)

func main() {
	database.ConnectToDB()
	controller.GetControllers()
	server.RunServer()
}
