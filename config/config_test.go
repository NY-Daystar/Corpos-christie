package config

import (
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
		Year: 2021,
		Tranches: []Tranche{
			{Min: 0, Max: 10084, Rate: 0},
			{Min: 10085, Max: 25710, Rate: 11},
			{Min: 25711, Max: 73516, Rate: 30},
			{Min: 73517, Max: 158122, Rate: 41},
			{Min: 158123, Max: 1000000, Rate: 45},
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
			Year: 2021,
			Tranches: []Tranche{
				{Min: 0, Max: 10084, Rate: 0},
				{Min: 10085, Max: 25710, Rate: 11},
				{Min: 25711, Max: 73516, Rate: 30},
				{Min: 73517, Max: 158122, Rate: 41},
				{Min: 158123, Max: 1000000, Rate: 45},
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
