package main

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	Router.POST("/newMessage", func(ctx *gin.Context) {
		var requestBody message
		err := ctx.BindJSON(&requestBody)
		failOnError(err, "Cannot Bind Request Body")
		ctx.JSON(200, "{ 'status': 'OK' }")
	})
}
