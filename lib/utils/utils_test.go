package utils

import (
	"testing"
)

// For testing
// $ cd utils
// $ go test -v

// Test string conversion to int
func TestStringConvertToInt(t *testing.T) {
	var vString string = "32000"

	_, err := ConvertStringToInt(vString)
	// Testing parts
	if err != nil {
		t.Errorf("Impossible to convert this string %v, err: %v", vString, err)
	}
}

// Test currentYear
func TestUserInCoupleNoChildren(t *testing.T) {
	var currentYear int = 2021

	var year = GetCurrentYear()
	if currentYear != year {
		t.Errorf("The currentYear '%v' is not return by the function GetCurrentYear() '%v'", currentYear, year)
	}
}
