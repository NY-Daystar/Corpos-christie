package settings

// Handle the icon in GUI settings

import (
	"fyne.io/fyne/v2"
)

// GetIcon Load icon file to show in window
// Returns icon in fyne object
func GetIcon(path string) fyne.Resource {
	icon, _ := fyne.LoadResourceFromPath(path)
	return icon
}
