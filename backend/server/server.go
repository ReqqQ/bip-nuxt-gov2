package server

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

const serverPort = ":3000"
const serverPrefix = "/"
const frontEndFolder = "./frontend"

var Router = gin.Default()

func RunServer() {
	runFrontEnd()
	Router.Run(serverPort)
}
func runFrontEnd() {
	Router.Use(static.Serve(serverPrefix, static.LocalFile(frontEndFolder, true)))
}
