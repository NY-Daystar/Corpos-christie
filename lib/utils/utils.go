// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package utils define functions to multiple uses
package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	// Add default padding for function setPadding
	DEFAULT_PADDING = 10
)

// ReadValue read input from terminal and returns its value
func ReadValue() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// ConvertStringToInt convert str string to an int and returns it
// return an error if the string is not convertible into an int
func ConvertStringToInt(str string) (int, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// ConvertFloat64ToString convert float64 to a string and returns it
func ConvertFloat64ToString(v float64) string {
	return fmt.Sprintf("%f", v)
}

// ConvertInt64ToString convert int64 to a string and returns it
func ConvertInt64ToString(v int64) string {
	// return fmt.Sprintf("%f", v)
	return strconv.FormatInt(v, 10)
}

// ConvertIntToString convert int to a string and returns it
func ConvertIntToString(v int) string {
	return fmt.Sprintf("%d", v)
}

// ConvertPercentageToFloat64 convert str which is string percentage like 5% into 5
func ConvertPercentageToFloat64(str string) (float64, error) {
	var s string = strings.TrimSuffix(str, "%")
	i, err := strconv.Atoi(s)
	f := float64(i)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// GetCurrentYear returns current year (ex: 2021)
func GetCurrentYear() int {
	year, _, _ := time.Now().Date()
	return year
}

// GetMaxLength get max length string among the tab slice and returns its length
func GetMaxLength(tab []string) int {
	var maxIndexLength int
	for _, v := range tab {
		if maxIndexLength < len(v) {
			maxIndexLength = len(v)
		}
	}
	return maxIndexLength
}

// getPadding get padding necessary between values in tab for each of them to align items
func getPadding(tab []string) int {
	return GetMaxLength(tab)
}

// SetPadding get the padding of the tab slice and add the padding into the element v
// returns v string including the padding
func SetPadding(tab []string, v string) string {
	var padding int = getPadding(tab)
	var gap int = padding - len(v) + DEFAULT_PADDING
	var space string = strings.Repeat(" ", gap)
	return space
}
