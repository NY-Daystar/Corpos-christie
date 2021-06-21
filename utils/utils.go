package utils

import (
	"fmt"
	"strconv"
)

// Read input from terminal
func ReadValue() string {
	var value string
	fmt.Scanf("%s", &value)
	return value
}

// TODO a mettre dans un folder utils
// Convert string value to int
func ConvertStringToInt(str string) (int, error) {
	f, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return f, nil
}
