package main

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	q, _ := ch.QueueDeclare(
		"test",
		false,
		false,
		false,
		false,
		nil,
	)

	// 4. 发布消息
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, _ := json.Marshal("hi")
	_ = ch.PublishWithContext(
		ctx,
		"",     // exchange: 空字符串表示使用默认 exchange
		q.Name, // routing key: 这里直接用队列名
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
