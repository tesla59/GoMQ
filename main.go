package main

import (
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	Router = gin.Default()

	// Init Routers
	InitRouter()

	Conn, Err = amqp.Dial("amqp://guest:guest@141.148.198.149:5672/")
	failOnError(Err, "Failed to connect to RabbitMQ")
	defer Conn.Close()

	Ch, Err = Conn.Channel()
	failOnError(Err, "Failed to open a channel")
	defer Ch.Close()

	// Start serving
	Router.Run(":7777")
}
