package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SetupMQ() {
	Conn, Err = amqp.Dial("amqp://guest:guest@141.148.198.149:5672/")
	failOnError(Err, "Failed to connect to RabbitMQ")
	Ch, Err = Conn.Channel()
	failOnError(Err, "Failed to open a channel")
}

func Exit() {
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("##### Shutting Down Server ######")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer Conn.Close()
	defer Ch.Close()
	if err := Server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("Server Exited")
}
