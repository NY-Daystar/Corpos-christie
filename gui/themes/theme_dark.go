package themes

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type DarkTheme struct{}

func (DarkTheme) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	if c == theme.ColorNameBackground {
		return color.NRGBA{R: 0xf, G: 0xf, B: 0x15, A: 0xff}
	}
	if c == theme.ColorNameButton {
		return color.Alpha16{A: 0x0}
	}
	if c == theme.ColorNameDisabled {
		return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x42}
	}
	if c == theme.ColorNameError {
		return color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	}
	if c == theme.ColorNameDisabledButton {
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x42}
	}
	if c == theme.ColorNameFocus {
		return color.NRGBA{R: 0x21, G: 0x96, B: 0xf3, A: 0x7f}
	}
	if c == theme.ColorNameForeground {
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	}
	if c == theme.ColorNameHover {
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xf}
	}
	if c == theme.ColorNameInputBackground {
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x19}
	}
	if c == theme.ColorNamePlaceHolder {
		return color.NRGBA{R: 0xb2, G: 0xb2, B: 0xb2, A: 0xff}
	}
	if c == theme.ColorNamePressed {
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x66}
	}
	if c == theme.ColorNamePrimary {
		return color.NRGBA{R: 0x21, G: 0x96, B: 0xf3, A: 0xff}
	}
	if c == theme.ColorNameScrollBar {
		return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x99}
	}
	if c == theme.ColorNameShadow {
		return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x66}
	}
	if c == theme.ColorNameSelection {
		return color.NRGBA{R: 0xf, G: 0xf, B: 0x15, A: 0xff}
	}
	if c == theme.ColorNameMenuBackground {
		return color.NRGBA{R: 0xf, G: 0xf, B: 0x15, A: 0xff}
	}
	if c == theme.ColorNameOverlayBackground {
		return color.NRGBA{R: 0xf, G: 0xf, B: 0x15, A: 0xff}
	} else {
		return theme.DefaultTheme().Color(c, v)
	}
}

func (DarkTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return theme.DefaultTheme().Font(s)
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return theme.DefaultTheme().Font(s)
}

func (DarkTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (DarkTheme) Size(s fyne.ThemeSizeName) float32 {
	switch s {
	case theme.SizeNameCaptionText:
		return 11
	case theme.SizeNameInlineIcon:
		return 20
	case theme.SizeNamePadding:
		return 4
	case theme.SizeNameScrollBar:
		return 16
	case theme.SizeNameScrollBarSmall:
		return 3
	case theme.SizeNameSeparatorThickness:
		return 1
	case theme.SizeNameText:
		return 14
	case theme.SizeNameInputBorder:
		return 2
	default:
		return theme.DefaultTheme().Size(s)
	}
}
