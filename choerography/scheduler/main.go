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

	go func() {
		fmt.Println("[SchedulerService] Waiting for events...")
		// once a patient is verified, the scheduler service schedules the procedure
		msgs, err := common.ConsumeEvent(ch, "PatientVerified")
		if err != nil {
			log.Fatalf("Failed to consume event: %v", err)
		}

		for range msgs {
			fmt.Println("[SchedulerService] Processing event: PatientVerified")
			// Simulate work
			if err := scheduleProcedure(); err != nil {
				// if scheduling fails, trigger compensation
				common.PublishEvent(ch, "events", "ProcedureScheduleFailed", "Failed to schedule procedure")
				fmt.Println("[SchedulerService] Compensation triggered: ProcedureScheduleFailed")
			} else {
				// if scheduling is successful, publish an event
				common.PublishEvent(ch, "events", "ProcedureScheduled", "Procedure scheduled successfully")
				fmt.Println("[SchedulerService] Event published: ProcedureScheduled")
			}
		}
	}()

	go func() {
		fmt.Println("[SchedulerService] Waiting for compensation events...")
		// if billing fails, the inventory service releases the reserved supplies
		// this is the compensation for the inventory service
		// the scheduler service listens for this event and cancels the procedure schedule
		msgs, err := common.ConsumeEvent(ch, "ReservedSuppliesReleased")
		if err != nil {
			log.Fatalf("Failed to consume event: %v", err)
		}

		for range msgs {
			fmt.Println("[SchedulerService] Processing event: ReservedSuppliesReleased")
			// Simulate work
			err := cancelProcedureSchedule()
			if err != nil {
				log.Fatalf("Failed to cancel procedure schedule: %v", err)
			}

			// if cancelling the procedure schedule is successful, publish an event
			// this event is consumed by the patient service to notify the patient
			common.PublishEvent(ch, "events", "ProcedureScheduleCancelled", "Procedure schedule cancelled")
		}
	}()

	select {}
}

func scheduleProcedure() error {
	// Simulate scheduling logic
	fmt.Println("Step 2: Scheduling procedure...")
	return nil // or return an error to simulate failure
}

func cancelProcedureSchedule() error {
	// Simulate cancellation logic
	fmt.Println("Compensation 2: Cancel procedure schedule.")
	return nil
}
