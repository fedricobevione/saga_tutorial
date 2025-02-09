package main

import (
	"log"

	"github.com/federicobevione/saga_tutorial/orchestration/activities"
	"github.com/federicobevione/saga_tutorial/orchestration/workflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, "healthcare-queue", worker.Options{})

	w.RegisterWorkflow(workflow.HealthcareWorkflow)
	w.RegisterActivity(activities.VerifyPatientActivity)
	w.RegisterActivity(activities.NotifyProcedureScheduleCancellationActivity)
	w.RegisterActivity(activities.ScheduleProcedureActivity)
	w.RegisterActivity(activities.CancelProcedureScheduleActivity)
	w.RegisterActivity(activities.ReserveSuppliesActivity)
	w.RegisterActivity(activities.ReleaseReservedSuppliesActivity)
	w.RegisterActivity(activities.ProcessBillingActivity)
	w.RegisterActivity(activities.ProcessBillingCompensationActivity)
	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
