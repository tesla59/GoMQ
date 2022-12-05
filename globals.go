package main

import (
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	Router *gin.Engine
	Conn   *amqp.Connection
	Ch     *amqp.Channel
	Q      amqp.Queue
	Err    error
)
