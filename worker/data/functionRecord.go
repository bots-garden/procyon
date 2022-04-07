package data

import "github.com/google/uuid"

type FunctionRecord struct {
	WasmFunctionHttpPort int
	TaskId uuid.UUID
	DefaultRevision bool
}