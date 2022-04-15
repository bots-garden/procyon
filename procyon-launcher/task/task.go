package task

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/struCoder/pidusage"
)

type Task struct {
	Id            uuid.UUID
	Name          string
	State         State
	PreviousState State
	Config        Config
	WasmRunner    WasmRunner
	HttpPort      int
	StartTime     time.Time
	FinishTime    time.Time
	CPU           float64
	Memory        float64
}

// TODO: get the size of the wasm file (for that query the registry)
// -> "WasmRegistryUrl": "https://localhost:9999/hi/hi.wasm"
// -> add some routes(api) to the registry

//TODO: log the time
// add a method to get the metrics of the process
// ref: https://github.com/struCoder/pidusage
// ??? should I use a coroutine inside the task to refresh the data ?

func (tsk *Task) ChangeState(newState State) {
	tsk.PreviousState = tsk.State
	tsk.State = newState
}

// !!! it's only for a Pending task
func getPIDUsage(tsk *Task) {
	for {
		sysinfo, _ := pidusage.GetStat(tsk.WasmRunner.Process.Pid)
		tsk.CPU = sysinfo.CPU
		tsk.Memory = sysinfo.CPU
		time.Sleep(1 * time.Second)
	}

}

func (tsk *Task) StartWasmRunner() *WasmRunnerResult {

	tsk.WasmRunner = WasmRunner{
		RunnerConfig: &tsk.Config,
		RunnerId:     uuid.New(), //!!! useful or not?
	}

	//tsk.RunnerId = tsk.WasmRunner.RunnerId

	wasmRunnerResult := tsk.WasmRunner.Start()
	if wasmRunnerResult.Error != nil {
		log.Println("😡 [starting]", wasmRunnerResult.Error)
		return wasmRunnerResult
	}

	log.Println("🎉 runner ", tsk.WasmRunner.RunnerId, " is running with config: ", tsk.WasmRunner.RunnerConfig)

	go getPIDUsage(tsk)

	return wasmRunnerResult
}

func (tsk *Task) StopWasmRunner() *WasmRunnerResult {

	wasmRunnerResult := tsk.WasmRunner.Stop()
	if wasmRunnerResult.Error != nil {
		log.Println("😡 [stopping]", wasmRunnerResult.Error)
		return wasmRunnerResult
	}

	log.Println("👋 runner ", tsk.WasmRunner.RunnerId, " has been stopped")

	//!!! add a flag?
	tsk.WasmRunner.Remove()

	return wasmRunnerResult
}
