package queue

import "github.com/rabbitmq/amqp091-go"

func NewQueue(ch *amqp091.Channel, name string) (amqp091.Queue, error) {
	return ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
}
