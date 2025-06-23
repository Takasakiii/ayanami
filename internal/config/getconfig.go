package config

import "os"

func GetConfig() Config {
	if _, err := os.Stat(configFile); err == nil {
		return loadConfig()
	} else if os.IsNotExist(err) {
		config := newConfig()
		writeDefaultConfig(&config)
		return config
	} else {
		panic(err)
	}
}
