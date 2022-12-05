package main

import (
	"context"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
	"net/http"
	"log"
)

func InitRouter() {
	Router = gin.Default()

	Server = &http.Server{
		Addr:    ":7777",
		Handler: Router,
	}

	go func() {
		// service connections
		if err := Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	Routes()
}

func Routes() {
	Router.POST("/newMessage", func(ctx *gin.Context) {
		var requestBody message
		err := ctx.BindJSON(&requestBody)
		failOnError(err, "Cannot Bind Request Body")

		Q, Err = Ch.QueueDeclare(
			requestBody.Queue, // name
			false,             // durable
			false,             // delete when unused
			false,             // exclusive
			false,             // no-wait
			nil,               // arguments
		)
		failOnError(Err, "Failed to declare a queue")

		c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		Err = Ch.PublishWithContext(c,
			"",     // exchange
			Q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(requestBody.Body),
			})
		failOnError(err, "Failed to publish a message")

		defer ctx.JSON(200, "{ 'status': 'OK' }")
	})
}
