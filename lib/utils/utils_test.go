// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package utils define functions to multiple uses
package utils

import (
	"testing"
)

// For testing
// $ cd utils
// $ go test -v

// Test string conversion to int
func TestStringConvertToInt(t *testing.T) {
	var stringRef string = "32000"
	var expected int = 32000

	val, err := ConvertStringToInt(stringRef)
	t.Logf("Value converted %d", val)

	if err != nil {
		t.Errorf("Impossible to convert this string %s, err: %v", stringRef, err)
	} else if val != expected {
		t.Errorf("Value '%d' is not the same as ref '%s'", val, stringRef)
	}
}

// Test string percentage conversion to float64
func TestConvertPercentageToFloat64(t *testing.T) {
	var stringRef string = "15%"
	var expected float64 = 15

	val, err := ConvertPercentageToFloat64(stringRef)
	t.Logf("Value converted %f", val)

	if err != nil {
		t.Errorf("Impossible to convert this string %s, err: %v", stringRef, err)
	} else if val != expected {
		t.Errorf("Value '%f' is not the same as ref '%s'", val, stringRef)
	}
}

// Test currentYear
func TestUserInCoupleNoChildren(t *testing.T) {
	var currentYear int = 2021

	var year = GetCurrentYear()
	if currentYear != year {
		t.Errorf("The currentYear '%d' is not return by the function GetCurrentYear() '%d'", currentYear, year)
	}
}

// Test if function return the maxLength among this item
func TestMaxLength(t *testing.T) {
	var longItem string = "test max length"
	var items = []string{"tax", "options", longItem, "db clean"} // 'long item' has to be the long string in array
	var refMaxLength int = len(longItem)
	t.Logf("Length of %s : %d", longItem, refMaxLength)

	var maxLength int = GetMaxLength(items)
	t.Logf("Length find by function: %d", maxLength)

	if refMaxLength != maxLength {
		t.Errorf("The refMaxLength '%d' is not the same has maxLength '%d'", refMaxLength, maxLength)
	}
}

// Test if function return the right padding
func TestGetPadding(t *testing.T) {
	var longItem string = "test get padding function"
	var items = []string{"tax", "options", longItem, "db clean"} // 'long item' has to be the long string in array
	var refGetPadding int = len(longItem)
	t.Logf("Length of %s : %d", longItem, refGetPadding)

	var padding int = getPadding(items)
	t.Logf("Length find by function: %d", padding)

	if refGetPadding != padding {
		t.Errorf("The refMaxLength '%d' is not the same has maxLength '%d'", refGetPadding, padding)
	}
}

// Test if the padding is set
func TestSetPadding(t *testing.T) {
	var items = []string{"tax", "options", "db"}

	// get padding for 'tax', 'options', 'db' values
	var paddings []int
	for _, val := range items {
		paddings = append(paddings, len(SetPadding(items, val))+len(val))
	}

	// check if all value are the same
	var samePaddings bool = func(a []int) bool {
		for i := 1; i < len(a); i++ {
			if a[i] != a[0] {
				return false
			}
		}
		return true
	}(paddings)

	// Padding has to be equal at the end between every value if we retranch their length to uniformize it
	if !samePaddings {
		t.Errorf("The paddings of paddings in the array %+v are not equal, paddings %+v", items, paddings)
	}
}
