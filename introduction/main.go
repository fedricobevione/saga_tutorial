package main

import (
	"fmt"
	"log"
)

// Define step and compensation types
type Step func() error
type Compensation func()

// Saga structure
type Saga struct {
	steps         []Step
	compensations []Compensation
}

func (s *Saga) AddStep(step Step, compensation Compensation) {
	s.steps = append(s.steps, step)
	s.compensations = append([]Compensation{compensation}, s.compensations...)
}

func (s *Saga) Execute() {
	for i, step := range s.steps {
		if err := step(); err != nil {
			log.Printf("Step %d failed: %v. Rolling back...\n", i+1, err)
			for _, compensation := range s.compensations {
				compensation()
			}
			return
		}
	}
	fmt.Println("Transaction completed successfully!")
}

func main() {
	saga := &Saga{}

	// Step 1: Verify patient and insurance details
	saga.AddStep(
		func() error {
			fmt.Println("Step 1: Verifying patient and insurance details...")
			// Simulate success
			return nil
		},
		func() { fmt.Println("Compensation 1: Notify patient of verification failure.") },
	)

	// Step 2: Schedule procedure
	saga.AddStep(
		func() error {
			fmt.Println("Step 2: Scheduling procedure...")
			// Simulate success
			return nil
		},
		func() { fmt.Println("Compensation 2: Cancel procedure schedule.") },
	)

	// Step 3: Reserve medical supplies
	saga.AddStep(
		func() error {
			fmt.Println("Step 3: Reserving medical supplies...")
			// Simulate success
			return nil
		},
		func() { fmt.Println("Compensation 3: Release reserved supplies.") },
	)

	// Step 4: Process billing
	saga.AddStep(
		func() error {
			fmt.Println("Step 4: Processing billing...")
			// Simulate failure
			return fmt.Errorf("billing service unavailable")
		},
		func() { fmt.Println("Compensation 4: Reverse any pending charges.") },
	)

	// Execute the saga
	saga.Execute()
}
