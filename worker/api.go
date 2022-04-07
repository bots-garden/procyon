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

var currentHttpPort = settings.GetSettings().Http.Start

/* 📝 Triggered by a curl request
curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 1,
      "wasmFileName": "hello.wasm",
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hello.wasm",
      "functionName": "hello",
      "functionRevision": "first",
      "defaultRevision": true
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

	// TODO: save the data of the function in a shared place (for alcor)
	// !!! Right now I use an http service

	functionConfig := task.Config{
		Executor: taskEvent.Executor,
		WasmFileName: taskEvent.WasmFileName,
		//WasmFunctionHttpPort: taskEvent.WasmFunctionHttpPort,
		WasmFunctionHttpPort: currentHttpPort,
		WasmRegistryUrl: taskEvent.WasmRegistryUrl,
		FunctionName: taskEvent.FunctionName,
		FunctionRevision: taskEvent.FunctionRevision,
		DefaultRevision: taskEvent.DefaultRevision,

	}

	// TODO: make a table of available ports
	// to handle the deletion of a function
	currentHttpPort+=1

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

type FunctionRecord struct {
	WasmFunctionHttpPort int
	TaskId uuid.UUID
	DefaultRevision bool
}

/* This method/route is called by the reverse proxy

{

    "hello-first": {
        "TaskId": "536d4475-7e22-4437-991a-740a8aef290d",
        "WasmFunctionHttpPort": 3000
				"DefaultRevision", true
    },
    "hello-orange": {
        "TaskId": "d9dd4be8-a377-4a6f-b87f-7e74cebd9483",
        "WasmFunctionHttpPort": 3004
		}
    "hey-first": {
        "TaskId": "2990d84b-2832-46b9-8c97-c5e39928da96",
        "WasmFunctionHttpPort": 3001
    }
}
*/
func (a *Api) GetFunctionsListHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseWriter.WriteHeader(200)

	var functionsMap map[string]FunctionRecord
	functionsMap = make(map[string]FunctionRecord)

	// parse map
	for key, element := range a.Worker.TasksDb {

		functionsMap[element.Config.FunctionName+"-"+element.Config.FunctionRevision] = FunctionRecord{
			WasmFunctionHttpPort: element.Config.WasmFunctionHttpPort,
			TaskId: key,
			DefaultRevision: element.Config.DefaultRevision,
		}
		/*
		if element.Config.DefaultRevision == true {
			functionsMap[element.Config.FunctionName+"-"+"*"] = FunctionRecord{
				WasmFunctionHttpPort: element.Config.WasmFunctionHttpPort,
				TaskId: key,
			}
		}
		*/

	}

	json.NewEncoder(responseWriter).Encode(functionsMap)
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
//TODO: 🌺 gardening to do
func (a *Api) SwitchTaskRevisionHandler(responseWriter http.ResponseWriter, request *http.Request) {
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

	taskWithRevision := a.Worker.TasksDb[idOfTask]
	taskWithRevision.Config.DefaultRevision = !taskWithRevision.Config.DefaultRevision
	
	//taskCopy := *taskWithRevision
	//taskCopy.ChangeState(task.Completed)
	
	log.Println(
		"💢💢 task:", taskWithRevision.Id, 
		"revision:",taskWithRevision.Config.FunctionRevision,
		"defaul revision:", taskWithRevision.Config.DefaultRevision, 
		"from runner:", taskWithRevision.WasmRunner.RunnerId)

	//a.Worker.AddTask(&taskCopy)

	responseWriter.WriteHeader(200)
	
}

func (a *Api) TaskInfoHandler(responseWriter http.ResponseWriter, request *http.Request) {

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

	taskWithRevision := a.Worker.TasksDb[idOfTask]
	
	
	//taskCopy := *taskWithRevision
	//taskCopy.ChangeState(task.Completed)
	
	log.Println(
		"💢💢 task:", taskWithRevision.Id, 
		"revision:",taskWithRevision.Config.FunctionRevision,
		"defaul revision:", taskWithRevision.Config.DefaultRevision, 
		"from runner:", taskWithRevision.WasmRunner.RunnerId)

	//a.Worker.AddTask(&taskCopy)

	responseWriter.WriteHeader(200)
	json.NewEncoder(responseWriter).Encode(taskWithRevision)
	
}

func (a *Api) SwitchRevision(responseWriter http.ResponseWriter, request *http.Request) {
	// r.Put("/{functionName}/{functionRevision}", a.SwitchRevision)

	functionName := chi.URLParam(request, "functionName")
	if functionName == "" {
		log.Println("😡 no functionName passed in request")
		responseWriter.WriteHeader(400)
	}
	functionRevision := chi.URLParam(request, "functionRevision")
	if functionRevision == "" {
		log.Println("😡 no functionRevision passed in request")
		responseWriter.WriteHeader(400)
	}
	switchRevisionValue := chi.URLParam(request, "switch")
	if switchRevisionValue == "" {
		log.Println("😡 no switchRevisionValue (on/off) passed in request")
		responseWriter.WriteHeader(400)
	}

	// search the id of the related task
	for key, element := range a.Worker.TasksDb {

		if element.Config.FunctionName==functionName && element.Config.FunctionRevision==functionRevision {

			if switchRevisionValue == "on" {
				element.Config.DefaultRevision = true
				} else {
					if switchRevisionValue == "off" {
						element.Config.DefaultRevision = false
					} else {
						// no change
						log.Println("😡 bad value for switchRevisionValue:", switchRevisionValue)
					}
				}

			log.Println(
				"🔴🤖 Key:", key, "=>", 
				"TaskId:", element.Id, 
				"Function", element.Config.FunctionName, element.Config.FunctionRevision, element.Config.DefaultRevision)

				break;
		}
		

  }

	responseWriter.WriteHeader(200)
	
}


func (a *Api) InitRouter() {
	a.Settings = settings.GetSettings()
	a.Router = chi.NewRouter()
	a.Router.Route("/tasks", func(r chi.Router) {
		r.Post("/", a.AddTaskHandler)
		r.Get("/", a.GetTasksListHandler)

		//r.Put("/{functionName}/{functionRevision}", a.SwitchRevision) // change the status of the revision 

		r.Route("/{taskID}", func(r chi.Router) {
			r.Delete("/", a.StopTaskHandler)
			// change the status of the revision TODO: test
			r.Put("/",a.SwitchTaskRevisionHandler)  
			r.Get("/", a.TaskInfoHandler)
		})
	})
	a.Router.Route("/functions", func(r chi.Router) {
		r.Get("/", a.GetFunctionsListHandler)
	})

	a.Router.Route("/revisions", func(r chi.Router) {
		r.Put("/{functionName}/{functionRevision}/default/{switch}", a.SwitchRevision) // change the status of the revision 
	})


}

func (a *Api) Start() { // Address???
	a.InitRouter()
	log.Println("🌍 Listening on " + a.Port)
	log.Fatal(http.ListenAndServe(":"+ a.Port, a.Router))
}

