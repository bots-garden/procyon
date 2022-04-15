package task

import (
	"log"
	"os"
	"os/exec"

	"github.com/google/uuid"
)

type WasmRunner struct {
	RunnerConfig   *Config
	RunnerId uuid.UUID //!!! useful or not?
	Process  *os.Process
}

// Start() is Run()
func (wasmRunner *WasmRunner) Start() *WasmRunnerResult {

	log.Println("📝 wasm runner config:", wasmRunner.RunnerConfig)
	log.Println("🚀 starting the wasm runner...")

	cmd := &exec.Cmd{
		Path:   wasmRunner.RunnerConfig.ExecutorPath,
		Args:   wasmRunner.RunnerConfig.Args,
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	newEnv := append(os.Environ(), wasmRunner.RunnerConfig.Env...)
	cmd.Env = newEnv

	err := cmd.Start()

	if err != nil {
		log.Println("💥 [exec]", err.Error())
		return &WasmRunnerResult{
			Action:   "start",
			RunnerId: wasmRunner.RunnerId,
			Result:   "KO",
			Error:    err,
		}
	}

	wasmRunner.Process = cmd.Process

	return &WasmRunnerResult{
		Action:   "start",
		RunnerId: wasmRunner.RunnerId,
		Pid:      cmd.Process.Pid,
		Result:   "OK",
		Error:    nil,
	}
}

//!!! test if sat or galao (I think it's somewhere else)
func (wasmRunner *WasmRunner) Remove() *WasmRunnerResult {
	err := os.Remove(wasmRunner.RunnerConfig.WasmFilePath)
	if err != nil {
		log.Println("💥 [remove]", err.Error())
		return &WasmRunnerResult{
			Action:   "remove",
			RunnerId: wasmRunner.RunnerId,
			Result:   "KO",
			Error:    err,
		}
	}

	return &WasmRunnerResult{
		Action:   "remove",
		RunnerId: wasmRunner.RunnerId,
		Result:   "OK",
		Error:    nil,
	}
}

func (wasmRunner *WasmRunner) Stop() *WasmRunnerResult {

	log.Println("📝 wasm runner config:", wasmRunner.RunnerConfig)

	log.Println("🛑 stoping the wasm runner...", wasmRunner.RunnerId)

	err := wasmRunner.Process.Kill()

	if err != nil {
		log.Println("💥 [stop]", err.Error())
		return &WasmRunnerResult{
			Action:   "stop",
			RunnerId: wasmRunner.RunnerId,
			Result:   "KO",
			Error:    err,
		}
	}

	return &WasmRunnerResult{
		Action:   "stop",
		RunnerId: wasmRunner.RunnerId,
		Result:   "OK",
		Error:    nil,
	}
}
