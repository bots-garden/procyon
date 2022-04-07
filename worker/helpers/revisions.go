package helpers

import (
	"log"

	"github.com/bots-garden/procyon/task"
	"github.com/bots-garden/procyon/worker/data"
	"github.com/google/uuid"
)

func GetDefaultRevisionsList(tasksDb map[uuid.UUID]*task.Task) map[string]data.FunctionRecord {
	// TODO: 🌺 use Redis
	var revisionsMap map[string]data.FunctionRecord
	revisionsMap = make(map[string]data.FunctionRecord)

	// parse map
	for key, element := range tasksDb {

		if element.Config.DefaultRevision == true {
			revisionsMap[element.Config.FunctionName] = data.FunctionRecord{
				WasmFunctionHttpPort: element.Config.WasmFunctionHttpPort,
				TaskId: key,
				DefaultRevision: element.Config.DefaultRevision,
			}
		}
		
	}
	return revisionsMap
}

// Switch on (or off) the revision of a function
// Then apply the reverse switch to the other revisions of the function
func SwitchFunctionRevision(tasksDb map[uuid.UUID]*task.Task, functionName string, functionRevision string, switchRevisionValue string) {
  // TODO: 🌺 use Redis
	// search the id of the related task
	for key, element := range tasksDb {

		if element.Config.FunctionName==functionName {
			if element.Config.FunctionRevision==functionRevision {
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
						"🔵 Revision switch:", switchRevisionValue,  key, "=>", 
						"TaskId:", element.Id, 
						"Function", element.Config.FunctionName, element.Config.FunctionRevision, element.Config.DefaultRevision)
			} else { // this is not the default revision
				if switchRevisionValue == "on" {
					element.Config.DefaultRevision = false
					} else {
						if switchRevisionValue == "off" {
							element.Config.DefaultRevision = true
						} else {
							// no change
							log.Println("😡 bad value for switchRevisionValue:", switchRevisionValue)
						}
					}
			}

		}
  }

}

