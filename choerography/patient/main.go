package main

import (
	"fmt"
	"log"

	"github.com/fedricobevione/saga_tutorial/choeraography/common"
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
		fmt.Println("[PatientService] Waiting for events...")
		// once a procedure  schedule is cancelled, the patient service notifies the patient
		// this is the compensation for the patient service
		msgs, err := common.ConsumeEvent(ch, "ProcedureScheduleCancelled")
		if err != nil {
			log.Fatalf("Failed to consume event: %v", err)
		}

		for range msgs {
			fmt.Println("[PatientService] Processing event: ProcedureScheduleCancelled")
			if err := notifyProcedureScheduleCancellation(); err != nil {
				log.Fatalf("Failed to notify patient of verification failure: %v", err)
			}
		}
	}()

	go func() {
		// once billing is successful, the patient service can notify the patient
		// this is the happy path for the patient service
		msgs, err := common.ConsumeEvent(ch, "BillingSuccessful")
		if err != nil {
			log.Fatalf("Failed to consume event: %v", err)
		}

		for range msgs {
			fmt.Println("[PatientService] Processing event: BillingSuccessful - Conclusion: Patient verified successfully")
		}
	}()

	// publish an event to start the process
	common.PublishEvent(ch, "events", "PatientVerified", "Patient details verified")
	fmt.Println("[PatientService] Event published: PatientVerified")

	select {}
}

func notifyProcedureScheduleCancellation() error {
	// Simulate logic
	fmt.Println("Compensation 1: Notify patient of verification failure.")
	return nil
}
