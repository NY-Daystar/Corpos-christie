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
	CONFIG_REFERENCE.Tranches = []Tranche{
		{Min: 0, Max: 10084, Percentage: 0},
		{Min: 10085, Max: 25710, Percentage: 11},
		{Min: 25711, Max: 73516, Percentage: 30},
		{Min: 73517, Max: 158122, Percentage: 41},
		{Min: 158123, Max: 1000000, Percentage: 45}}
}

// Test if tranche are well setup
func TestValidConfig(t *testing.T) {
	t.Logf("Reference config %+v", CONFIG_REFERENCE)

	var cfg *Config = new(Config)
	cfg.LoadConfiguration("../config.json")
	t.Logf("Config loaded %+v", cfg)

	if !reflect.DeepEqual(CONFIG_REFERENCE.Tranches, cfg.Tranches) {
		t.Errorf("Expected that the configRef \n%v\n should be equal to \n%v", CONFIG_REFERENCE, cfg)
	}
}

// Test loading of the default configuration
func TestLoadConfigWithNoFileSoLoadDefaultConfig(t *testing.T) {
	t.Logf("Reference config %+v", CONFIG_REFERENCE)

	var cfg *Config = new(Config)
	cfg.LoadConfiguration("config_file_not_exist.json") // load a file which doesn't exist
	t.Logf("Config loaded %+v", cfg)

	if !reflect.DeepEqual(CONFIG_REFERENCE.Tranches, cfg.Tranches) {
		t.Errorf("Expected that the configRef \n%v\n should be equal to \n%v", CONFIG_REFERENCE, cfg)
	}
}
