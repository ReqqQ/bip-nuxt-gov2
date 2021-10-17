package controller

import (
	"cookieApi/backend/server"
)

func GetControllers() {
	api := server.Router.Group("/api")
	{
		api.GET("/", GetUser)
	}
}
