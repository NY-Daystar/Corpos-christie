package utils

import (
	"fmt"
	"strconv"
	"time"
)

// Read input from terminal
func ReadValue() string {
	var value string
	fmt.Scanf("%s", &value)
	return value
}

// Convert string value to int
func ConvertStringToInt(str string) (int, error) {
	f, err := strconv.Atoi(str)
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
