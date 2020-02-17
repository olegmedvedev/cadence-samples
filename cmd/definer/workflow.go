package main

import (
	"cadence-samples/internal"
	"context"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
	"time"
)

// This is registration process where you register all your workflow handlers.
func init() {
	workflow.RegisterWithOptions(sampleWorkflowDefiner, workflow.RegisterOptions{Name: "SampleWorkflowDefiner"})
	activity.RegisterWithOptions(testActivity, activity.RegisterOptions{Name: internal.TestActivityName})
}

// SampleWorkflowDefiner workflow decider
func sampleWorkflowDefiner(ctx workflow.Context) (result string, err error) {
	workflow.GetLogger(ctx).Info("SampleWorkflowDefiner")
	return "COMPLETED", nil
}

func testActivity(ctx context.Context) error {
	logger := activity.GetLogger(ctx)
	logger.Info("test activity started")
	time.Sleep(time.Second * 10)
	logger.Info("test activity done")
	return nil
}
