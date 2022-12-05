package main

import (
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"net/http"
)

var (
	Router *gin.Engine
	Server *http.Server
	Conn   *amqp.Connection
	Ch     *amqp.Channel
	Q      amqp.Queue
	Err    error
)
