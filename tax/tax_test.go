// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package tax is the algorithm to calculate taxes
package tax

import (
	"math"
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
		Year: 2022,
		Tranches: []config.Tranche{
			{Min: 0, Max: 10225, Rate: "0%"},
			{Min: 10226, Max: 26070, Rate: "11%"},
			{Min: 26071, Max: 74545, Rate: "30%"},
			{Min: 74546, Max: 160336, Rate: "41%"},
			{Min: 160337, Max: math.MaxInt64, Rate: "45%"},
		},
	}
	CONFIG.TaxList = []config.Tax{
		{
			Year: 2022,
			Tranches: []config.Tranche{

				{Min: 0, Max: 10225, Rate: "0%"},
				{Min: 10226, Max: 26070, Rate: "11%"},
				{Min: 26071, Max: 74545, Rate: "30%"},
				{Min: 74546, Max: 160336, Rate: "41%"},
				{Min: 160337, Max: 1000000, Rate: "45%"},
			},
		},
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

// Calculate tax for a single person with 30000 of income
func TestCalculateTaxForSinglePerson(t *testing.T) {
	var user user.User = user.User{}
	user.Income = 30000

	result := CalculateTax(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{income: 30000, tax: 2922, remainder: 27078}
	t.Logf("Expected:\t\t%+v", expected)

	if result.income != expected.income || result.tax != expected.tax || result.remainder != expected.remainder {
		t.Errorf("Expected that the Income %s should be equal to %s", colors.Red(expected.income), colors.Red(result.income))
		t.Errorf("Expected that the Tax %s should be equal to %s", colors.Red(expected.tax), colors.Red(result.tax))
		t.Errorf("Expected that the Remainder %s should be equal to %s", colors.Red(expected.remainder), colors.Red(result.remainder))
	}
}

// Calculate tax for a couple with 2 children, testing shares with a couple and 2 childrens
func TestCalculateTaxForCoupleWith2Children(t *testing.T) {
	user := user.User{
		Income:     60000,
		IsInCouple: true,
		Children:   2,
	}

	result := CalculateTax(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{income: 60000, tax: 3225, remainder: 56775}
	t.Logf("Expected:\t\t%+v", expected)

	if result.income != expected.income || result.tax != expected.tax || result.remainder != expected.remainder {
		t.Errorf("Expected that the Income %s should be equal to %s", colors.Red(expected.income), colors.Red(result.income))
		t.Errorf("Expected that the Tax %s should be equal to %s", colors.Red(expected.tax), colors.Red(result.tax))
		t.Errorf("Expected that the Remainder %s should be equal to %s", colors.Red(expected.remainder), colors.Red(result.remainder))
	}
}

// Calculate tax for a couple with 3 children, testing shares with a couple and 3 childrens
func TestCalculateTaxForCoupleWith3Children(t *testing.T) {
	user := user.User{
		Income:     100000,
		IsInCouple: true,
		Children:   3,
	}

	result := CalculateTax(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{income: 100000, tax: 6501, remainder: 93499}
	t.Logf("Expected:\t\t%+v", expected)

	if result.income != expected.income || result.tax != expected.tax || result.remainder != expected.remainder {
		t.Errorf("Expected that the Income %s should be equal to %s", colors.Red(expected.income), colors.Red(result.income))
		t.Errorf("Expected that the Tax %s should be equal to %s", colors.Red(expected.tax), colors.Red(result.tax))
		t.Errorf("Expected that the Remainder %s should be equal to %s", colors.Red(expected.remainder), colors.Red(result.remainder))
	}
}

// Calculate tax for a couple with no children, testing shares with a couple and 0 childrens
func TestCalculateTaxForCoupleWithNoChildren(t *testing.T) {
	user := user.User{
		Income:     60000,
		IsInCouple: true,
		Children:   0,
	}

	result := CalculateTax(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{income: 60000, tax: 5843, remainder: 54157}
	t.Logf("Expected:\t\t%+v", expected)

	if result.income != expected.income || result.tax != expected.tax || result.remainder != expected.remainder {
		t.Errorf("Expected that the Income %s should be equal to %s", colors.Red(expected.income), colors.Red(result.income))
		t.Errorf("Expected that the Tax %s should be equal to %s", colors.Red(expected.tax), colors.Red(result.tax))
		t.Errorf("Expected that the Remainder %s should be equal to %s", colors.Red(expected.remainder), colors.Red(result.remainder))
	}
}

func TestCalculateTaxForIsolatedParent(t *testing.T) {
	user := user.User{
		Income:     30000,
		IsInCouple: false,
		Children:   2,
	}

	result := CalculateTax(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{income: 30000, tax: 488, remainder: 29512}
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

// Calculate reverse tax for a couple with 2 children, testing shares with a couple and 2 childrens
func TestCalculateReverseTaxForCoupleWith2Children(t *testing.T) {
	user := user.User{
		Remainder:  53124,
		IsInCouple: true,
		Children:   2,
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

// Calculate reverse tax for a couple with 3 children, testing shares with a couple and 3 childrens
func TestCalculateReverseTaxForCoupleWith3Children(t *testing.T) {
	user := user.User{
		Remainder:  93437,
		IsInCouple: true,
		Children:   3,
	}

	result := calculateReverseTax(&user, CONFIG)
	t.Logf("Function result:\t%+v", result)

	expected := Result{income: 100000, tax: 6563, remainder: 93437}
	t.Logf("Expected:\t\t%+v", expected)

	if result.income != expected.income && result.tax != expected.tax && result.remainder != expected.remainder {
		t.Errorf("Expected that the Income %s should be equal to %s", colors.Red(expected.income), colors.Red(result.income))
		t.Errorf("Expected that the Tax %s should be equal to %s", colors.Red(expected.tax), colors.Red(result.tax))
		t.Errorf("Expected that the Remainder %s should be equal to %s", colors.Red(expected.remainder), colors.Red(result.remainder))
	}
}

// Create user single with no children and check shares
func TestUserSingleOnlyIncome(t *testing.T) {
	var sharesRef float64 = 1.

	var user user.User = user.User{
		IsInCouple: false,
		Children:   0,
	}

	var shares float64 = getShares(user)
	t.Logf("User reference %+v", user)

	// Testing shares
	if sharesRef != shares {
		t.Errorf("Expected that the Shares \n%f\n should be equal to \n%f", sharesRef, shares)
	}
}

// Create user in couple with no children and check shares
func TestUserInCoupleNoChildren(t *testing.T) {
	var sharesRef float64 = 2.

	var user user.User = user.User{
		IsInCouple: true,
		Children:   0,
	}

	var shares float64 = getShares(user)
	t.Logf("User reference %+v", user)

	// Testing shares
	if sharesRef != shares {
		t.Errorf("Expected that the Shares \n%f\n should be equal to \n%f", sharesRef, shares)
	}
}

// Create user in couple with 3 children and check shares
func TestUserInCoupleWith3Children(t *testing.T) {
	var sharesRef float64 = 4

	var user user.User = user.User{
		IsInCouple: true,
		Children:   3,
	}

	var shares float64 = getShares(user)
	t.Logf("User reference %+v", user)

	// Testing shares
	if sharesRef != shares {
		t.Errorf("Expected that the Shares \n%f\n should be equal to \n%f", sharesRef, shares)
	}
}

// Create user single with 4 children and check shares
func TestUserInSingleWith4Children(t *testing.T) {
	var sharesExpected float64 = 4.5

	var user user.User = user.User{
		IsInCouple: false,
		Children:   4,
	}

	var shares float64 = getShares(user)
	t.Logf("User reference %+v", user)

	// Testing shares
	if sharesExpected != shares {
		t.Errorf("Expected that the Shares \n%f\n should be equal to \n%f", sharesExpected, shares)
	}
}

// Create user in couple with 4 children and check shares
func TestUserInCoupleWith4Children(t *testing.T) {
	var sharesRef float64 = 5

	var user user.User = user.User{
		IsInCouple: true,
		Children:   4,
	}

	var shares float64 = getShares(user)
	t.Logf("User reference %+v", user)

	// Testing shares
	if sharesRef != shares {
		t.Errorf("Expected that the Shares \n%f\n should be equal to \n%f", sharesRef, shares)
	}
}
