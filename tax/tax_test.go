// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package tax is the algorithm to calculate taxes
package tax

import (
	"testing"

	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/model"
)

// For testing
// $ cd tax
// $ go test -v

// Global variables
var cfg *config.Config

// Init global variables
func init() {
	cfg = config.New()
	cfg.ChangeTax(2022)
}

// Calculate tax
func TestCalculateTax(t *testing.T) {

	tests := []struct {
		user     model.User
		expected Result
	}{
		{
			user:     model.User{Income: 30000},
			expected: Result{Income: 30000, Tax: 2922, Remainder: 27078},
		},
		{
			user:     model.User{Income: 60000, IsInCouple: true, Children: 2},
			expected: Result{Income: 60000, Tax: 3225, Remainder: 56775},
		},
		{
			user:     model.User{Income: 100000, IsInCouple: true, Children: 3},
			expected: Result{Income: 100000, Tax: 6501, Remainder: 93499},
		},
		{
			user:     model.User{Income: 60000, IsInCouple: true, Children: 0},
			expected: Result{Income: 60000, Tax: 5843, Remainder: 54157},
		},
		{
			user:     model.User{Income: 30000, IsInCouple: false, Children: 2},
			expected: Result{Income: 30000, Tax: 488, Remainder: 29512},
		},
	}

	for _, testCase := range tests {
		result := CalculateTax(&testCase.user, cfg)
		t.Logf("Function result:\t%+v", result)

		if result.Income != testCase.expected.Income || result.Tax != testCase.expected.Tax || result.Remainder != testCase.expected.Remainder {
			t.Errorf("Expected that the Income %d should be equal to %d", result.Income, testCase.expected.Income)
			t.Errorf("Expected that the Tax %f should be equal to %f", result.Tax, testCase.expected.Tax)
			t.Errorf("Expected that the Remainder %f should be equal to %f", result.Remainder, testCase.expected.Remainder)
		}
	}
}

// Calculate reverse tax
func TestCalculateReverseTax(t *testing.T) {

	tests := []struct {
		user     model.User
		expected Result
	}{
		{
			user:     model.User{Remainder: 28395},
			expected: Result{Income: 31880, Tax: 3485, Remainder: 28395},
		},
		{
			user:     model.User{Remainder: 53124, IsInCouple: true, Children: 2},
			expected: Result{Income: 55896, Tax: 2772, Remainder: 53124},
		},
		{
			user:     model.User{Remainder: 93437, IsInCouple: true, Children: 3},
			expected: Result{Income: 99929, Tax: 6492, Remainder: 93437},
		},
	}

	for _, testCase := range tests {
		result := CalculateReverseTax(&testCase.user, cfg)
		t.Logf("Function result:\t%+v", result)

		if result.Income != testCase.expected.Income || result.Tax != testCase.expected.Tax || result.Remainder != testCase.expected.Remainder {
			t.Errorf("Expected that the Income %d should be equal to %d", result.Income, testCase.expected.Income)
			t.Errorf("Expected that the Tax %f should be equal to %f", result.Tax, testCase.expected.Tax)
			t.Errorf("Expected that the Remainder %f should be equal to %f", result.Remainder, testCase.expected.Remainder)
		}
	}
}

// Create user single with no children and check shares
func TestUserSingleOnlyIncome(t *testing.T) {
	var sharesRef = 1.

	var user = model.User{
		IsInCouple: false,
		Children:   0,
	}

	var shares = getShares(user)
	t.Logf("User reference %+v", user)

	// Testing shares
	if sharesRef != shares {
		t.Errorf("Expected that the Shares \n%f\n should be equal to \n%f", sharesRef, shares)
	}
}

// Create user in couple with no children and check shares
func TestUserInCoupleNoChildren(t *testing.T) {
	var sharesRef = 2.

	var user = model.User{
		IsInCouple: true,
		Children:   0,
	}

	var shares = getShares(user)
	t.Logf("User reference %+v", user)

	// Testing shares
	if sharesRef != shares {
		t.Errorf("Expected that the Shares \n%f\n should be equal to \n%f", sharesRef, shares)
	}
}

// Create user in couple with 3 children and check shares
func TestUserInCoupleWith3Children(t *testing.T) {
	var sharesRef = 4.

	var user = model.User{
		IsInCouple: true,
		Children:   3,
	}

	var shares = getShares(user)
	t.Logf("User reference %+v", user)

	// Testing shares
	if sharesRef != shares {
		t.Errorf("Expected that the Shares \n%f\n should be equal to \n%f", sharesRef, shares)
	}
}

// Create user single with 4 children and check shares
func TestUserInSingleWith4Children(t *testing.T) {
	var sharesExpected = 4.5

	var user = model.User{
		IsInCouple: false,
		Children:   4,
	}

	var shares = getShares(user)
	t.Logf("User reference %+v", user)

	// Testing shares
	if sharesExpected != shares {
		t.Errorf("Expected that the Shares \n%f\n should be equal to \n%f", sharesExpected, shares)
	}
}

// Create user in couple with 4 children and check shares
func TestUserInCoupleWith4Children(t *testing.T) {
	var sharesRef = 5.

	var user = model.User{
		IsInCouple: true,
		Children:   4,
	}

	var shares = getShares(user)
	t.Logf("User reference %+v", user)

	// Testing shares
	if sharesRef != shares {
		t.Errorf("Expected that the Shares \n%f\n should be equal to \n%f", sharesRef, shares)
	}
}
