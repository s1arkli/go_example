package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	msgs, _ := ch.Consume(
		"test",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	//接受消息
	for msg := range msgs {
		fmt.Println(string(msg.Body))
	}
}
