package main

import (
	//"log"
	"log"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/bots-garden/procyon/settings"
	"github.com/bots-garden/procyon/task"
	"github.com/bots-garden/procyon/worker"
	//"gitlab.com/k33g/galago-o/worker/api"
)

func getConfigs() (*task.Config, *task.Config, *task.Config) {

	workerSettings := settings.GetSettings()
	log.Println("🎃", workerSettings)

	helloConfig := task.Config{
		Executor:             task.Galago,
		WasmFileName:         "hello.wasm",
		WasmFunctionHttpPort: 8082,
		WasmRegistryUrl:      "https://localhost:9999/hello/hello.wasm",
		Env:                  []string{},
		Args:                 []string{},
	}
	helloConfig.Initialize(workerSettings)

	heyConfig := task.Config{
		Executor:             task.Galago,
		WasmFileName:         "hey.wasm",
		WasmFunctionHttpPort: 8083,
		WasmRegistryUrl:      "https://localhost:9999/hey/hey.wasm",
		Env:                  []string{},
		Args:                 []string{},
	}
	heyConfig.Initialize(workerSettings)

	hiConfig := task.Config{
		Executor:             task.Sat,
		WasmFileName:         "",
		WasmFunctionHttpPort: 8084,
		WasmRegistryUrl:      "https://localhost:9999/hi/hi.wasm",
		Env:                  []string{},
		Args:                 []string{},
	}
	hiConfig.Initialize(workerSettings)

	return &helloConfig, &heyConfig, &hiConfig
}

func withWorkers(delayBeforeStop time.Duration) {

	helloConfig, heyConfig, hiConfig := getConfigs()

	helloTask := task.Task{
		Id:       uuid.New(),
		Name:     "hello-function-task",
		Config:   *helloConfig,
		HttpPort: 8082,
		State:    task.Scheduled, // 🖐
	}

	heyTask := task.Task{
		Id:       uuid.New(),
		Name:     "hey-function-task",
		Config:   *heyConfig,
		HttpPort: 8083,
		State:    task.Scheduled, // 🖐
	}

	hiTask := task.Task{
		Id:       uuid.New(),
		Name:     "hi-function-task",
		Config:   *hiConfig,
		HttpPort: 8084,
		State:    task.Scheduled, // 🖐
	}

	worker := worker.Worker{
		Queue:     *queue.New(),
		TasksDb:   make(map[uuid.UUID]*task.Task),
		TaskCount: 0,
		Name:      "my-worker",
	}

	log.Println("🧡 starting task")

	worker.AddTask(&helloTask)
	worker.AddTask(&heyTask)
	worker.AddTask(&hiTask)

	/*
		result := worker.RunTask()
		if result.Error != nil {
			panic(result.Error)
		}
		helloTask.RunnerId = result.RunnerId // put this in the task

		log.Println("✅ task: ", helloTask.Id, " is running in runner: ", helloTask.RunnerId)
		log.Println("🔥 task states: ", helloTask.PreviousState, helloTask.State)
	*/

	worker.RunTask()
	worker.RunTask()
	worker.RunTask()

	log.Println("📦", worker.TasksDb)

	worker.RunTask() // No Task in the queue

	time.Sleep(delayBeforeStop)

	log.Println("💙 stopping task")

	helloTask.ChangeState(task.Completed)
	heyTask.ChangeState(task.Completed)
	hiTask.ChangeState(task.Completed)

	//log.Println("🔥🔥 task states: ", helloTask.PreviousState, helloTask.State)

	worker.AddTask(&helloTask)
	worker.AddTask(&heyTask)
	worker.AddTask(&hiTask)

	/*
		result = worker.RunTask()
		if result.Error != nil {
			panic(result.Error)
		}
	*/
	worker.RunTask()
	worker.RunTask()
	worker.RunTask()

}

func runTasks(worker *worker.Worker) {
	for {
		if worker.Queue.Len() != 0 {
			result := worker.RunTask()
			if result.Error != nil {
				log.Println("😡 error running task:", result.Error)
			}
		} else {
			log.Println("🤖 no task to process currently")
		}
		log.Println("😴 sleeping for 10 seconds")
		time.Sleep(10 * time.Second)
	}
}

func main() {

	// start the registry before

	//withWorkers(time.Second * 15)

	
		httpPort := settings.GetEnv("WASM_WORKER_PORT", "9090")

		log.Println("🚀 starting Wasm worker...")

		wasmWorker := worker.Worker{
			Queue: *queue.New(),
			TasksDb: make(map[uuid.UUID]*task.Task),
			TaskCount: 0,
			Name: "my-wasm-worker",
		}


		api := worker.Api{
			Address: "", Port: httpPort, Worker: &wasmWorker,
		}

		go runTasks(&wasmWorker)

		api.Start()
	

}
