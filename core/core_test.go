package core

import (
	"corpos-christie/config"
	"log"
	"testing"
)

// To test:
// $ cd core
// $ go test -v

// Test a valid process
func TestValidProcess(t *testing.T) {
	//TODO Voir si on a moyen de mettre cette config en variable global dans le script
	var cfg config.Config
	cfg.Tranches = []config.Tranche{
		{Min: 0, Max: 10084, Percentage: 0},
		{Min: 10085, Max: 25710, Percentage: 11},
		{Min: 25711, Max: 73516, Percentage: 30}}

	var income int = 32000
	Process(income, &cfg)

	//TODO mettre un log panic ou fatal f

	expected := Result{Income: 32000, Tax: 3605, Remainder: 28395}
	log.Printf("%v", expected)
}
