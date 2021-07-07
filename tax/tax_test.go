// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package tax is the algorithm to calculate taxes
package tax

import (
	"testing"

	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/lib/colors"
	"github.com/LucasNoga/corpos-christie/user"
)

// For testing
// $ cd tax
// $ go test -v

// Global variables
var CONFIG *config.Config

// Init global variables
func init() {
	CONFIG = new(config.Config)
	CONFIG.Tax = config.Tax{
		Year: 2021,
		Tranches: []config.Tranche{
			{Min: 0, Max: 10084, Rate: "0%"},
			{Min: 10085, Max: 25710, Rate: "11%"},
			{Min: 25711, Max: 73516, Rate: "30%"},
			{Min: 73517, Max: 158122, Rate: "41%"},
			{Min: 158123, Max: 1000000, Rate: "45%"},
		},
	}
	CONFIG.TaxList = []config.Tax{
		{
			Year: 2021,
			Tranches: []config.Tranche{
				{Min: 0, Max: 10084, Rate: "0%"},
				{Min: 10085, Max: 25710, Rate: "11%"},
				{Min: 25711, Max: 73516, Rate: "30%"},
				{Min: 73517, Max: 158122, Rate: "41%"},
				{Min: 158123, Max: 1000000, Rate: "45%"},
			},
		},
		{
			Year: 2020,
			Tranches: []config.Tranche{
				{Min: 0, Max: 10064, Rate: "0%"},
				{Min: 10065, Max: 25659, Rate: "11%"},
				{Min: 25660, Max: 73369, Rate: "30%"},
				{Min: 73370, Max: 157806, Rate: "41%"},
				{Min: 157807, Max: 1000000, Rate: "45%"},
			},
		},
		{
			Year: 2019,
			Tranches: []config.Tranche{
				{Min: 0, Max: 10064, Rate: "0%"},
				{Min: 10065, Max: 27794, Rate: "14%"},
				{Min: 27795, Max: 74517, Rate: "30%"},
				{Min: 74518, Max: 157806, Rate: "41%"},
				{Min: 157807, Max: 1000000, Rate: "45%"},
			},
		},
	}
}

