package worker

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/bots-garden/procyon/task"
)

type Worker struct {
	Queue     queue.Queue
	TasksDb   map[uuid.UUID]*task.Task
	TaskCount int // keep track of tasks a worker has at any given time
	Name      string
}

func (worker *Worker) Metrics() {
	log.Println("I will collect metrics")
}

/*
- Pull a task of the queue
- Convert it from an interface to a task.Task type // ??? 🤔
- Retrieve the task from the workers's Db
- Check if the state transition is valid
- If the task from the queue is in a state "Scheduled", call StartTask
- If the task from the queue is in a state "Completed", call StopTask
- Else, there is an invalid transition, so retuen an error
*/
func (worker *Worker) RunTask() task.WasmRunnerResult {
	log.Println("🤖 running task...")
	tsk := worker.Queue.Dequeue() // pop a task of the worker queue

	if tsk == nil {
		log.Println("🤖 no task in the queue")
		return task.WasmRunnerResult{
			Error:  nil,
			Action: "task dequeue",
			Result: "",
		}
	}
	taskQueued := tsk.(*task.Task) // convert the ppoped task to the proper type

	var result task.WasmRunnerResult

	if task.ValidStateTransition(taskQueued.PreviousState, taskQueued.State) {

		switch taskQueued.State {
		case task.Scheduled:
			result = worker.StartTask(taskQueued)
		case task.Completed:
			result = worker.StopTask(taskQueued)
		default:
			result.Error = errors.New(("🥶 Houston, we have a problem"))
		}
	} else {
		err := fmt.Errorf("Invalid transition from %v to %v", taskQueued.PreviousState, taskQueued.State)
		result.Error = err
	}

	return result

}

/*
- Update the StartTime field on the task t.
- Create an instance of the Docker struct to talk to the Docker daemon.
- Call the Run() method on the Docker struct.
- Check if there were any errors in starting the task.
- Add the running container’s ID to the tasks t.Runtime.ContainerId field.
- Save the updated task t to the worker’s Db field.
- Return the result of the operation.
*/
func (worker *Worker) StartTask(t *task.Task) task.WasmRunnerResult {
	log.Println("🚗 starting task...")

	t.StartTime = time.Now().UTC()

	result := t.StartWasmRunner()
	if result.Error != nil {
		log.Println("😡 [worker.StartTask]", result.Error)
		t.State = task.Failed
		worker.PersistTask(t)
		return *result
	}
	t.ChangeState(task.Running)

	log.Println("🕞 [worker.StartTask]", t.StartTime, "states:", t.PreviousState, "to", t.State)

	worker.PersistTask(t)

	return *result
}

/*
- Create an instance of the Docker struct that allows us to talk to the Docker daemon using the Docker SDK.
- Call the Stop() method on the Docker struct.
- Check if there were any errors in stopping the task.
- Update the FinishTime field on the task t.
- Save the updated task t to the worker’s Db field.
- Print an informative message and return the result of the operation.
*/
func (worker *Worker) StopTask(t *task.Task) task.WasmRunnerResult {

	log.Println("🛑 stopping task...")

	result := t.StopWasmRunner()
	if result.Error != nil {
		log.Println("😡 [worker.StopTask]", result.Error)
		//t.State = task.Failed
		//return result
	}
	t.FinishTime = time.Now().UTC()

	t.ChangeState(task.Completed)

	log.Println("🕞 [worker.StopTask]", t.StartTime, "states:", t.PreviousState, "to", t.State)

	log.Println("👋 task:", t.Id, "runner stopped and removed", t.WasmRunner.RunnerId)

	worker.PersistTask(t)

	return *result

}

func (worker *Worker) AddTask(t *task.Task) {
	log.Println("⭐ [added task]", t.StartTime, "states:", t.PreviousState, "to", t.State)

	worker.Queue.Enqueue(t)
}

func (worker *Worker) PersistTask(t *task.Task) {
	worker.TasksDb[t.Id] = t
}