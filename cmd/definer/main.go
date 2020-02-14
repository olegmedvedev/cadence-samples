package main

import (
	"github.com/uber-common/cadence-samples/cmd/samples/common"
	"go.uber.org/cadence/worker"
)

func startWorkers(h *common.SampleHelper) {
	workerOptions := worker.Options{
		MetricsScope: h.Scope,
		Logger:       h.Logger,
	}
	h.StartWorkers(h.Config.DomainName, ApplicationName, workerOptions)
}

func main() {
	var h common.SampleHelper
	h.SetupServiceConfig()
	startWorkers(&h)
	select {}
}
