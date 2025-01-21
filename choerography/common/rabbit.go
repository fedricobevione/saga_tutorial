package common

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func PublishEvent(ch *amqp091.Channel, exchange, routingKey, message string) {
	err := ch.Publish(
		exchange, routingKey, false, false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	fmt.Printf("[Event] Published: %s - %s\n", routingKey, message)
}

func ConsumeEvent(ch *amqp091.Channel, queue string) (<-chan amqp091.Delivery, error) {
	msgs, err := ch.Consume(
		queue, "", true, false, false, false, nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to register a consumer: %w", err)
	}
	return msgs, nil
}
