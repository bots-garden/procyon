package task

import (
	"time"

	"github.com/google/uuid"
)

type TaskEvent struct {
	Id                   uuid.UUID
	Executor             ExecutorType // 1: galago 2: sat
	WasmFileName         string
	WasmFunctionHttpPort int
	WasmRegistryUrl      string
	Timestamp            time.Time // record the time the event was requested
	//Task      Task
	//State     State
}

/*
```json
{
  "executor": 1,
  "wasmFileName": "hello.wasm",
  "wasmFunctionHttpPort": 8082,
  "wasmRegistryUrl": "https://localhost:9999/hello/hello.wasm"
}
```
*/
