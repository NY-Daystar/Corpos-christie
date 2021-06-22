package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Load json file from string path (config.json)
func LoadConfiguration(cfg *Config) {
	const file string = "./config.json"
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(cfg)
}
