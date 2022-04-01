package task

import (
	"time"

	"github.com/google/uuid"
)

type TaskEvent struct {
	Id                   uuid.UUID
	Executor             ExecutorType // 1: sat 2: galago (other kind of runner)
	WasmFileName         string
	WasmFunctionHttpPort int
	WasmRegistryUrl      string
	FunctionName         string
	FunctionRevision     string
	DefaultRevision      bool
	Timestamp            time.Time // record the time the event was requested
}
