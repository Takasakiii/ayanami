package config

import (
	"encoding/json"
	"os"
)

func writeDefaultConfig(config *Config) {
	file, err := os.Create(configFile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	if err != nil {
		panic(err)
	}

	bytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		panic(err)
	}

	_, err = file.Write(bytes)
	if err != nil {
		panic(err)
	}
}
