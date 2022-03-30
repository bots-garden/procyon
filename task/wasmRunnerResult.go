package task

import "github.com/google/uuid"

type WasmRunnerResult struct {
	Error    error
	Action   string
	RunnerId uuid.UUID
	Pid      int
	Result   string
}

//TODO: create a constant  or a type for the result ?
//TODO: or create a "result status"?
