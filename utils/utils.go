// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package utils define functions to multiple uses
package utils

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
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

// Regex validation for email
func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
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

// src [string] : zip path
// dest [string] : folder path to unzip
//
// Unzipped zip file to a folder
func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		if strings.Contains(fpath, "..") {
			continue
		}

		// Si c'est un dossier, crée-le
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Si c'est un fichier, extrait-le
		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.CopyN(outFile, rc, 1024)

		// Fermer les fichiers ouverts
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

// CopyFolder: copy all files into destDir
func CopyFolder(srcDir, destDir string) error {
	err := filepath.Walk(srcDir, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDir, srcPath)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDir, relPath)

		if info.IsDir() {
			if err := os.MkdirAll(destPath, info.Mode()); err != nil {
				return err
			}
		} else {
			if _, err := os.Stat(destPath); err == nil {
				if err := os.Remove(destPath); err != nil {
					return err
				}
			}

			if err := copyFile(srcPath, destPath); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

// copyFile: copy on file to dest folder
func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	sourceFile.Close()
	destFile.Close()
	return os.Chmod(dest, info.Mode())
}

// copyFile: Get absolute path of executable
func GetExecutablePath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("path of executable not found: %v", err)
	}

	execPath, err = filepath.EvalSymlinks(execPath)
	if err != nil {
		return "", fmt.Errorf("symbolic link not resolved: %v", err)
	}

	execDir := filepath.Dir(execPath)
	// fmt.Println("Chemin complet de l'exécutable :", execPath)
	// fmt.Println("Répertoire de l'exécutable :", execDir)

	return execDir, nil
}
