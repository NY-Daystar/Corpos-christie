package updater

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/utils"
	go_version "github.com/hashicorp/go-version"

	"go.uber.org/zap"
)

// GitHubRelease data relative to release github
type GitHubRelease struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name string `json:"name"`
		URL  string `json:"browser_download_url"`
	} `json:"assets"`
}

// StartUpdater check if update is available and install it
func StartUpdater(logger *zap.Logger) (string, error) {
	release, err := checkForUpdate(config.APP_VERSION)
	if err != nil {
		return "", fmt.Errorf("error checking for update: %v", err)
	}
	if release == nil {
		return "", fmt.Errorf("no update available")
	}

	zipPath, err := downloadRelease(release)
	if err != nil {
		return "", fmt.Errorf("error download release: %v", err)
	}

	appPath, err := installRelease(zipPath)
	if err != nil {
		return "", fmt.Errorf("error install release: %v", err)
	}

	return appPath, nil
}

//	version [string]: actual version of application and check if superior one exists
//	returns all data of github release if new version is available otherwise return false
//
// checkForUpdate verify with github api if new release is available and download it
func checkForUpdate(version string) (*GitHubRelease, error) {
	url := "https://api.github.com/repos/ny-daystar/corpos-christie/releases/latest"
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	var release GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}

	var releaseTag = GetTag(release.TagName)

	if releaseTag == "" {
		fmt.Println("No release, reached rate limiter")
		return nil, errors.New("rate limiter reached")
	}

	tag, _ := go_version.NewVersion(releaseTag)
	actualVersion, _ := go_version.NewVersion(version)

	if tag.LessThan(actualVersion) {
		fmt.Printf("Actual version (%s) is greater or equal than latest version (%s). No need to update\n", actualVersion, tag)
	} else {
		fmt.Printf("Actual version (%s) is lower than latest version (%s)\n", actualVersion, tag)
		return &release, nil
	}
	return nil, resp.Body.Close()
}

//	release [GitHubRelease]: release of new version
//	returns path to zip locally
//
// downloadRelease Download link of release and return path of zip
func downloadRelease(release *GitHubRelease) (string, error) {
	if release == nil {
		return "", fmt.Errorf("No new version available")
	}

	fmt.Println("New version available:", release.TagName)

	var arch = getArchitecture()
	fmt.Printf("Arch : %v -> %v\n", runtime.GOOS, arch)

	var delivery = fmt.Sprintf("%s-%s-%s.zip", arch, config.APP_NAME, GetTag(release.TagName))
	fmt.Printf("Asset to download : %s\n", delivery)

	var downloadLink = ""
	for _, asset := range release.Assets {
		fmt.Println("Found asset:", asset.Name)
		if asset.Name == delivery {
			downloadLink = asset.URL
			break
		}
	}

	if downloadLink == "" {
		return "", fmt.Errorf("download link not found")
	}

	fmt.Println("Downloading from:", downloadLink)

	downloadFolder, _ := getDownloadFolder()

	var dest = path.Join(downloadFolder, fmt.Sprintf("%s-%s.zip", config.APP_NAME, release.TagName))

	typee, err := utils.DownloadFile(downloadLink, dest)
	if err != nil {
		fmt.Printf("Type %d, Error downloading the update: %v\n", typee, err)
		return "", err
	}

	fmt.Println("Downloaded new version")
	return dest, nil
}

//	zipPath [string]: path to release downloaded
//	returns path to folder unzipped
//
// installRelease unzip file and copy into
func installRelease(zipPath string) (string, error) {
	downloadFolder, err := getDownloadFolder()
	if err != nil {
		return "", err
	}
	var unzipPath = path.Join(downloadFolder, config.APP_NAME)
	utils.Unzip(zipPath, unzipPath)
	return unzipPath, nil
}

// Get version of release based on tag
// If no matching return ""
func GetTag(tag string) string {
	var re = regexp.MustCompile(`v(\d+\.\d+\.\d+)`)
	var result = re.FindStringSubmatch(tag)
	if len(result) == 0 {
		return ""
	}
	return result[1]
}

//	version [string]: actual version of application and check if superior one exists
//	returns true if new version is available otherwise return false
//
// IsNewUpdateAvailable Check for GUI and return true if new update
func IsNewUpdateAvailable(version string) (bool, error) {
	release, err := checkForUpdate(version)

	if err != nil || release == nil {
		return false, err
	}

	return true, nil
}

// Return good name with right architecture
func getArchitecture() string {
	switch a := runtime.GOOS; a {
	case "windows":
		return "windows"
	case "linux":
		return "linux"
	case "darwin":
		return "mac"
	default:
		return "windows"
	}
}

// GetDownloadFolder returns path of folder "Download" from OS
func getDownloadFolder() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	switch runtime.GOOS {
	case "windows":
		return filepath.Join(homeDir, "Downloads"), nil
	case "darwin": // macOS
		return filepath.Join(homeDir, "Downloads"), nil
	case "linux":
		return filepath.Join(homeDir, "Downloads"), nil
	default:
		return "", fmt.Errorf("OS unknown: %s", runtime.GOOS)
	}
}
