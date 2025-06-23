package config

import (
	"encoding/json"
	"log"
	"os"
)

func loadConfig() Config {
	config := newConfig()
	configFile, err := os.Open(configFile)
	defer func(configFile *os.File) {
		_ = configFile.Close()
	}(configFile)

	if err != nil {
		log.Printf("[WARN] Failed to open config file: %v", err)
		return config
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		log.Printf("[WARN] Failed to parse config file: %v", err)
		return config
	}

	return config
}
