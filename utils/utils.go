// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package utils define functions to multiple uses
package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/data/binding"
	"github.com/NY-Daystar/corpos-christie/config"
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

// ConvertStringToFloat64 convert str string to a float64 and returns it
// return an error if the string is not convertible into an float64
func ConvertStringToFloat64(str string) (float64, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return float64(i), nil
}

// ConvertBindStringToInt convert bindstring (fyne) to an int and returns it
// return an error if the string is not convertible into an int
func ConvertBindStringToInt(str binding.String) int {
	bstr, _ := str.Get()
	i, err := strconv.Atoi(bstr)
	if err != nil {
		return 0
	}
	return i
}

// ConvertFloat64ToString convert float64 to a string and returns it
func ConvertFloat64ToString(v float64) string {
	return fmt.Sprintf("%f", v)
}

// ConvertInt64ToString convert int64 to a string and returns it
func ConvertInt64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

// ConvertIntToString convert int to a string and returns it
func ConvertIntToString(v int) string {
	return fmt.Sprintf("%d", v)
}

// ConvertPercentageToFloat64 convert str which is string percentage like 5% into 5
func ConvertPercentageToFloat64(str string) (float64, error) {
	var s = strings.TrimSuffix(str, " %")
	i, err := strconv.Atoi(s)
	f := float64(i)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// FindIndex get index in slice of string if target is in
// If not found return -1
func FindIndex(slice []string, target string) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
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

// ReadFileLastNLines read the last N lines of the files
func ReadFileLastNLines(filePath string, n int) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	// Keep only last N lines
	if len(lines) > n {
		lines = lines[len(lines)-n:]
	}

	// Add line number
	var result string
	for i, line := range lines {
		result += fmt.Sprintf("%3d: %s\n", i+1, line)
	}

	return result, file.Sync()
}

// Read history files and return list of string
func GetHistory(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Sync()

	return lines
}

// getPadding get padding necessary between values in tab for each of them to align items
func getPadding(tab []string) int {
	return GetMaxLength(tab)
}

// SetPadding get the padding of the tab slice and add the padding into the element v
// returns v string including the padding
func SetPadding(tab []string, v string) string {
	var padding = getPadding(tab)
	var gap = padding - len(v) + DEFAULT_PADDING
	var space = strings.Repeat(" ", gap)
	return space
}

// getAppDataPath returns the path to the AppData directory on Windows
// and the home directory on Linux.
func GetAppDataPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if runtime.GOOS == "windows" {
		appDataPath := os.Getenv("APPDATA")
		if appDataPath == "" {
			return "", fmt.Errorf("APPDATA environment variable is not set")
		}
		return appDataPath, nil
	} else {
		return homeDir, nil
	}
}

// get logs filePath
func GetLogsFile() string {
	appDataPath, _ := GetAppDataPath()
	var logsFolder = path.Join(appDataPath, config.APP_NAME, "logs")
	return path.Join(logsFolder, "log.json")
}

// get history filePath
func GetHistoryFile() string {
	appDataPath, _ := GetAppDataPath()
	var appFolder = path.Join(appDataPath, config.APP_NAME)
	var historyPath = path.Join(appFolder, "history.json")
	if _, err := os.Stat(historyPath); err != nil {
		os.Create(historyPath)
	}
	return historyPath
}

// get setings filePath
func GetSettingsFile() string {
	appDataPath, _ := GetAppDataPath()
	var appFolder = path.Join(appDataPath, config.APP_NAME)
	var settingsPath = path.Join(appFolder, "settings.json")
	if _, err := os.Stat(settingsPath); err != nil {
		os.Create(settingsPath)
	}
	return settingsPath
}

// delete file with his path
func DeleteFile(filePath string) {
	os.Remove(filePath)
}

// DownloadFile from url to destination return int and error
// If success then return 0 and no error
func DownloadFile(url, dest string) (int, error) {
	out, err := os.Create(dest)
	if err != nil {
		return 1, err
	}

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode == 404 {
		return 404, err
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return 3, err
	}
	resp.Body.Close()
	return 0, out.Sync()
}
