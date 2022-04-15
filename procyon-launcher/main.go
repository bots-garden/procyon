package main

import (
	//"log"
	"log"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/bots-garden/procyon/procyon-launcher/settings"
	"github.com/bots-garden/procyon/procyon-launcher/task"
	"github.com/bots-garden/procyon/procyon-launcher/worker"
	//"gitlab.com/k33g/galago-o/worker/api"
)


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
		log.Println("😴 sleeping for 5 seconds")
		time.Sleep(5 * time.Second)
	}
}

func main() {

	// start the registry before

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
