// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package gui defines component and script to launch gui application
package gui

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/gui/settings"
	"github.com/LucasNoga/corpos-christie/gui/themes"
	"github.com/LucasNoga/corpos-christie/lib/utils"
	"github.com/LucasNoga/corpos-christie/tax"
	"github.com/LucasNoga/corpos-christie/user"
	"gopkg.in/yaml.v3"
)

// GUI represents the program parameters to launch in gui the application
type GUI struct {
	Config              *config.Config      // Config to use correctly the program
	User                *user.User          // User param to use program
	ThemeName           string              // name of the theme for fyne theme (Dark or Light)
	Theme               themes.Theme        // Fyne theme for the application
	App                 fyne.App            // Fyne application
	Window              fyne.Window         // Fyne window
	Language            settings.Yaml       // Yaml struct with all language data
	Currency            binding.String      // Currency to display
	labelCurrency       *widget.Label       // Label linked to currency
	labelIncome         *widget.Label       // Label for income
	entryIncome         *widget.Entry       // Input Entry to set income
	labelStatus         *widget.Label       // Label for status
	radioStatus         *widget.RadioGroup  // Input Radio buttons to get status
	labelChildren       *widget.Label       // Label for children
	selectChildren      *widget.SelectEntry // Input Select to know how children
	buttonSave          *widget.Button      // Label for save button
	labelTax            *widget.Label       // Label for tax
	labelTaxValue       *widget.Label       // Label for tax value
	labelRemainder      *widget.Label       // Label for remainder
	labelRemainderValue *widget.Label       // Label for remainder value
	labelShares         *widget.Label       // Label for shares
	labelSharesValue    *widget.Label       // Label for shares value
	labelsTrancheTaxes  []*widget.Label     // List of tranches tax label
	labelsAbout         []binding.String    // List of label in about modal
	labelsTaxHeaders    []binding.String    // List of label for tax details headers
}

// Start Launch GUI application
func (gui GUI) Start() {
	gui.App = app.New()
	gui.Window = gui.App.NewWindow(config.APP_NAME)

	// Size and Position
	const WIDTH = 1100
	const HEIGHT = 540
	gui.Window.Resize(fyne.NewSize(WIDTH, HEIGHT))
	gui.Window.CenterOnScreen()

	// Set Theme
	var theme string = settings.GetTheme()
	gui.setTheme(theme)

	// Set Language
	var language string = settings.GetLanguage()
	gui.setLanguage(language)

	// Set Currency
	var currency string = settings.GetCurrency()
	gui.Currency = binding.NewString()
	gui.setCurrency(currency)

	// Set Icon
	var icon fyne.Resource = settings.GetIcon()
	gui.Window.SetIcon(icon)

	// Set menu
	var menu *fyne.MainMenu = gui.setMenu()
	gui.Window.SetMainMenu(menu)

	// Handle Events and widgets
	gui.setLayouts()
	gui.setEvents()

	gui.Window.ShowAndRun()
}

// SetTheme change Theme of the application
func (gui *GUI) setTheme(theme string) {
	log.Printf("Debug theme: %+v", theme) // TODO log debug to show change theme
	var t themes.Theme
	if theme == themes.DARK {
		t = themes.DarkTheme{}
	} else {
		t = themes.LightTheme{}
	}
	gui.App.Settings().SetTheme(t)
}

// SetLanguage change language of the application
func (gui *GUI) setLanguage(code string) {
	log.Printf("Debug languages: %+v", code) // TODO log debug to show change language
	var language settings.Yaml = settings.Yaml{Code: code}
	var languageFile string = fmt.Sprintf("%s/%s.yaml", config.LANGUAGES_PATH, language.Code)
	yamlFile, _ := ioutil.ReadFile(languageFile)

	err := yaml.Unmarshal(yamlFile, &gui.Language)
	if err != nil {
		log.Fatalf("Unmarshal language file %s: %v", languageFile, err)
	}

	log.Printf("Debug languages: %+v", language) // TODO log debug to show change language
}

// setCurrency change language of the application
func (gui *GUI) setCurrency(currency string) {
	log.Printf("Debug currency: %+v", currency) // TODO log debug to show change currency
	gui.Currency.Set(currency)
	gui.labelCurrency = widget.NewLabel(currency)
}

// setEvents Set the events/trigger of gui widgets
func (gui *GUI) setEvents() {
	gui.entryIncome.OnChanged = func(input string) {
		gui.calculate()
	}
	gui.radioStatus.OnChanged = func(input string) {
		gui.calculate()
	}
	gui.selectChildren.OnChanged = func(input string) {
		gui.calculate()
	}

}

// getIncome Get value of widget entry
func (gui *GUI) getIncome() int {
	intVal, err := strconv.Atoi(gui.entryIncome.Text)
	if err != nil {
		return 0
	}
	return intVal
}

// getStatus Get value of widget radioGroup
func (gui *GUI) getStatus() bool {
	return gui.radioStatus.Selected == "Couple"
}

// getChildren get value of widget select
func (gui *GUI) getChildren() int {
	children, err := strconv.Atoi(gui.selectChildren.Entry.Text)
	if err != nil {
		return 0
	}
	return children
}

// reload Refresh widget who needed specially when language changed
func (gui *GUI) Reload() {

	// Reload about content
	for i, text := range gui.Language.GetAbouts() {
		gui.labelsAbout[i].Set(text)
	}

	// Reload header tax details
	for i, text := range gui.Language.GetTaxHeaders() {
		gui.labelsTaxHeaders[i].Set(text)
	}
}

