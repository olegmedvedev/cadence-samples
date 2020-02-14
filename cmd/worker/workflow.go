package main

import (
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
	"time"
)

const (
	ApplicationName = "testGroup"
)

// This is registration process where you register all your workflow handlers.
func init() {
	workflow.RegisterWithOptions(SampleWorkflow, workflow.RegisterOptions{Name: "SampleWorkflow"})
}

// SampleWorkflow workflow decider
func SampleWorkflow(ctx workflow.Context) (result string, err error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: 30 * time.Second,
		StartToCloseTimeout:    1 * time.Minute,
	}
	ctx1 := workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow started.")

	err = workflow.ExecuteActivity(ctx1, "testActivity").Get(ctx1, nil)
	if err != nil {
		logger.Error("Failed to do activity", zap.Error(err))
		return "", err
	}

	logger.Info("Workflow completed.")
	return "COMPLETED", nil
}
