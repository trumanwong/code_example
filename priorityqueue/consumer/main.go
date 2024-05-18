package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/trumanwong/go-tools/mq"
	"os"
	"time"
)

func main() {
	arguments := make(amqp.Table)
	arguments["x-max-priority"] = int64(9)
	mq.NewRabbitMQ(&mq.Options{
		Name:          os.Getenv("RABBITMQ_QUEUE_NAME"),
		Addr:          os.Getenv("RABBITMQ_ADDR"),
		PrefetchCount: 1,
		PrefetchSize:  0,
		Global:        false,
		Consume: func(msgs <-chan amqp.Delivery) {
			for d := range msgs {
				fmt.Println("receive data: ", string(d.Body))
				time.Sleep(10 * time.Second)
				_ = d.Ack(false)
			}
		},
		Arguments: arguments,
	})
	select {}
}