// calculate Get values of gui to calculate tax
func (gui *GUI) calculate() {
	gui.User.Income = gui.getIncome()
	gui.User.IsInCouple = gui.getStatus()
	gui.User.Children = gui.getChildren()

	result := tax.CalculateTax(gui.User, gui.Config)
	log.Printf("Result - %#v ", result) // TODO log debug

	var taxValue string = utils.ConvertInt64ToString(int64(result.Tax))
	var remainderValue string = utils.ConvertInt64ToString(int64(result.Remainder))
	var shareValue string = utils.ConvertInt64ToString(int64(result.Shares))

	// Set data in tax layout
	gui.labelTaxValue.SetText(taxValue)
	gui.labelRemainderValue.SetText(remainderValue)
	gui.labelSharesValue.SetText(shareValue)

	// Set Tax details

	var trancheNumber int = 5 // TOODO a configurer via une functioin ou en attriibut de la gui
	for i := 0; i < trancheNumber; i++ {
		var taxTranche string = utils.ConvertIntToString(int(result.TaxTranches[i].Tax))
		gui.labelsTrancheTaxes[i].SetText(taxTranche + " " + "€") // TODO check devise
	}
}

// createMenu create mainMenu for window
func (gui *GUI) setMenu() *fyne.MainMenu {
	return fyne.NewMainMenu(
		gui.createFileMenu(),
		gui.createHelpMenu(),
	)
}

// createFileMenu create file item in toolbar to handle app settings
func (gui *GUI) createFileMenu() *fyne.Menu {
	fileMenu := fyne.NewMenu(gui.Language.File,
		fyne.NewMenuItem(gui.Language.Settings, func() {
			dialog.ShowCustom(gui.Language.Settings, gui.Language.Close,
				container.NewVBox(
					gui.createSelectTheme(),
					widget.NewSeparator(),
					gui.createSelectLanguage(),
					widget.NewSeparator(),
					gui.createSelectCurrency(),
				), gui.Window)
		}),
		fyne.NewMenuItem(gui.Language.Quit, func() { gui.App.Quit() }),
	)
	return fileMenu
}

// createSelectTheme create select to change theme
func (gui *GUI) createSelectTheme() *fyne.Container {
	selectTheme := widget.NewSelect(gui.Language.GetThemes(), func(val string) {
		gui.setTheme(val)
		// TODO save data in .settings
	})
	selectTheme.SetSelected(gui.ThemeName)
	return container.NewHBox(
		widget.NewLabel(gui.Language.ThemeCode),
		selectTheme,
	)
}

// createSelectLanguage create select to change language
func (gui *GUI) createSelectLanguage() *fyne.Container {
	selectLanguage := widget.NewSelect(gui.Language.GetLanguages(), nil)
	selectLanguage.OnChanged = func(s string) {
		index := selectLanguage.SelectedIndex()
		var getLanguage = func() string {
			switch index {
			case 0:
				return settings.ENGLISH
			case 1:
				return settings.FRENCH
			}
			// TODO error log
			return settings.FRENCH
		}

		language := getLanguage()
		gui.setLanguage(language)

		gui.Reload()
		// TODO save data in .settings
	}
	return container.NewHBox(
		widget.NewLabel(gui.Language.LanguageCode),
		selectLanguage,
	)
}

// createSelectCurrency create select to change currency
func (gui *GUI) createSelectCurrency() *fyne.Container {
	selectCurrency := widget.NewSelect(settings.GetCurrencies(), func(currency string) {
		gui.setCurrency(currency)
		gui.Reload()
		// TODO save data in .settings
	})
	return container.NewHBox(
		widget.NewLabel(gui.Language.Currency),
		selectCurrency,
	)
}

// createHelpMenu create help item in toolbar to show about app
func (gui *GUI) createHelpMenu() *fyne.Menu {
	url, _ := url.Parse(config.APP_LINK)
	for _, text := range gui.Language.GetAbouts() {
		var bindString binding.String = binding.NewString()
		bindString.Set(text)
		gui.labelsAbout = append(gui.labelsAbout, bindString)
	}

	// Setup layouts with data
	firstLine := container.NewHBox(
		widget.NewLabelWithData(gui.labelsAbout[0]),
		widget.NewLabel(config.APP_NAME),
		widget.NewLabelWithData(gui.labelsAbout[1]),
	)
	secondLine := container.NewHBox(
		widget.NewLabelWithData(gui.labelsAbout[2]),
		widget.NewHyperlink("GitHub Project", url),
		widget.NewLabelWithData(gui.labelsAbout[3]),
	)
	thirdLine := widget.NewLabelWithData(gui.labelsAbout[4])
	fourthLine := container.NewHBox(
		widget.NewLabel("Version:"),
		canvas.NewText(fmt.Sprintf("v%s", config.APP_VERSION), color.NRGBA{R: 218, G: 20, B: 51, A: 255}),
	)
	fifthLine := widget.NewLabel(fmt.Sprintf("%s: %s", gui.Language.Author, config.APP_AUTHOR))

	helpMenu := fyne.NewMenu(gui.Language.Help,
		fyne.NewMenuItem(gui.Language.About, func() {
			dialog.ShowCustom(gui.Language.About, gui.Language.Close,
				container.NewVBox(
					firstLine,
					secondLine,
					thirdLine,
					fourthLine,
					fifthLine,
				), gui.Window)
		}))
	return helpMenu
}
