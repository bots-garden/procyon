package helpers

import (
	"github.com/bots-garden/procyon/procyon-launcher/task"
	"github.com/bots-garden/procyon/procyon-launcher/worker/data"
	"github.com/google/uuid"
)

func GetFunctionsList(tasksDb map[uuid.UUID]*task.Task) map[string]data.FunctionRecord {
	// TODO: 🌺 use Redis
	var functionsMap map[string]data.FunctionRecord
	functionsMap = make(map[string]data.FunctionRecord)

	// parse map
	for key, element := range tasksDb {

		functionsMap[element.Config.FunctionName+"-"+element.Config.FunctionRevision] = data.FunctionRecord{
			WasmFunctionHttpPort: element.Config.WasmFunctionHttpPort,
			TaskId: key,
			DefaultRevision: element.Config.DefaultRevision,
		}

	}
	return functionsMap
}