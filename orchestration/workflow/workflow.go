// workflow.go
package workflow

import (
	"time"

	"github.com/fedricobevione/saga_tutorial/orchestration/activities"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/multierr"
)

func HealthcareWorkflow(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting Healthcare Workflow")

	// Define retry options for activities
	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute,
		MaximumAttempts:    5,
	}
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retryPolicy,
	})

	// Step 1: Verify Patient
	logger.Info("Verifying Patient")
	err := workflow.ExecuteActivity(ctx, activities.VerifyPatientActivity).Get(ctx, nil)
	if err != nil {
		logger.Error("Patient verification failed", "Error", err)
		return err
	}

	// compensation for patient verification
	defer func() {
		if err != nil {
			errCompensation := workflow.ExecuteActivity(ctx, activities.NotifyProcedureScheduleCancellationActivity).Get(ctx, nil)
			err = multierr.Append(err, errCompensation)
		}
	}()

	// Step 2: Schedule Procedure
	logger.Info("Scheduling Procedure")
	err = workflow.ExecuteActivity(ctx, activities.ScheduleProcedureActivity).Get(ctx, nil)
	if err != nil {
		logger.Error("Procedure scheduling failed", "Error", err)
		return err
	}

	// compensation for procedure scheduling
	defer func() {
		if err != nil {
			errCompensation := workflow.ExecuteActivity(ctx, activities.CancelProcedureScheduleActivity).Get(ctx, nil)
			err = multierr.Append(err, errCompensation)
		}
	}()

	// Step 3: Reserve Supplies
	logger.Info("Reserving Supplies")
	err = workflow.ExecuteActivity(ctx, activities.ReserveSuppliesActivity).Get(ctx, nil)
	if err != nil {
		logger.Error("Supply reservation failed", "Error", err)
		return err
	}

	// compensation for supply reservation
	defer func() {
		if err != nil {
			errCompensation := workflow.ExecuteActivity(ctx, activities.ReleaseReservedSuppliesActivity).Get(ctx, nil)
			err = multierr.Append(err, errCompensation)
		}
	}()

	// Step 4: Process Billing
	logger.Info("Processing Billing")
	err = workflow.ExecuteActivity(ctx, activities.ProcessBillingActivity).Get(ctx, nil)
	if err != nil {
		logger.Error("Billing failed", "Error", err)
		return err
	}

	// compensation for billing
	defer func() {
		if err != nil {
			errCompensation := workflow.ExecuteActivity(ctx, activities.ProcessBillingCompensationActivity).Get(ctx, nil)
			err = multierr.Append(err, errCompensation)
		}
	}()

	logger.Info("Healthcare Workflow completed successfully")
	return nil
}
