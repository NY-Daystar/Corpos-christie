// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package config define the loading of configuration of the program
package config

import (
	"math"
	"reflect"
	"testing"
)

// For testing
// $ cd config
// $ go test -v

// Global variables
var CONFIG_REFERENCE Config

// Init global variables
func init() {
	CONFIG_REFERENCE.Tax = Tax{
		Year: 2022,
		Tranches: []Tranche{
			{Min: 0, Max: 10225, Rate: "0%"},
			{Min: 10226, Max: 26070, Rate: "11%"},
			{Min: 26071, Max: 74545, Rate: "30%"},
			{Min: 74546, Max: 160336, Rate: "41%"},
			{Min: 160337, Max: math.MaxInt64, Rate: "45%"},
		},
	}
}

// Test if tranche are well setup
func TestValidConfig(t *testing.T) {
	t.Logf("Reference config %+v", CONFIG_REFERENCE)

	var cfg Config
	cfg.LoadConfiguration("../config.json")
	t.Logf("Config loaded %+v", cfg)

	if !reflect.DeepEqual(CONFIG_REFERENCE.Tax.Tranches, cfg.Tax.Tranches) {
		t.Errorf("Expected that the configRef \n%v\n should be equal to \n%v", CONFIG_REFERENCE.Tax.Tranches, cfg.Tax)
	}
}

// Test loading of the default configuration
func TestLoadConfigWithNoFileSoLoadDefaultConfig(t *testing.T) {
	t.Logf("Reference config %+v", CONFIG_REFERENCE)

	var cfg *Config = new(Config)
	_, err := cfg.LoadConfiguration("config_file_not_exist.json") // load a file which doesn't exist
	if err != nil {
		cfg.LoadDefaultConfiguration()
	}
	t.Logf("Config loaded %+v", cfg)

	if !reflect.DeepEqual(CONFIG_REFERENCE.Tax.Tranches, cfg.Tax.Tranches) {
		t.Errorf("Expected that the configRef \n%v\n should be equal to \n%v", CONFIG_REFERENCE.Tax.Tranches, cfg.Tax)
	}
}

// Test to compare a json data structure to the golang config structure
func TestConfigLoadedFitWithInterface(t *testing.T) {
	configJson := make(map[string]interface{})
	configJson["taxlist"] = []Tax{
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
	}

	var cfg *Config = new(Config)
	_, err := cfg.LoadConfiguration("config_file_not_exist.json") // load a file which doesn't exist
	if err != nil {
		cfg.LoadDefaultConfiguration()
	}

	if !reflect.DeepEqual(configJson["taxlist"], cfg.TaxList) {
		t.Errorf("Expected that the configJson \n%v\n should be equal to \n%v", configJson["taxlist"], cfg.TaxList)
	}
}
