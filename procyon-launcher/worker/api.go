package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bots-garden/procyon/procyon-launcher/settings"
	"github.com/bots-garden/procyon/procyon-launcher/task"
	"github.com/bots-garden/procyon/procyon-launcher/worker/helpers"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type Api struct {
	Address  string
	Port     string
	Worker   *Worker
	Router   *chi.Mux
	Settings settings.Settings
}

type ErrResponse struct {
	HTTPStatusCode int
	Message        string
}

var currentHttpPort = settings.GetSettings().Http.Start

func (a *Api) AddTaskHandler(responseWriter http.ResponseWriter, request *http.Request) {

	data := json.NewDecoder(request.Body)
	data.DisallowUnknownFields()
	taskEvent := task.TaskEvent{}
	err := data.Decode(&taskEvent)

	if err != nil {
		responseWriter.WriteHeader(400)

		e := ErrResponse{
			HTTPStatusCode: 400,
			Message:        fmt.Sprintf("Error unmarshalling body: %v\n", err),
		}
		json.NewEncoder(responseWriter).Encode(e)
		return
	}

	functionTask := helpers.AddTask(taskEvent, a.Settings)

	a.Worker.AddTask(&functionTask)

	responseWriter.WriteHeader(201)
	json.NewEncoder(responseWriter).Encode(functionTask)

}

func (a *Api) GetTasksListHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseWriter.WriteHeader(200)
	json.NewEncoder(responseWriter).Encode(a.Worker.TasksDb) // todo: format
}

func (a *Api) GetFunctionsListHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseWriter.WriteHeader(200)

	var functionsMap = helpers.GetFunctionsList(a.Worker.TasksDb)

	json.NewEncoder(responseWriter).Encode(functionsMap)
}

func (a *Api) GetDefaultRevisionsListHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseWriter.WriteHeader(200)

	var revisionsMap = helpers.GetDefaultRevisionsList(a.Worker.TasksDb)

	json.NewEncoder(responseWriter).Encode(revisionsMap)
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
		log.Println("😡 no task with id:", idOfTask, "found")
		responseWriter.WriteHeader(404)
	}

	taskCopy := helpers.StopTask(a.Worker.TasksDb, idOfTask)
	a.Worker.AddTask(&taskCopy)
	responseWriter.WriteHeader(204)

}

func (a *Api) GetTaskInfoHandler(responseWriter http.ResponseWriter, request *http.Request) {

	taskId := chi.URLParam(request, "taskID")
	if taskId == "" {
		log.Println("😡 no taskId passed in request")
		responseWriter.WriteHeader(400)
	}

	idOfTask, _ := uuid.Parse(taskId)
	_, ok := a.Worker.TasksDb[idOfTask]
	if !ok {
		log.Println("😡 no task with id:", idOfTask, "found")
		responseWriter.WriteHeader(404)
	}

	taskWithRevision := helpers.GetTaskInfo(a.Worker.TasksDb, idOfTask)

	responseWriter.WriteHeader(200)
	json.NewEncoder(responseWriter).Encode(taskWithRevision)

}

func (a *Api) SwitchFunctionRevisionHandler(responseWriter http.ResponseWriter, request *http.Request) {

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

	helpers.SwitchFunctionRevision(
		a.Worker.TasksDb,
		functionName,
		functionRevision,
		switchRevisionValue)

	responseWriter.WriteHeader(200)

}

func (a *Api) InitRouter() {
	a.Settings = settings.GetSettings()
	a.Router = chi.NewRouter()
	a.Router.Route("/tasks", func(r chi.Router) {
		r.Post("/", a.AddTaskHandler)
		r.Get("/", a.GetTasksListHandler)

		r.Route("/{taskID}", func(r chi.Router) {
			r.Delete("/", a.StopTaskHandler)
			r.Get("/", a.GetTaskInfoHandler)
		})
	})
	a.Router.Route("/functions", func(r chi.Router) {
		r.Get("/", a.GetFunctionsListHandler)
	})

	a.Router.Route("/revisions", func(r chi.Router) {
		r.Get("/default", a.GetDefaultRevisionsListHandler)
		r.Put("/{functionName}/{functionRevision}/default/{switch}", a.SwitchFunctionRevisionHandler) // change the status of the revision
	})
}

func (a *Api) Start() { // Address???
	a.InitRouter()
	log.Println("🌍 Listening on " + a.Port)
	log.Fatal(http.ListenAndServe(":"+a.Port, a.Router))
}