// Calculate tax for a single person with 32000 of income
func TestCalculateTaxForSinglePerson(t *testing.T) {
	user := user.User{
		Income: 32000,
	}

	result := calculateTax(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{income: 32000, tax: 3605, remainder: 28395}
	t.Logf("Expected:\t\t%+v", expected)

	if result.income != expected.income || result.tax != expected.tax || result.remainder != expected.remainder {
		t.Errorf("Expected that the Income %s should be equal to %s", colors.Red(expected.income), colors.Red(result.income))
		t.Errorf("Expected that the Tax %s should be equal to %s", colors.Red(expected.tax), colors.Red(result.tax))
		t.Errorf("Expected that the Remainder %s should be equal to %s", colors.Red(expected.remainder), colors.Red(result.remainder))
	}
}

// Set another year to calculate tax for a single person with 32000 of income year 2019
func TestCalculateTaxForSinglePersonIn2019(t *testing.T) {
	user := user.User{
		Income: 32000,
	}
	year := 2019

	// Set tax metrics to 2019
	CONFIG.ChangeTax(year)
	t.Logf("New tax set:\t%+v", CONFIG.GetTax())

	result := calculateTax(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{income: 32000, tax: 3744, remainder: 28256}
	t.Logf("Expected:\t\t%+v", expected)

	if result.income != expected.income || result.tax != expected.tax || result.remainder != expected.remainder {
		t.Errorf("Expected that the Income %s should be equal to %s", colors.Red(expected.income), colors.Red(result.income))
		t.Errorf("Expected that the Tax %s should be equal to %s", colors.Red(expected.tax), colors.Red(result.tax))
		t.Errorf("Expected that the Remainder %s should be equal to %s", colors.Red(expected.remainder), colors.Red(result.remainder))
	}

	// Reset tax metrics to 2021
	CONFIG.ChangeTax(2021)
}

// Calculate tax for a couple with 2 children, testing parts with a couple and 2 childrens
func TestCalculateTaxForCoupleWith2Children(t *testing.T) {
	user := user.User{
		Income:     55950,
		IsInCouple: true,
		Children:   2,
	}

	result := calculateTax(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{income: 55950, tax: 2826, remainder: 53124}
	t.Logf("Expected:\t\t%+v", expected)

	if result.income != expected.income || result.tax != expected.tax || result.remainder != expected.remainder {
		t.Errorf("Expected that the Income %s should be equal to %s", colors.Red(expected.income), colors.Red(result.income))
		t.Errorf("Expected that the Tax %s should be equal to %s", colors.Red(expected.tax), colors.Red(result.tax))
		t.Errorf("Expected that the Remainder %s should be equal to %s", colors.Red(expected.remainder), colors.Red(result.remainder))
	}
}

// Calculate reverse tax for a single person to get at the end 28395
func TestCalculateReverseTaxForSinglePerson(t *testing.T) {
	user := user.User{
		Remainder: 28395,
	}

	result := calculateReverseTax(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{income: 32000, tax: 3605, remainder: 28395}
	t.Logf("Expected:\t\t%+v", expected)

	if result.income != expected.income && result.tax != expected.tax && result.remainder != expected.remainder {
		t.Errorf("Expected that the Income %s should be equal to %s", colors.Red(expected.income), colors.Red(result.income))
		t.Errorf("Expected that the Tax %s should be equal to %s", colors.Red(expected.tax), colors.Red(result.tax))
		t.Errorf("Expected that the Remainder %s should be equal to %s", colors.Red(expected.remainder), colors.Red(result.remainder))
	}
}

// Calculate reverse tax for a couple with 2 children, testing parts with a couple and 2 childrens
func TestCalculateReverseTaxForCoupleWith2Children(t *testing.T) {
	user := user.User{
		Remainder:  53124,
		IsInCouple: true,
		Children:   2,
		Parts:      3,
	}

	result := calculateReverseTax(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{income: 55950, tax: 2826, remainder: 53124}
	t.Logf("Expected:\t\t%+v", expected)

	if result.income != expected.income && result.tax != expected.tax && result.remainder != expected.remainder {
		t.Errorf("Expected that the Income %s should be equal to %s", colors.Red(expected.income), colors.Red(result.income))
		t.Errorf("Expected that the Tax %s should be equal to %s", colors.Red(expected.tax), colors.Red(result.tax))
		t.Errorf("Expected that the Remainder %s should be equal to %s", colors.Red(expected.remainder), colors.Red(result.remainder))
	}
}

// Create user single with no children and check parts
func TestUserSingleOnlyIncome(t *testing.T) {
	var partsRef float64 = 1.

	var user user.User = user.User{
		IsInCouple: false,
		Children:   0,
	}

	var parts float64 = getParts(user)
	t.Logf("User reference %+v", user)

	// Testing parts
	if partsRef != parts {
		t.Errorf("Expected that the Parts \n%f\n should be equal to \n%v", partsRef, parts)
	}
}

// Create user in couple with no children and check parts
func TestUserInCoupleNoChildren(t *testing.T) {
	var partsRef float64 = 2.

	var user user.User = user.User{
		IsInCouple: true,
		Children:   0,
	}

	var parts float64 = getParts(user)
	t.Logf("User reference %+v", user)

	// Testing parts
	if partsRef != parts {
		t.Errorf("Expected that the Parts \n%f\n should be equal to \n%v", partsRef, parts)
	}
}

// Create user in couple with 3 children and check parts
func TestUserInCoupleWith3Children(t *testing.T) {
	var partsRef float64 = 3.5

	var user user.User = user.User{
		IsInCouple: true,
		Children:   3,
	}

	var parts float64 = getParts(user)
	t.Logf("User reference %+v", user)

	// Testing parts
	if partsRef != parts {
		t.Errorf("Expected that the Parts \n%f\n should be equal to \n%v", partsRef, parts)
	}
}
