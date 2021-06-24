package core

import (
	"corpos-christie/colors"
	"corpos-christie/config"
	"corpos-christie/user"
	"reflect"
	"testing"
)

// For testing
// $ cd core
// $ go test -v

// Global variables
var CONFIG *config.Config

// Init global variables
func init() {
	CONFIG = new(config.Config)
	CONFIG.Tranches = []config.Tranche{
		{Min: 0, Max: 10084, Percentage: 0},
		{Min: 10085, Max: 25710, Percentage: 11},
		{Min: 25711, Max: 73516, Percentage: 30},
		{Min: 73517, Max: 158122, Percentage: 41},
		{Min: 158123, Max: 1000000, Percentage: 45}}
}

// Test a valid process with 32000 of income for single person
func TestValidProcess(t *testing.T) {
	user := user.User{
		Income: 32000,
	}

	result := Process(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{Income: 32000, Tax: 3605, Remainder: 28395}
	t.Logf("Expected:\t\t%+v", expected)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected that the Income %v should be equal to %v", colors.Red(expected.Income), colors.Red(result.Income))
		t.Errorf("Expected that the Tax %v should be equal to %v", colors.Red(expected.Tax), colors.Red(result.Tax))
		t.Errorf("Expected that the Remainder %v should be equal to %v", colors.Red(expected.Remainder), colors.Red(result.Remainder))
	}
}

// Test parts with a couple and 2 childrens
func TestProcessForCoupleWith2Children(t *testing.T) {
	user := user.User{
		Income:     55950,
		IsInCouple: true,
		Children:   2,
		Parts:      3,
	}

	result := Process(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{Income: 55950, Tax: 2826, Remainder: 53124}
	t.Logf("Expected:\t\t%+v", expected)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected that the Income %v should be equal to %v", colors.Red(expected.Income), colors.Red(result.Income))
		t.Errorf("Expected that the Tax %v should be equal to %v", colors.Red(expected.Tax), colors.Red(result.Tax))
		t.Errorf("Expected that the Remainder %v should be equal to %v", colors.Red(expected.Remainder), colors.Red(result.Remainder))
	}
}
