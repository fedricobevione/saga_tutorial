package main

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	log.Println("Connecting to RabbitMQ...")
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()
	log.Println("Successfully connected to RabbitMQ.")

	log.Println("Opening a channel...")
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()
	log.Println("Channel opened successfully.")

	log.Println("Declaring exchange 'events'...")
	err = ch.ExchangeDeclare("events", "direct", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
	log.Println("Exchange 'events' declared successfully.")

	log.Println("Declaring queue 'PatientVerified'...")
	_, err = ch.QueueDeclare("PatientVerified", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}
	log.Println("Queue 'PatientVerified' declared successfully.")

	log.Println("Binding queue 'PatientVerified' to exchange 'events' with routing key 'PatientVerified'...")
	err = ch.QueueBind("PatientVerified", "PatientVerified", "events", false, nil)
	if err != nil {
		log.Fatalf("Failed to bind a queue: %v", err)
	}
	log.Println("Queue 'PatientVerified' bound successfully.")

	log.Println("Declaring queue 'SuppliesReserved'...")
	_, err = ch.QueueDeclare("SuppliesReserved", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}
	log.Println("Queue 'SuppliesReserved' declared successfully.")

	log.Println("Binding queue 'SuppliesReserved' to exchange 'events' with routing key 'SuppliesReserved'...")
	err = ch.QueueBind("SuppliesReserved", "SuppliesReserved", "events", false, nil)
	if err != nil {
		log.Fatalf("Failed to bind a queue: %v", err)
	}
	log.Println("Queue 'SuppliesReserved' bound successfully.")

	log.Println("Declaring queue 'SuppliesReserveFailed'...")
	_, err = ch.QueueDeclare("ReservedSuppliesReleased", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
	log.Println("Queue 'ReservedSuppliesReleased' declared successfully.")

	log.Println("Binding queue 'ReservedSuppliesReleased' to exchange 'events' with routing key 'ReservedSuppliesReleased'...")
	err = ch.QueueBind("ReservedSuppliesReleased", "ReservedSuppliesReleased", "events", false, nil)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
	log.Println("Queue 'ReservedSuppliesReleased' bound successfully.")

	log.Println("Declaring queue 'SuppliesReserveFailed'...")
	_, err = ch.QueueDeclare("ProcedureScheduled", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
	log.Println("Queue 'ProcedureScheduled' declared successfully.")

	log.Println("Binding queue 'ProcedureScheduled' to exchange 'events' with routing key 'ProcedureScheduled'...")
	err = ch.QueueBind("ProcedureScheduled", "ProcedureScheduled", "events", false, nil)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
	log.Println("Queue 'ProcedureScheduled' bound successfully.")

	log.Println("Declaring queue 'ProcedureScheduleCancelled'...")
	_, err = ch.QueueDeclare("ProcedureScheduleCancelled", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
	log.Println("Queue 'ProcedureScheduleCancelled' declared successfully.")

	log.Println("Binding queue 'ProcedureScheduleCancelled' to exchange 'events' with routing key 'ProcedureScheduleCancelled'...")
	err = ch.QueueBind("ProcedureScheduleCancelled", "ProcedureScheduleCancelled", "events", false, nil)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
	log.Println("Queue 'ProcedureScheduleCancelled' bound successfully.")

	log.Println("Declaring queue 'ProcedureScheduleFailed'...")
	_, err = ch.QueueDeclare("BillingSuccessful", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
	log.Println("Queue 'BillingSuccessful' declared successfully.")

	log.Println("Binding queue 'BillingSuccessful' to exchange 'events' with routing key 'BillingSuccessful'...")
	err = ch.QueueBind("BillingSuccessful", "BillingSuccessful", "events", false, nil)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
	log.Println("Queue 'BillingSuccessful' bound successfully.")

	log.Println("Declaring queue 'BillingFailed'...")
	_, err = ch.QueueDeclare("BillingFailed", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
	log.Println("Queue 'BillingFailed' declared successfully.")

	log.Println("Binding queue 'BillingFailed' to exchange 'events' with routing key 'BillingFailed'...")
	err = ch.QueueBind("BillingFailed", "BillingFailed", "events", false, nil)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
	log.Println("Queue 'BillingFailed' bound successfully.")

	log.Println("Declaring queue 'PendingChargesReversed'...")
	_, err = ch.QueueDeclare("PendingChargesReversed", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
	log.Println("Queue 'PendingChargesReversed' declared successfully.")

	log.Println("Binding queue 'PendingChargesReversed' to exchange 'events' with routing key 'PendingChargesReversed'...")
	err = ch.QueueBind("PendingChargesReversed", "PendingChargesReversed", "events", false, nil)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
	log.Println("Queue 'PendingChargesReversed' bound successfully.")
}
