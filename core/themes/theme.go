package themes

import (
	"image/color"

	"fyne.io/fyne/v2"
)

const (
	DARK  string = "Dark"
	LIGHT string = "Light"
)

// Theme define fyne theme between (Light and Dark)
type Theme interface {
	Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color
	Font(s fyne.TextStyle) fyne.Resource
	Icon(n fyne.ThemeIconName) fyne.Resource
	Size(s fyne.ThemeSizeName) float32
}
