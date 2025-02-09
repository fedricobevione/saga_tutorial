package main

import (
	"context"
	"log"

	"github.com/fedricobevione/saga_tutorial/orchestration/workflow"
	"go.temporal.io/sdk/client"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	options := client.StartWorkflowOptions{
		ID:        "my-id",
		TaskQueue: "healthcare-queue",
	}

	ctx := context.Background()

	// Start the workflow
	we, err := c.ExecuteWorkflow(ctx, options, workflow.HealthcareWorkflow)
	if err != nil {
		log.Fatalln("error starting TransferMoney workflow", err)
	}

	// wait for workflow completion
	err = we.Get(ctx, nil)
	if err != nil {
		log.Fatalln("workflow error", err)
	}
}
