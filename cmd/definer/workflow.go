package main

import (
	"context"
	"fmt"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
	"time"
)

const (
	ApplicationName = "testGroup"
)

// This is registration process where you register all your workflow handlers.
func init() {
	workflow.RegisterWithOptions(SampleWorkflowDefiner, workflow.RegisterOptions{Name: "SampleWorkflowDefiner"})

	activity.RegisterWithOptions(testActivity, activity.RegisterOptions{Name: "testActivity"})
	fmt.Print("Registered testActivity")
}

// SampleWorkflowDefiner workflow decider
func SampleWorkflowDefiner(ctx workflow.Context) (result string, err error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow started.")

	workflow.Sleep(ctx, time.Second * 10)

	logger.Info("Workflow completed.")
	return "COMPLETED", nil
}

func testActivity(ctx context.Context) error {
	logger := activity.GetLogger(ctx)
	logger.Info("test activity started")
	time.Sleep(time.Second * 30)
	logger.Info("test activity done")
	return nil
}
