package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//TODO mettre en mode objet
// Load json file from string path (config.json)
// func LoadConfiguration(cfg *Config, file string) {
// 	// default path
// 	if file == "" {
// 		file = "./config.json"
// 	}
// 	// const file string = "./config.json"
// 	configFile, err := os.Open(file)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		log.Printf("Loading Default configuration...")
// 		LoadDefaultConfiguration(cfg)
// 	}
// 	defer configFile.Close()
// 	jsonParser := json.NewDecoder(configFile)
// 	jsonParser.Decode(cfg)
// }

func (cfg *Config) LoadConfiguration(file string) {
	// default path
	if file == "" {
		file = "./config.json"
	}
	// const file string = "./config.json"
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		log.Printf("Loading Default configuration...")
		cfg.LoadDefaultConfiguration()
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(cfg)
}

//TODO mettre en mode objet
// Load default configuration file if we don't have a 'config.json file'
func (cfg *Config) LoadDefaultConfiguration() {
	cfg.Name = "Corpos-Christie"
	cfg.Version = "1.0.0"
	cfg.Tranches = []Tranche{
		{Min: 0, Max: 10084, Percentage: 0},
		{Min: 10085, Max: 25710, Percentage: 11},
		{Min: 25711, Max: 73516, Percentage: 30},
		{Min: 73517, Max: 158122, Percentage: 41},
		{Min: 158123, Max: 1000000, Percentage: 45}}
}
