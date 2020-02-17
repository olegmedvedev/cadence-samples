package main

import (
	"cadence-samples/internal"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
	"time"
)

const sampleWorkflowName = "SampleWorkflow"

// This is registration process where you register all your workflow handlers.
func init() {
	workflow.RegisterWithOptions(SampleWorkflow, workflow.RegisterOptions{Name: sampleWorkflowName})
}

// SampleWorkflow workflow decider
func SampleWorkflow(ctx workflow.Context) (result string, err error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: 30 * time.Second,
		StartToCloseTimeout:    time.Minute,
		TaskList:               internal.DefinerWorkerTaskList,
	}
	ctx1 := workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow started.")

	err = workflow.ExecuteActivity(ctx1, internal.TestActivityName).Get(ctx1, nil)
	if err != nil {
		logger.Error("Failed to do activity", zap.Error(err))
		return "", err
	}

	logger.Info("Workflow completed.")
	return "COMPLETED", nil
}
