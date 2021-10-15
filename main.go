package main

import (
	"cookieApi/backend/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	api := server.Router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	server.RunServer()
}
