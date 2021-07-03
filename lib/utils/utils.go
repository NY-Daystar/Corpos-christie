package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	DEFAULT_PADDING = 10 // Add defautl padding for function setPadding
)

// Read input from terminal
func ReadValue() string {
	var value string
	fmt.Scanf("%s", &value)
	return value
}

// Convert string value to int
func ConvertStringToInt(str string) (int, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// Convert string percentage like 5% into 5
func ConvertPercentageToFloat64(str string) (float64, error) {
	var s string = strings.TrimSuffix(str, "%")
	i, err := strconv.Atoi(s)
	f := float64(i)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// Return current year
func GetCurrentYear() int {
	year, _, _ := time.Now().Date()
	return year
}

// Get Max length from slice string
func GetMaxLength(tab []string) int {
	var maxIndexLength int
	for _, v := range tab {
		if maxIndexLength < len(v) {
			maxIndexLength = len(v)
		}
	}
	return maxIndexLength
}

// Get padding between value in tab for each of them to align items
func getPadding(tab []string) int {
	return GetMaxLength(tab)
}

// Set a padding with a value among a list of data
func SetPadding(tab []string, v string) string {
	var padding int = getPadding(tab)
	var gap int = padding - len(v) + DEFAULT_PADDING
	var space string = strings.Repeat(" ", gap)
	return space
}
