package themes

import (
	"image/color"

	"fyne.io/fyne/v2"
)

// To create a theme
// $ go get github.com/lusingander/fyne-theme-generator
// $ go run github.com/lusingander/fyne-theme-generator
// Theme define fyne theme between (Light and Dark)
type Theme interface {
	Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color
	Font(s fyne.TextStyle) fyne.Resource
	Icon(n fyne.ThemeIconName) fyne.Resource
	Size(s fyne.ThemeSizeName) float32
}
