package config

import (
	"log"
	"testing"
)

// To test:
// $ go test config/config.go config/config_test.go -v
// $ go test config/config.go config/config_test.go

// Test a valid config
func TestValidConfig(t *testing.T) {
	var cfg Config
	cfg.Tranches = []Tranche{{Min: 0, Max: 10084, Percentage: 0}, {Min: 10085, Max: 25710, Percentage: 11}, {Min: 25711, Max: 73516, Percentage: 30}}
	log.Printf("%v", cfg)
}
