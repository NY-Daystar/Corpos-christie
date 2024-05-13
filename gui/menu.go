package gui

import (
	"fmt"
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/settings"
)

// GUIMenu represents menu of window application
type GUIMenu struct {
	Controller *GUIController
	App        fyne.App    // Fyne application
	Window     fyne.Window // Fyne window
}

// NewMenu create main menu for window application
func NewMenu(controller *GUIController) *GUIMenu {
	return &GUIMenu{
		Controller: controller,
		App:        controller.View.App,
		Window:     controller.View.Window,
	}
}

func (menu *GUIMenu) Start() *fyne.MainMenu {
	fmt.Printf("START\n")
	return fyne.NewMainMenu(
		menu.createFileMenu(),
		menu.createHelpMenu(),
	)
}

// createFileMenu create file item in toolbar to handle app settings
func (menu *GUIMenu) createFileMenu() *fyne.Menu {
	fmt.Printf("createFileMenu\n")
	fileMenu := fyne.NewMenu(menu.Controller.Model.Language.File,
		fyne.NewMenuItem(menu.Controller.Model.Language.Settings, func() {
			fmt.Printf("createFileMenu\n")
			dialog.ShowCustom(menu.Controller.Model.Language.Settings, menu.Controller.Model.Language.Close,
				container.NewVBox(
					menu.createSelectTheme(),
					widget.NewSeparator(),
					menu.createSelectLanguage(),
					widget.NewSeparator(),
					menu.createSelectCurrency(),
					widget.NewSeparator(),
					menu.createLabelLogs(),
				), menu.Window)
		}),
		fyne.NewMenuItem(menu.Controller.Model.Language.Quit, func() { menu.App.Quit() }),
	)
	return fileMenu
}

// createHelpMenu create help item in toolbar to show about app
func (menu *GUIMenu) createHelpMenu() *fyne.Menu {
	url, _ := url.Parse(config.APP_LINK)

	menu.Controller.Model.LabelsAbout = binding.NewStringList()
	menu.Controller.Model.LabelsAbout.Set(menu.Controller.Model.Language.GetAbouts())
	var labels []binding.DataItem
	for index := range menu.Controller.Model.Language.GetAbouts() {
		about, _ := menu.Controller.Model.LabelsAbout.GetItem(index)
		labels = append(labels, about)
	}

	// Setup layouts with data
	firstLine := container.NewHBox(
		widget.NewLabelWithData(labels[0].(binding.String)),
		widget.NewLabel(config.APP_NAME),
		widget.NewLabelWithData(labels[1].(binding.String)),
	)
	secondLine := container.NewHBox(
		widget.NewLabelWithData(labels[2].(binding.String)),
		widget.NewHyperlink("GitHub Project", url),
		widget.NewLabelWithData(labels[3].(binding.String)),
	)
	thirdLine := widget.NewLabelWithData(labels[4].(binding.String))
	fourthLine := container.NewHBox(
		widget.NewLabel("Version:"),
		canvas.NewText(fmt.Sprintf("v%s", config.APP_VERSION), color.NRGBA{R: 218, G: 20, B: 51, A: 255}),
	)
	fifthLine := widget.NewLabel(fmt.Sprintf("%s: %s", menu.Controller.Model.Language.Author, config.APP_AUTHOR))

	helpMenu := fyne.NewMenu(menu.Controller.Model.Language.Help,
		fyne.NewMenuItem(menu.Controller.Model.Language.About, func() {
			dialog.ShowCustom(menu.Controller.Model.Language.About, menu.Controller.Model.Language.Close,
				container.NewVBox(
					firstLine,
					secondLine,
					thirdLine,
					fourthLine,
					fifthLine,
				), menu.Window)
		}))
	return helpMenu
}

// createSelectTheme create select to change theme
func (menu *GUIMenu) createSelectTheme() *fyne.Container {
	selectTheme := widget.NewSelect(menu.Controller.Model.Language.GetThemes(), nil)

	selectTheme.OnChanged = func(s string) {
		index := selectTheme.SelectedIndex()
		menu.Controller.SetTheme(index) // update model
	}

	selectTheme.SetSelectedIndex(menu.Controller.Model.Settings.Theme)

	return container.NewHBox(
		widget.NewLabel(menu.Controller.Model.Language.ThemeCode),
		selectTheme,
	)
}

// createSelectLanguage create select to change language
func (menu *GUIMenu) createSelectLanguage() *fyne.Container {
	selectLanguage := widget.NewSelect(menu.Controller.Model.Language.GetLanguages(), nil)
	selectLanguage.SetSelectedIndex(menu.Controller.Model.GetLanguageIndex(menu.Controller.Model.Language.Code))
	selectLanguage.OnChanged = func(s string) {
		index := selectLanguage.SelectedIndex()
		var getLanguage = func() string {
			switch index {
			case 0:
				return settings.ENGLISH
			case 1:
				return settings.FRENCH
			case 2:
				return settings.SPANISH
			case 3:
				return settings.GERMAN
			case 4:
				return settings.ITALIAN
			default:
				return settings.ENGLISH
			}
		}

		language := getLanguage()
		menu.Controller.SetLanguage(language) // Update model
	}

	return container.NewHBox(
		widget.NewLabel(menu.Controller.Model.Language.LanguageCode),
		selectLanguage,
	)
}

// createSelectCurrency create select to change currency
func (menu *GUIMenu) createSelectCurrency() *fyne.Container {
	selectCurrency := widget.NewSelect(settings.GetCurrencies(), func(currency string) {
		menu.Controller.SetCurrency(currency) // Update model
	})

	currency, _ := menu.Controller.Model.Currency.Get()
	selectCurrency.SetSelected(currency)
	return container.NewHBox(
		widget.NewLabel(menu.Controller.Model.Language.Currency),
		selectCurrency,
	)
}

// createLabelLogs create label to show logs
func (menu *GUIMenu) createLabelLogs() *fyne.Container {
	return container.NewHBox(
		widget.NewLabel(menu.Controller.Model.Language.Logs),
		widget.NewLabel(config.LOGS_PATH),
	)
}
