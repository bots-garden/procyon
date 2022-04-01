package task

import (
	"log"
	"strconv"

	"github.com/bots-garden/procyon/settings"
)

type ExecutorType int

// ??? or use interface
const (
	Unknown ExecutorType = iota
	Galago
	Sat
)

type Config struct {
	//Name                 string // "hello-function-config"
	Executor             ExecutorType
	ExecutorPath         string // "./galago-wasm-runner/galago-runner"
	WasmFileName         string // "hello.wasm"
	WasmFilePath         string // "./functions/hello.wasm"
	WasmFunctionHttpPort int
	WasmRegistryUrl      string // "https://localhost:9999/hello/hello.wasm"
	FunctionName         string
	FunctionRevision     string
	Env  []string // WASM_EXECUTOR_HTTP=9090
	Args []string // The first argument is the runer
}

func (config *Config) Initialize(settings settings.Settings) ExecutorType {

	log.Println("👁", settings)

	config.Args = append(config.Args, config.ExecutorPath)

	if config.Executor == Galago {
		
		config.WasmFilePath = settings.Functions.WasmFilesDirectory+"/"+config.WasmFileName

		config.ExecutorPath = settings.Executors.GalagoExecutorPath

		config.Args = append(config.Args, config.WasmFilePath, config.WasmRegistryUrl)

		config.Env = append(config.Env, "WASM_EXECUTOR_HTTP="+strconv.Itoa(config.WasmFunctionHttpPort))

		return Galago
	}
	
	if config.Executor == Sat {

		config.WasmFilePath = ""

		config.ExecutorPath = settings.Executors.SatExecutorPath

		config.Args = append(config.Args, config.WasmRegistryUrl)
		config.Env = append(config.Env, "SAT_HTTP_PORT="+strconv.Itoa(config.WasmFunctionHttpPort))
		return Sat
	}
	return Unknown
}
