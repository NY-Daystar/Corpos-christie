package themes

import (
	"testing"

	"fyne.io/fyne/v2"
)

// For testing
// $ cd gui/themes
// $ go test -v

func TestDarkTheme(t *testing.T) {
	var theme = DarkTheme{}
	theme.Color("background", 0)
	theme.Color("button", 0)
	theme.Color("error", 0)
	theme.Color("disabled", 0)
	theme.Color("disabledButton", 0)
	theme.Color("focus", 0)
	theme.Color("foreground", 0)
	theme.Color("hover", 0)
	theme.Color("inputBackground", 0)
	theme.Color("placeholder", 0)
	theme.Color("pressed", 0)
	theme.Color("primary", 0)
	theme.Color("scrollBar", 0)
	theme.Color("shadow", 0)

	theme.Font(fyne.TextStyle{Bold: true})
	theme.Font(fyne.TextStyle{Italic: true})
	theme.Font(fyne.TextStyle{Bold: true, Italic: true})
	theme.Font(fyne.TextStyle{Monospace: true})
	theme.Font(fyne.TextStyle{Symbol: true})

	theme.Icon(fyne.ThemeIconName(""))

	theme.Size(fyne.ThemeSizeName("helperText"))
	theme.Size(fyne.ThemeSizeName("iconInline"))
	theme.Size(fyne.ThemeSizeName("innerPadding"))
	theme.Size(fyne.ThemeSizeName("lineSpacing"))
	theme.Size(fyne.ThemeSizeName("padding"))
	theme.Size(fyne.ThemeSizeName("scrollBar"))
	theme.Size(fyne.ThemeSizeName("scrollBarSmall"))
	theme.Size(fyne.ThemeSizeName("separator"))
	theme.Size(fyne.ThemeSizeName("text"))
	theme.Size(fyne.ThemeSizeName("headingText"))
	theme.Size(fyne.ThemeSizeName("subHeadingText"))
	theme.Size(fyne.ThemeSizeName("inputBorder"))
	t.Logf("%v", theme)
}

func TestLightTheme(t *testing.T) {
	var theme = &LightTheme{}
	theme.Color("background", 0)
	theme.Color("button", 0)
	theme.Color("error", 0)
	theme.Color("disabled", 0)
	theme.Color("disabledButton", 0)
	theme.Color("focus", 0)
	theme.Color("foreground", 0)
	theme.Color("hover", 0)
	theme.Color("inputBackground", 0)
	theme.Color("placeholder", 0)
	theme.Color("pressed", 0)
	theme.Color("primary", 0)
	theme.Color("scrollBar", 0)
	theme.Color("shadow", 0)

	theme.Font(fyne.TextStyle{Bold: true})
	theme.Font(fyne.TextStyle{Italic: true})
	theme.Font(fyne.TextStyle{Bold: true, Italic: true})
	theme.Font(fyne.TextStyle{Monospace: true})
	theme.Font(fyne.TextStyle{Symbol: true})

	theme.Icon(fyne.ThemeIconName(""))

	theme.Size(fyne.ThemeSizeName("helperText"))
	theme.Size(fyne.ThemeSizeName("iconInline"))
	theme.Size(fyne.ThemeSizeName("innerPadding"))
	theme.Size(fyne.ThemeSizeName("lineSpacing"))
	theme.Size(fyne.ThemeSizeName("padding"))
	theme.Size(fyne.ThemeSizeName("scrollBar"))
	theme.Size(fyne.ThemeSizeName("scrollBarSmall"))
	theme.Size(fyne.ThemeSizeName("separator"))
	theme.Size(fyne.ThemeSizeName("text"))
	theme.Size(fyne.ThemeSizeName("headingText"))
	theme.Size(fyne.ThemeSizeName("subHeadingText"))
	theme.Size(fyne.ThemeSizeName("inputBorder"))
	t.Logf("%v", theme)
}
