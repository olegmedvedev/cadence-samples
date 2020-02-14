# Cadence Sample

```

1. Run definer
2. Run worker -m=worker
3. Run worker -m=trigger

```

#### Error:
```
2020-02-14T12:43:34.543+0200    WARN    internal/internal_task_pollers.go:332   Failed to process decision task.        {"Domain": "samples-domain", "TaskList": "testGroup", "WorkerID": "13688@Lim-MedvedevO@testGroup", "WorkflowType": "SampleWorkflow", "WorkflowID"
: "TEST-17184b1e-2607-4ee0-a5e7-84d495eec314", "RunID": "08e1acba-0e9f-4db4-b563-2473c62bddd6", "error": "unable to find workflow type: SampleWorkflow. Supported types: [SampleWorkflowDefiner]"}


```