package main

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"example/6.rabbitMQ/queue"
)

var (
	queueName_01 = "test"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	queue.NewQueue(ch, queueName_01)

	//  发布消息
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < 5; i++ {
		_ = ch.PublishWithContext(
			ctx,
			"",
			queueName_01, // 队列名称
			false,
			false,
			amqp.Publishing{
				Body: []byte(fmt.Sprintf("第%d次消息", i+1)),
			},
		)

		time.Sleep(1 * time.Second)
	}

}
