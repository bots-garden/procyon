package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/bots-garden/procyon/settings"
	"github.com/bots-garden/procyon/task"
)

type Api struct {
	Address string
	Port string
	Worker *Worker
	Router *chi.Mux
	Settings settings.Settings
}

type ErrResponse struct {
	HTTPStatusCode int 
	Message string
}

/* 📝 Triggered by a curl request
curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 2,
      "wasmFileName": "hello.wasm",
      "wasmFunctionHttpPort": 8081,
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hello.wasm"
    }
  ' http://localhost:9090/tasks
*/
func (a *Api) AddTaskHandler(responseWriter http.ResponseWriter, request *http.Request) {
	data := json.NewDecoder(request.Body)
	data.DisallowUnknownFields()
	taskEvent := task.TaskEvent{}
	err := data.Decode(&taskEvent)

	if err != nil {
		responseWriter.WriteHeader(400)
		
		e := ErrResponse{
			HTTPStatusCode: 400,
			Message: fmt.Sprintf("Error unmarshalling body: %v\n", err),
		}
		json.NewEncoder(responseWriter).Encode(e)
		return
	}
	// TODO: if taskEvent.WasmFunctionHttpPort empty, give a port
	// TODO: save the data of the function in a shared place (for alcor)
	
	functionConfig := task.Config{
		Executor: taskEvent.Executor,
		WasmFileName: taskEvent.WasmFileName,
		WasmFunctionHttpPort: taskEvent.WasmFunctionHttpPort,
		WasmRegistryUrl: taskEvent.WasmRegistryUrl,
	}
	functionConfig.Initialize(a.Settings)

	functionTask := task.Task{
		Id: uuid.New(),
		Name: "task[" +functionConfig.WasmFileName + "]",
		Config: *&functionConfig,
		HttpPort: functionConfig.WasmFunctionHttpPort,
		State: task.Scheduled,
	}

	a.Worker.AddTask(&functionTask)

	responseWriter.WriteHeader(201)
	json.NewEncoder(responseWriter).Encode(functionTask)

}

func (a *Api) GetTasksListHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseWriter.WriteHeader(200)
	json.NewEncoder(responseWriter).Encode(a.Worker.TasksDb) // todo: format
}

func (a *Api) StopTaskHandler(responseWriter http.ResponseWriter, request *http.Request) {
	taskId := chi.URLParam(request, "taskID")
	if taskId == "" {
		log.Println("😡 no taskId passed in request")
		responseWriter.WriteHeader(400)
	}

	idOfTask, _ := uuid.Parse(taskId)
	_, ok := a.Worker.TasksDb[idOfTask]
	if !ok {
		log.Println("😡 no task wit id:", idOfTask, "found")
		responseWriter.WriteHeader(404)
	}

	taskToStop := a.Worker.TasksDb[idOfTask]
	taskCopy := *taskToStop
	taskCopy.ChangeState(task.Completed)
	

	log.Println("💢 added task:", taskToStop.Id, "to stop runner:", taskToStop.WasmRunner.RunnerId)

	a.Worker.AddTask(&taskCopy)

	responseWriter.WriteHeader(204)
	
}

func (a *Api) InitRouter() {
	a.Settings = settings.GetSettings()
	a.Router = chi.NewRouter()
	a.Router.Route("/tasks", func(r chi.Router) {
		r.Post("/", a.AddTaskHandler)
		r.Get("/", a.GetTasksListHandler)
		r.Route("/{taskID}", func(r chi.Router) {
			r.Delete("/", a.StopTaskHandler)
		})
	})
}

func (a *Api) Start() { // Address???
	a.InitRouter()
	log.Println("🌍 Listening on " + a.Port)
	log.Fatal(http.ListenAndServe(":"+ a.Port, a.Router))
}

