package worker

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bots-garden/procyon/procyon-launcher/settings"
	"github.com/bots-garden/procyon/procyon-launcher/task"
	"github.com/bots-garden/procyon/procyon-launcher/worker/helpers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ApiConfig struct {
	Address  string
	Port     string
	Worker   *Worker
	Settings settings.Settings
}

//var currentHttpPort = settings.GetSettings().Http.Start

func (a *ApiConfig) AddTaskHandler(c *gin.Context) {

	data := json.NewDecoder(c.Request.Body)
	data.DisallowUnknownFields()
	taskEvent := task.TaskEvent{}
	err := data.Decode(&taskEvent)

	if err != nil {
		c.Writer.WriteHeader(400)

		e := ErrResponse{
			HTTPStatusCode: 400,
			Message:        fmt.Sprintf("Error unmarshalling body: %v\n", err),
		}
		json.NewEncoder(c.Writer).Encode(e)
		return
	}

	functionTask := helpers.AddTask(taskEvent, a.Settings)

	a.Worker.AddTask(&functionTask)

	c.Writer.WriteHeader(201)
	json.NewEncoder(c.Writer).Encode(functionTask)
}

func (a *ApiConfig) GetTasksListHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Writer.WriteHeader(200)
	json.NewEncoder(c.Writer).Encode(a.Worker.TasksDb) // todo: format
}

func (a *ApiConfig) GetFunctionsListHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Writer.WriteHeader(200)

	var functionsMap = helpers.GetFunctionsList(a.Worker.TasksDb)

	json.NewEncoder(c.Writer).Encode(functionsMap)
}

func (a *ApiConfig) GetDefaultRevisionsListHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Writer.WriteHeader(200)

	var revisionsMap = helpers.GetDefaultRevisionsList(a.Worker.TasksDb)

	json.NewEncoder(c.Writer).Encode(revisionsMap)
}

func (a *ApiConfig) StopTaskHandler(c *gin.Context) {
	taskId := c.Param("taskID")
	if taskId == "" {
		log.Println("😡 no taskId passed in request")
		c.Writer.WriteHeader(400)
	}

	idOfTask, _ := uuid.Parse(taskId)
	_, ok := a.Worker.TasksDb[idOfTask]
	if !ok {
		log.Println("😡 no task with id:", idOfTask, "found")
		c.Writer.WriteHeader(404)
	}

	taskCopy := helpers.StopTask(a.Worker.TasksDb, idOfTask)
	a.Worker.AddTask(&taskCopy)
	c.Writer.WriteHeader(204)
}

func (a *ApiConfig) GetTaskInfoHandler(c *gin.Context) {

	taskId := c.Param("taskID")

	if taskId == "" {
		log.Println("😡 no taskId passed in request")
		c.Writer.WriteHeader(400)
	}

	idOfTask, _ := uuid.Parse(taskId)
	_, ok := a.Worker.TasksDb[idOfTask]
	if !ok {
		log.Println("😡 no task with id:", idOfTask, "found")
		c.Writer.WriteHeader(404)
	}

	taskWithRevision := helpers.GetTaskInfo(a.Worker.TasksDb, idOfTask)

	c.Writer.WriteHeader(200)
	json.NewEncoder(c.Writer).Encode(taskWithRevision)
}

func (a *ApiConfig) SwitchFunctionRevisionHandler(c *gin.Context) {
	functionName := c.Param("functionName")

	if functionName == "" {
		log.Println("😡 no functionName passed in request")
		c.Writer.WriteHeader(400)
	}
	functionRevision := c.Param("functionRevision")
	if functionRevision == "" {
		log.Println("😡 no functionRevision passed in request")
		c.Writer.WriteHeader(400)
	}
	switchRevisionValue := c.Param("switch")
	if switchRevisionValue == "" {
		log.Println("😡 no switchRevisionValue (on/off) passed in request")
		c.Writer.WriteHeader(400)
	}

	helpers.SwitchFunctionRevision(
		a.Worker.TasksDb,
		functionName,
		functionRevision,
		switchRevisionValue)

	c.Writer.WriteHeader(200)
}

func (a *ApiConfig) Start() {

	a.Settings = settings.GetSettings()

	r := gin.Default()

	r.POST("/tasks", a.AddTaskHandler)
	r.GET("/tasks", a.GetTasksListHandler)

	r.DELETE("/tasks/:taskID", a.StopTaskHandler)
	r.GET("/tasks/:taskID", a.GetTaskInfoHandler)

	r.GET("/functions", a.GetFunctionsListHandler)

	r.GET("/revisions/default", a.GetDefaultRevisionsListHandler)
	r.PUT("/revisions/:functionName/:functionRevision/default/:switch", a.SwitchFunctionRevisionHandler) // change the status of the revision

	log.Println("🌍 Listening on " + a.Port)
	r.Run(":" + a.Port)
}
