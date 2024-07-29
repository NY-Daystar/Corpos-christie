package updater

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
func StartUpdater(logger *zap.Logger) {
	release, err := checkForUpdate(config.APP_VERSION)
	if err != nil {
		logger.Sugar().Errorf("Error checking for update: %v", err)
	}

	if release != nil {
		fmt.Println("New version available:", release.TagName)
		for _, asset := range release.Assets {
			fmt.Println("Found asset:", asset.Name)
			// Adjust to find the correct binary for the OS
			if (runtime.GOOS == "windows" && asset.Name == "windows-corpos-christie-2.0.0.zip") ||
				(runtime.GOOS == "linux" && asset.Name == "app-linux") ||
				(runtime.GOOS == "darwin" && asset.Name == "app-macos") {
				fmt.Println("Downloading from:", asset.URL)

				var dest string
				if runtime.GOOS == "windows" {
					dest = "./updater/output.zip" // TODO mettre le dossier de telechargement de l'OS
				} else {
					dest = "./updater/output" // TODO mettre le dossier de telechargement de l'OS
				}

				typee, err := utils.DownloadFile(asset.URL, dest)
				if err != nil {
					fmt.Printf("Type %d, Error downloading the update: %v\n", typee, err)
					return
				}

				fmt.Println("Downloaded new version")

				// TODO execute MSI

				// Replace the old binary with the new one
				// execPath, err := os.Executable()
				// if err != nil {
				// 	fmt.Println("Error finding executable path:", err)
				// 	return
				// }

				// err = os.Rename(dest, execPath)
				// if err != nil {
				// 	fmt.Println("Error replacing the binary:", err)
				// 	return
				// }

				// fmt.Println("Update successful, restarting application")
				// cmd := exec.Command(execPath)
				// cmd.Start()
				// os.Exit(0)
			}
		}
	} else {
		// if no internet access
		fmt.Println("No new version available")
	}
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

	tag, err := go_version.NewVersion(releaseTag)
	if err != nil {
		fmt.Printf("3 %v", err)
	}
	versionCompared, err := go_version.NewVersion(version)
	if err != nil {
		fmt.Printf("4 %v", err)
	}

	if tag.LessThan(versionCompared) {
		fmt.Printf("Actual version (%s) is less than latest version (%s)\n", tag, versionCompared)
		// return &release, nil
	} else {
		fmt.Printf("Actual version (%s) is greater or equal than latest version (%s)\n", tag, versionCompared)
		return &release, nil
	}
	// if tag > version {
	// 	return &release, nil
	// }
	return nil, resp.Body.Close()
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
