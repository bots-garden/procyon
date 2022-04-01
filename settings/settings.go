package settings

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// see procyon.json
type Settings struct {
	Executors struct {
		SatExecutorPath    string `json:"satExecutorPath"`
		GalagoExecutorPath string `json:"galagoExecutorPath"` // for another executor than Sat
	} `json:"executors"`
	Functions struct {
		WasmFilesDirectory string `json:"wasmFilesDirectory"` // not used by Sat
	} `json:"functions"`
	Http struct {
		Start int `json:"start"` // Procyon use this number for every wasm service (it is incremented at every launch)
	}
}

// Read settings of Procyon (in `./procyon.json`)
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
