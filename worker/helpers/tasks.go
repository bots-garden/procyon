package helpers

import (
	"log"

	"github.com/bots-garden/procyon/settings"
	"github.com/bots-garden/procyon/task"
	"github.com/google/uuid"
)

func GetTaskInfo(tasksDb map[uuid.UUID]*task.Task, idOfTask uuid.UUID) *task.Task {

	taskWithRevision := tasksDb[idOfTask]
	
	log.Println(
		"💢💢 task:", taskWithRevision.Id,
		"revision:", taskWithRevision.Config.FunctionRevision,
		"defaul revision:", taskWithRevision.Config.DefaultRevision,
		"from runner:", taskWithRevision.WasmRunner.RunnerId)

	return taskWithRevision
}

func StopTask(tasksDb map[uuid.UUID]*task.Task, idOfTask uuid.UUID) task.Task {

	taskToStop := tasksDb[idOfTask]
	taskCopy := *taskToStop
	taskCopy.ChangeState(task.Completed)

	log.Println("💢 added task:", taskToStop.Id, "to stop runner:", taskToStop.WasmRunner.RunnerId)

	return taskCopy
}

var currentHttpPort = settings.GetSettings().Http.Start

func AddTask(taskEvent task.TaskEvent, settings settings.Settings) task.Task {

		// TODO: save the data of the function in a shared place (for reverse proxy)

		functionConfig := task.Config{
			Executor:             taskEvent.Executor,
			WasmFileName:         taskEvent.WasmFileName,
			WasmFunctionHttpPort: currentHttpPort,
			WasmRegistryUrl:      taskEvent.WasmRegistryUrl,
			FunctionName:         taskEvent.FunctionName,
			FunctionRevision:     taskEvent.FunctionRevision,
			DefaultRevision:      taskEvent.DefaultRevision,
		}
	
		// TODO: make a table of available ports
		// to handle the deletion of a function
		currentHttpPort += 1
	
		functionConfig.Initialize(settings)
	
		functionTask := task.Task{
			Id:       uuid.New(),
			Name:     "task[" + functionConfig.WasmFileName + "]",
			Config:   *&functionConfig,
			HttpPort: functionConfig.WasmFunctionHttpPort,
			State:    task.Scheduled,
		}

		return functionTask
}