package main

import (
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func main() {
	Router = gin.Default()

	// Init Routers
	InitRouter()

	// Start serving
	Router.Run(":7777")
}
