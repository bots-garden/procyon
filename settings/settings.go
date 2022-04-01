package settings

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Settings struct {
	Executors struct {
		GalagoExecutorPath string `json:"galagoExecutorPath"`
		SatExecutorPath    string `json:"satExecutorPath"`
	} `json:"executors"`
	Functions struct {
		WasmFilesDirectory string `json:"wasmFilesDirectory"`
	} `json:"functions"`
	Http struct {
		Start int `json:"start"`
	}
}
/*
{
  "executors": {
    "galagoExecutorPath": "./galago-wasm-runner/galago-runner",
    "satExecutorPath": "./runners/sat"
  },
  "functions": {
    "wasmFilePath": "./functions"
  }
}
*/

func GetSettings() Settings {
	settingsFile, err := ioutil.ReadFile("./procyon.json")
	if err != nil {
		log.Fatal(err)
	}
	settings := Settings{}

	err = json.Unmarshal([]byte(settingsFile), &settings)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("📝", settings)
	return settings
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
