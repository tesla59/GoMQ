package main

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func main() {
	router = gin.Default()

	router.POST("/newMessage", func(ctx *gin.Context) {
		var requestBody message
		err := ctx.BindJSON(&requestBody)
		failOnError(err, "Cannot Bind Request Body")
		ctx.JSON(200, "{ 'status': 'OK' }")
	})

	router.Run(":7777")
}
