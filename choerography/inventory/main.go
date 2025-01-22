package main

import (
	"fmt"
	"log"

	"github.com/fedricobevione/saga_tutorial/choerography/common"
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

	go func() {
		fmt.Println("[InventoryService] Waiting for events...")
		// once a procedure is scheduled, the inventory service reserves the supplies
		msgs, err := common.ConsumeEvent(ch, "ProcedureScheduled")
		if err != nil {
			log.Fatalf("Failed to consume event: %v", err)
		}

		for range msgs {
			fmt.Println("[InventoryService] Processing event: ProcedureScheduled")
			// Simulate work
			if err := reserveSupplies(); err != nil {
				// if reservation fails, trigger compensation
				common.PublishEvent(ch, "events", "SuppliesReserveFailed", "Failed to reserve supplies")
				fmt.Println("[InventoryService] Compensation triggered: ProcedureScheduleFailed")
			} else {
				// if reservation is successful, publish an event
				common.PublishEvent(ch, "events", "SuppliesReserved", "Supplies reserved")
				fmt.Println("[InventoryService] Event published: SuppliesReserved")
			}
		}
	}()

	go func() {
		fmt.Println("[InventoryService] Waiting for compensation events...")
		// if billing fails, the inventory service releases the reserved supplies
		// this is the compensation for the inventory service
		msgs, err := common.ConsumeEvent(ch, "BillingFailed")
		if err != nil {
			log.Fatalf("Failed to consume event: %v", err)
		}

		for range msgs {
			fmt.Println("[InventoryService] Processing event: BillingFailed")
			// Simulate work
			err := releaseReservedSupplies()
			if err != nil {
				log.Fatalf("Failed to release reserved supplies: %v", err)
			}

			// if releasing reserved supplies is successful, publish an event
			// this event is consumed by the scheduler service to cancel the procedure schedule
			common.PublishEvent(ch, "events", "ReservedSuppliesReleased", "Reserved supplies released")
		}
	}()

	select {}
}

func reserveSupplies() error {
	// Simulate scheduling logic
	fmt.Println("Step 3: Reserving medical supplies...")
	return nil // or return an error to simulate failure
}

func releaseReservedSupplies() error {
	// Simulate cancellation logic
	fmt.Println("Compensation 3: Release reserved supplies.")
	return nil
}
