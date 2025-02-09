package main

import (
	"fmt"
	"log"

	"github.com/federicobevione/saga_tutorial/choerography/common"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	fmt.Println("[BillingService] Waiting for events...")
	// once supplies are reserved, the billing service processes the billing
	msgs, err := common.ConsumeEvent(ch, "SuppliesReserved")
	if err != nil {
		log.Fatalf("Failed to consume event: %v", err)
	}

	for range msgs {
		fmt.Println("[BillingService] Processing event: SuppliesReserved")
		// Simulate work
		if err := processBilling(); err != nil {
			if err := processBillingcompensation(); err != nil {
				log.Fatalf("Failed to reverse pending charges: %v", err)
			}
			// if billing fails, trigger compensation
			// this is the compensation for the inventory service
			common.PublishEvent(ch, "events", "BillingFailed", "Failed to reserve supplies")
			fmt.Println("[BillingService] Compensation triggered: BillingFailed")
		} else {
			// if billing is successful, publish an event
			common.PublishEvent(ch, "events", "BillingSuccessful", "Billing successful")
			fmt.Println("[BillingService] Event published: BillingSuccessful")
		}
	}

	select {}
}

func processBilling() error {
	// Simulate logic
	fmt.Println("Step 4: Processing billing...")
	return fmt.Errorf("billing service unavailable")
}

func processBillingcompensation() error {
	// Simulate logic
	fmt.Println("Compensation 4: Reverse any pending charges.")
	return nil
}
