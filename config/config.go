package config

import (
	"corpos-christie/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Load configuration from config.json file
func (cfg *Config) LoadConfiguration(file string) (bool, error) {
	// default path
	if file == "" {
		file = "./config.json"
	}

	// Load json file
	result, err := cfg.loadJson(file)
	// var result map[string]interface{}, err Error := cfg.loadJson(file)
	if err != nil {
		return false, err
	}

	// Loop through the Items; we're not interested in the key, just the values
	cfg.Name = result["name"].(string)
	cfg.Version = result["version"].(string)

	// JSON object parses into a map with string keys
	taxMap := result["tax"].(map[string]interface{})

	for key, value := range taxMap {
		// Init item tax
		var tax Tax
		tax.Year, _ = utils.ConvertStringToInt(key)

		// Initiate every tranche and add it to the list
		var tranches []interface{} = value.(map[string]interface{})["tranches"].([]interface{})
		for _, v := range tranches {
			t := v.(map[string]interface{})
			var tranche Tranche = Tranche{
				Min:        int(t["min"].(float64)),
				Max:        int(t["max"].(float64)),
				Percentage: t["percentage"].(float64),
			}
			tax.Tranches = append(tax.Tranches, tranche)
		}
		// Add tax from the year into tax list
		cfg.TaxList = append(cfg.TaxList, tax)
	}

	// Define tax of the current year as reference
	cfg.loadTaxYear()

	return true, nil
}

// Define the tax of the current year
func (cfg *Config) loadTaxYear() {
	for _, tax := range cfg.TaxList {
		if tax.Year == utils.GetCurrentYear() { // get tax of current year
			cfg.Tax = tax
			break
		}
	}

	// If no tax tranches are defined from current year load default tax 2021
	if len(cfg.Tax.Tranches) == 0 {
		cfg.Tax = cfg.TaxList[0]
	}
}

// Load a json file into an interface array
func (cfg *Config) loadJson(file string) (map[string]interface{}, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	return result, nil
}

// Load default configuration file if we don't have a 'config.json file'
func (cfg *Config) LoadDefaultConfiguration() {
	log.Printf("Loading Default configuration...")
	cfg.Name = "Corpos-Christie"
	cfg.Version = "1.0.0"
	cfg.Tax = Tax{
		Year: 2021,
		Tranches: []Tranche{
			{Min: 0, Max: 10084, Percentage: 0},
			{Min: 10085, Max: 25710, Percentage: 11},
			{Min: 25711, Max: 73516, Percentage: 30},
			{Min: 73517, Max: 158122, Percentage: 41},
			{Min: 158123, Max: 1000000, Percentage: 45},
		},
	}
	cfg.TaxList = append(cfg.TaxList, cfg.Tax)
}
