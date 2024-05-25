// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package utils define functions to multiple uses
package utils

import (
	"testing"

	"fyne.io/fyne/v2/data/binding"
)

// For testing
// $ cd utils
// $ go test -v

// Test string conversion to int
func TestStringConvertToInt(t *testing.T) {
	var stringRef = "32000"
	var expected = 32000

	val, err := ConvertStringToInt(stringRef)
	t.Logf("Value converted %d", val)

	if err != nil {
		t.Errorf("Impossible to convert this string %d, err: %v", val, err)
	} else if val != expected {
		t.Errorf("Value '%d' is not the same as ref '%s'", val, stringRef)
	}
}

func TestBadStringConvertToInt(t *testing.T) {
	var stringRef = "abc"

	_, err := ConvertStringToInt(stringRef)
	t.Logf("Error generated %v", err)

	if err == nil {
		t.Errorf("Conversion has to generate error, err: %v", err)
	}
}

// Test bind string conversion to int
func TestBindStringConvertToInt(t *testing.T) {
	var ref = "15554"
	var stringRef binding.String = binding.BindString(&ref)
	var expected = 15554

	val := ConvertBindStringToInt(stringRef)
	t.Logf("Value converted %d", val)

	if val != expected {
		t.Errorf("Impossible to convert this string %d", val)
	}
}

func TestBadBindStringConvertToInt(t *testing.T) {
	var ref = "abc"
	var stringRef binding.String = binding.BindString(&ref)
	var expected = 0

	val := ConvertBindStringToInt(stringRef)
	t.Logf("Value converted %d", val)

	if val != expected {
		t.Errorf("Conversion has to return 0 and not: %d", val)
	}
}

// Test float to string
func TestConvertFloat64ToString(t *testing.T) {
	var floatRef = 5.54
	var expected = "5.54"

	val := ConvertFloat64ToString(floatRef)
	t.Logf("Value converted %s", val)

	// check on first 5 characters
	if val[:4] != expected[:4] {
		t.Errorf("Value '%s' is not the same as ref '%f'", val, floatRef)
	}
}

// Test int64 to string
func TestConvertInt64ToString(t *testing.T) {
	var intRef int64 = 7
	var expected = "7"

	val := ConvertInt64ToString(intRef)
	t.Logf("Value converted %s", val)

	if val != expected {
		t.Errorf("Value '%s' is not the same as ref '%d'", val, intRef)
	}
}

// Test int64 to string
func TestConvertIntToString(t *testing.T) {
	var intRef = 7
	var expected = "7"

	val := ConvertIntToString(intRef)
	t.Logf("Value converted %s", val)

	if val != expected {
		t.Errorf("Value '%s' is not the same as ref '%d'", val, intRef)
	}
}

// Test string percentage conversion to float64
func TestConvertPercentageToFloat64(t *testing.T) {
	var stringRef = "15 %"
	var expected = 15.

	val, err := ConvertPercentageToFloat64(stringRef)
	t.Logf("Value converted %f", val)

	if err != nil {
		t.Errorf("Impossible to convert this string %s, err: %v", stringRef, err)
	} else if val != expected {
		t.Errorf("Value '%f' is not the same as ref '%s'", val, stringRef)
	}
}

// Test string percentage conversion to float64
func TestConvertBadPercentageToFloat64(t *testing.T) {
	var stringRef = "15%"
	var expected = 0.

	val, err := ConvertPercentageToFloat64(stringRef)
	t.Logf("Value converted %f", val)

	if err == nil {
		t.Errorf("Conversion has to generate error, err: %v", err)
	} else if val != expected {
		t.Errorf("Value '%f' is not the same as ref '%s'", val, stringRef)
	}
}

// Test if function return the maxLength among this item
func TestMaxLength(t *testing.T) {
	var longItem = "test max length"
	var items = []string{"tax", "options", longItem, "db clean"} // 'long item' has to be the long string in array
	var refMaxLength = len(longItem)
	t.Logf("Length of %s : %d", longItem, refMaxLength)

	var maxLength = GetMaxLength(items)
	t.Logf("Length find by function: %d", maxLength)

	if refMaxLength != maxLength {
		t.Errorf("The refMaxLength '%d' is not the same has maxLength '%d'", refMaxLength, maxLength)
	}
}

// Test if function return the right padding
func TestGetPadding(t *testing.T) {
	var longItem = "test get padding function"
	var items = []string{"tax", "options", longItem, "db clean"} // 'long item' has to be the long string in array
	var refGetPadding = len(longItem)
	t.Logf("Length of %s : %d", longItem, refGetPadding)

	var padding = getPadding(items)
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
	var samePaddings = func(a []int) bool {
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

func TestDownloadWithGoodUrl(t *testing.T) {
	var url = "https://github.com/NY-Daystar/corpos-christie/releases/download/v2.1.0/windows-corpos-christie-2.0.0.zip"
	var destFile = "./release.zip"
	var expected = 0

	statusCode, err := DownloadFile(url, destFile)
	t.Logf("Status code %d", statusCode)

	if statusCode != expected {
		t.Errorf("Bad status code for url %s - status code : %d <> %d", url, statusCode, expected)
	}
	if err != nil {
		t.Errorf("Unable to download url %s, err: %v", url, err)
	}
}

func TestDownloadWithWrongUrl(t *testing.T) {
	var url = "https://github.com/NY-Daystar/WRONG-REPOSITORY/tree/"
	var destFile = "./release.zip"
	var expected = 404

	statusCode, _ := DownloadFile(url, destFile)

	if statusCode != expected {
		t.Errorf("Bad status code for url %s - status code : %d <> %d", url, statusCode, expected)
	}
}

func TestDownloadWithWrongProtocol(t *testing.T) {
	var url = "Wrong-protocol://github.com/NY-Daystar/corpos-christie/releases/download/v2.1.0/linux-corpos-christie-2.0.0.zip"
	var destFile = "./dest.zip"

	_, err := DownloadFile(url, destFile)
	t.Logf("Error: %v", err)

	if err == nil {
		t.Error("No error detected")
	}
}
