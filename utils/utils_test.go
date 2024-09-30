// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package utils define functions to multiple uses
package utils

import (
	"reflect"
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

func TestDownloadWithGoodUrl(t *testing.T) {
	var url = "https://github.com/NY-Daystar/corpos-christie/releases/download/v2.1.0/windows-corpos-christie-2.1.0.zip"
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
	var destFile = "./downloaded.zip"

	_, err := DownloadFile(url, destFile)
	t.Logf("Error: %v", err)

	if err == nil {
		t.Error("No error detected")
	}
}

func TestFindIndex(t *testing.T) {
	tests := []struct {
		slice    []string
		target   string
		expected int
	}{
		{
			slice:    []string{"one", "two", "three"},
			target:   "one",
			expected: 0,
		},
		{
			slice:    []string{"one", "two", "three"},
			target:   "two",
			expected: 1,
		},
		{
			slice:    []string{"one", "two", "three"},
			target:   "three",
			expected: 2,
		},
		{
			slice:    []string{"one", "two", "three"},
			target:   "four",
			expected: -1,
		},
	}

	for _, testCase := range tests {
		var index = FindIndex(testCase.slice, testCase.target)
		t.Logf("Index found: %v", index)

		if testCase.expected != index {
			t.Errorf("Test case failed with result: %d - expected: %d", index, testCase.expected)
		}
	}
}

func TestValidationMail(t *testing.T) {
	tests := []struct {
		mail     string
		expected bool
	}{
		{
			mail:     "123@gmail.com",
			expected: true,
		},
		{
			mail:     "mymail@gmail.com",
			expected: true,
		},
		{
			mail:     "my.mail@gmail.com",
			expected: true,
		},
		{
			mail:     "my_mail@gmail.com",
			expected: true,
		},
		{
			mail:     "my.mailgmail.com",
			expected: false,
		},
		{
			mail:     "my.mail@gmailcom",
			expected: false,
		},
		{
			mail:     "my.mailgmailcom",
			expected: false,
		},
	}

	for _, testCase := range tests {
		var check = IsValidEmail(testCase.mail)

		if testCase.expected != check {
			t.Errorf("Test case failed with mail: %s - expected: %v", testCase.mail, testCase.expected)
		}
	}
}

func TestFilePath(t *testing.T) {
	tests := map[string]interface{}{
		"GetAppDataPath":  GetAppDataPath,
		"GetLogsFile":     GetLogsFile,
		"GetHistoryFile":  GetHistoryFile,
		"GetSettingsFile": GetSettingsFile,
	}

	for _, method := range tests {
		t.Logf("Méthod called %v", method)
		callMethodByName(method) // Using reflexion to call method
	}
}

// use reflexion to call method
func callMethodByName(name interface{}) {
	method := reflect.ValueOf(name)
	method.Call(nil)
}
