// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package config define the loading of configuration of the program
package config

import (
	"fmt"
	"math"

	"github.com/LucasNoga/corpos-christie/lib/colors"
	"github.com/LucasNoga/corpos-christie/lib/utils"
)

// Config represents the configuration of the program with the tax metrics
type Config struct {
	Name    string
	Version string
	Tax     Tax
	TaxList []Tax
}

// Tax represent the metrics of french tax in a specific year
// This metrics are called 'tranche'
type Tax struct {
	Year     int       // Year of the tax specifications
	Tranches []Tranche // List of Tranches
}

// Tranche is a unit to define several metrics to calculate tax
type Tranche struct {
	Min  int    // Minimun in euros to get in the tranche
	Max  int    // Maximum in euros to get in the tranche
	Rate string // Rate taxable in euros in this tranche
}

// New create new configuration
func New(name, version string) *Config {
	var config Config = Config{
		Name:    name,
		Version: version,
		TaxList: []Tax{
			{
				Year: 2022,
				Tranches: []Tranche{

					{Min: 0, Max: 10225, Rate: "0%"},
					{Min: 10226, Max: 26070, Rate: "11%"},
					{Min: 26071, Max: 74545, Rate: "30%"},
					{Min: 74546, Max: 160336, Rate: "41%"},
					{Min: 160337, Max: math.MaxInt64, Rate: "45%"},
				},
			},
			{
				Year: 2021,
				Tranches: []Tranche{
					{Min: 0, Max: 10084, Rate: "0%"},
					{Min: 10085, Max: 25710, Rate: "11%"},
					{Min: 25711, Max: 73516, Rate: "30%"},
					{Min: 73517, Max: 158122, Rate: "41%"},
					{Min: 158123, Max: math.MaxInt64, Rate: "45%"},
				},
			},
			{
				Year: 2020,
				Tranches: []Tranche{
					{Min: 0, Max: 10064, Rate: "0%"},
					{Min: 10065, Max: 25659, Rate: "11%"},
					{Min: 25660, Max: 73369, Rate: "30%"},
					{Min: 73370, Max: 157806, Rate: "41%"},
					{Min: 157807, Max: math.MaxInt64, Rate: "45%"},
				},
			},
			{
				Year: 2019,
				Tranches: []Tranche{
					{Min: 0, Max: 10064, Rate: "0%"},
					{Min: 10065, Max: 27794, Rate: "14%"},
					{Min: 27795, Max: 74517, Rate: "30%"},
					{Min: 74518, Max: 157806, Rate: "41%"},
					{Min: 157807, Max: math.MaxInt64, Rate: "45%"},
				},
			},
		},
	}

	// set tax list of current year
	config.loadTaxYear()

	return &config
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
		Year: 2022,
		Tranches: []Tranche{
			{Min: 0, Max: 10225, Rate: "0%"},
			{Min: 10226, Max: 26070, Rate: "11%"},
			{Min: 26071, Max: 74545, Rate: "30%"},
			{Min: 74546, Max: 160336, Rate: "41%"},
			{Min: 160337, Max: math.MaxInt64, Rate: "45%"},
		},
	}
	cfg.TaxList = append(cfg.TaxList, cfg.Tax)
}
