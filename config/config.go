// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package config define the loading of configuration of the program
package config

import (
	"encoding/json"
	"fmt"
	"math"
	"os"

	"github.com/LucasNoga/corpos-christie/lib/colors"
	"github.com/LucasNoga/corpos-christie/lib/utils"
)

// Config represents the configuration of the program with the tax metrics
type Config struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Tax     Tax
	TaxList []Tax `json:"tax"`
}

// Tax represent the metrics of french tax in a specific year
// This metrics are called 'tranche'
type Tax struct {
	Year     int       `json:"year"`     // Year of the tax specifications
	Tranches []Tranche `json:"tranches"` // List of Tranches
}

// Tranche is a unit to define several metrics to calculate tax
type Tranche struct {
	Min  int    `json:"min"`  // Minimun in euros to get in the tranche
	Max  int    `json:"max"`  // Maximum in euros to get in the tranche
	Rate string `json:"rate"` // Rate taxable in euros in this tranche
}

// LoadConfiguration load in struct cfg Config the configuration present in the file config.json
// return true is the config is loaded false if not,
// also return an error if a fail happened while loading config
func (cfg *Config) LoadConfiguration(file string) (bool, error) {
	// default path
	if file == "" {
		file = "./config.json"
	}

	jsonFile, err := os.Open(file)
	if err != nil {
		return false, err
	}
	defer jsonFile.Close()

	// Load json file
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(cfg)
	if err != nil {
		return false, err
	}

	cfg.addMaxValue()

	// Define tax of the current year as reference
	cfg.loadTaxYear()

	return true, nil
}

// addMaxValue set an infinite value for the last tranche for each year of tax metrics present in cfg Config
func (cfg *Config) addMaxValue() {
	// for last tranch of each tax in the list
	for _, tax := range cfg.TaxList {
		tax.Tranches[len(tax.Tranches)-1].Max = math.MaxInt64
	}
}

// loadTaxYear set a default tax metrics among the year of tax metrics set in cfg Config
// If we have the metrics of current year we set this
// If not we set the last tax metrics present in the cfg Config
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

// GetTax returns the Tax metrics to calculate tax of user
func (cfg *Config) GetTax() Tax {
	return cfg.Tax
}

// ChangeTax get in Taxlist of cfg the metrics of the year wished
func (cfg *Config) ChangeTax(year int) {
	for _, tax := range cfg.TaxList {
		if tax.Year == year {
			cfg.Tax = tax
			return
		}
	}
	fmt.Printf(colors.Red("%d is not on the list\n"), year)
	fmt.Printf(colors.Red("Get default tax year: %d\n"), cfg.GetTax().Year)
}

// LoadDefaultConfiguration load a default configuration into struct cfg Config if the file 'config.json' is not found
func (cfg *Config) LoadDefaultConfiguration() {
	fmt.Println("Loading Default configuration...")
	cfg.Name = "Corpos-Christie"
	cfg.Version = "1.0.0"
	cfg.Tax = Tax{
		Year: 2021,
		Tranches: []Tranche{
			{Min: 0, Max: 10084, Rate: "0%"},
			{Min: 10085, Max: 25710, Rate: "11%"},
			{Min: 25711, Max: 73516, Rate: "30%"},
			{Min: 73517, Max: 158122, Rate: "41%"},
			{Min: 158123, Max: math.MaxInt64, Rate: "45%"},
		},
	}
	cfg.TaxList = append(cfg.TaxList, cfg.Tax)
}
