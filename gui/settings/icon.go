package settings

// Handle the icon in GUI settings

import (
	"fmt"

	"fyne.io/fyne/v2"
	"github.com/LucasNoga/corpos-christie/config"
)

// GetIcon Load icon file to show in window
// Returns icon in fyne object
func GetIcon() fyne.Resource {
	var iconName string = "logo.ico"
	var iconPath string = fmt.Sprintf("%s/%s", config.ASSETS_PATH, iconName)
	icon, _ := fyne.LoadResourceFromPath(iconPath)
	// TODO log debug to show icon loaded
	return icon
}
