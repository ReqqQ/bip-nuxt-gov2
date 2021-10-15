package server

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Server struct {}

func runServer()  {
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./frontend", true)))
	router.Run(":3000")
}