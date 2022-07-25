package gui

import (
	"fmt"
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/LucasNoga/corpos-christie/config"
)

// TODO a mettre dans le package widgets

// SetMenu Create mainMenu for window
func (g *GUI) SetMenu() *fyne.MainMenu {
	// TODO a split en petite fonction
	selectTheme := widget.NewSelect(g.Language.Theme.GetThemes(), func(val string) {
		g.SetTheme(val)
		// TODO save data in .settings
	})
	selectTheme.SetSelected(g.ThemeName)

	selectLanguage := widget.NewSelect(g.Language.Languages.GetLanguages(), nil)
	selectLanguage.OnChanged = func(s string) {
		index := selectLanguage.SelectedIndex()
		var getLanguage = func() string {
			switch index {
			case 0:
				return "en" // TODO enum
			case 1:
				return "fr"
			}
			// TODO error log
			return "fr"
		}

		language := getLanguage()

		g.SetLanguage(language)
		g.Reload()
		// TODO save data in .settings

	}

	fileMenu := fyne.NewMenu(g.Language.File,
		fyne.NewMenuItem(g.Language.Settings, func() {
			dialog.ShowCustom(g.Language.Settings, g.Language.Close, container.NewVBox(
				container.NewHBox(
					widget.NewLabel(g.Language.ThemeCode),
					selectTheme,
				),
				widget.NewSeparator(),
				container.NewHBox(
					widget.NewLabel(g.Language.LanguageCode),
					selectLanguage,
				),
			), g.Window)
		}),
		fyne.NewMenuItem(g.Language.Quit, func() { g.App.Quit() }),
	)

	url, _ := url.Parse(config.APP_LINK)

	helpMenu := fyne.NewMenu(g.Language.Help,
		fyne.NewMenuItem(g.Language.About, func() {
			dialog.ShowCustom(g.Language.About, g.Language.Close, container.NewVBox(
				widget.NewLabel(fmt.Sprintf("Welcome to %s, a Desktop app to calculate your taxes in France.", config.APP_NAME)), // TODO language
				container.NewHBox(
					widget.NewLabel("This"),                    // TODO language
					widget.NewHyperlink("GitHub Project", url), // TODO language
					widget.NewLabel("is open-source."),         // TODO language
				),
				widget.NewLabel("Developped in Go with Fyne."),
				container.NewHBox(
					widget.NewLabel("Version:"),
					canvas.NewText(fmt.Sprintf("v%s", config.APP_VERSION), color.NRGBA{R: 218, G: 20, B: 51, A: 255}),
				),
				widget.NewLabel(fmt.Sprintf("%s: %s", g.Language.Author, config.APP_AUTHOR)),
			), g.Window)
		}))
	return fyne.NewMainMenu(fileMenu, helpMenu)
}
