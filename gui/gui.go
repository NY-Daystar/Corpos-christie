// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package gui defines component and script to launch gui application
package gui

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"math"
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
	"github.com/LucasNoga/corpos-christie/tax"
	"github.com/LucasNoga/corpos-christie/user"
	"github.com/LucasNoga/corpos-christie/utils"
	"gopkg.in/yaml.v3"
)

// GUI represents the program parameters to launch in gui the application
type GUI struct {
	Config *config.Config // Config to use correctly the program
	User   *user.User     // User param to use program
	App    fyne.App       // Fyne application
	Window fyne.Window    // Fyne window

	// Settings
	ThemeName string         // name of the theme for fyne theme (Dark or Light)
	Theme     themes.Theme   // Fyne theme for the application
	Language  settings.Yaml  // Yaml struct with all language data
	Currency  binding.String // Currency to display

	// Widgets
	entryIncome    *widget.Entry       // Input Entry to set income
	radioStatus    *widget.RadioGroup  // Input Radio buttons to get status
	selectChildren *widget.SelectEntry // Input Select to know how children

	buttonSave *widget.Button // Label for save button

	// Bindings
	Tax                binding.String     // Bind for tax value
	Remainder          binding.String     // Bind for remainder value
	Shares             binding.String     // Bind for shares value
	labelShares        binding.String     // Bind for shares label
	labelIncome        binding.String     // Bind for income label
	labelStatus        binding.String     // Bind for status label
	labelChildren      binding.String     // Bind for children label
	labelTax           binding.String     // Bind for tax label
	labelRemainder     binding.String     // Bind for remainder label
	labelsAbout        binding.StringList // List of label in about modal
	labelsTaxHeaders   binding.StringList // List of label for tax details headers
	labelsMinTranche   binding.StringList // List of labels for min tranche in grid
	labelsMaxTranche   binding.StringList // List of labels for max tranche in grid
	labelsTrancheTaxes binding.StringList // List of tranches tax label
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
	var theme string = settings.GetDefaultTheme()
	gui.setTheme(theme)

	// Set Language
	var language string = settings.GetDefaultLanguage()
	gui.setLanguage(language)

	// Set Currency
	var currency string = settings.GetDefaultCurrency()
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
	if theme == settings.DARK {
		t = themes.DarkTheme{}
	} else {
		t = themes.LightTheme{}
	}
	gui.ThemeName = theme
	gui.App.Settings().SetTheme(t)
}

// SetLanguage change language of the application
func (gui *GUI) setLanguage(code string) {
	log.Printf("Debug code language: %+v", code) // TODO log debug to show change language

	var languageFile string = fmt.Sprintf("%s/%s.yaml", config.LANGUAGES_PATH, code)
	yamlFile, _ := ioutil.ReadFile(languageFile)

	var language settings.Yaml = settings.Yaml{Code: code}
	err := yaml.Unmarshal(yamlFile, &gui.Language)
	gui.Language.Code = code

	if err != nil {
		log.Fatalf("Unmarshal language file %s: %v", languageFile, err)
	}

	log.Printf("Debug languages: %+v", language) // TODO log debug to show change language
}

// setCurrency change language of the application
func (gui *GUI) setCurrency(currency string) {
	log.Printf("Debug currency: %+v", currency) // TODO log debug to show change currency
	gui.Currency.Set(currency)
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
	// Simple data bind
	gui.labelIncome.Set(gui.Language.Income)
	gui.labelStatus.Set(gui.Language.Status)
	gui.labelChildren.Set(gui.Language.Children)
	gui.labelTax.Set(gui.Language.Tax)
	gui.labelRemainder.Set(gui.Language.Remainder)
	gui.labelShares.Set(gui.Language.Share)

	// Handle widget
	gui.buttonSave.SetText(gui.Language.Save)

	// Reload about content
	gui.labelsAbout.Set(gui.Language.GetAbouts())

	// Reload header tax details
	gui.labelsTaxHeaders.Set(gui.Language.GetTaxHeaders())

	// Reload grid header
	currency, _ := gui.Currency.Get()
	gui.labelsTrancheTaxes.Set(*createTrancheTaxesLabels(gui.labelsTrancheTaxes.Length(), currency))

	// Reload grid min tranches
	var minList []string
	for index := 0; index < gui.labelsMinTranche.Length(); index++ {
		var min string = utils.ConvertIntToString(gui.Config.Tax.Tranches[index].Min) + " " + currency
		minList = append(minList, min)
	}
	gui.labelsMinTranche.Set(minList)

	// Reload grid max tranches
	var maxList []string
	for index := 0; index < gui.labelsMaxTranche.Length(); index++ {
		var max string = utils.ConvertIntToString(gui.Config.Tax.Tranches[index].Max) + " " + currency
		if gui.Config.Tax.Tranches[index].Max == math.MaxInt64 {
			max = "-"
		}
		maxList = append(maxList, max)
	}
	gui.labelsMaxTranche.Set(maxList)
}

// calculate Get values of gui to calculate tax
func (gui *GUI) calculate() {
	gui.User.Income = gui.getIncome()
	gui.User.IsInCouple = gui.getStatus()
	gui.User.Children = gui.getChildren()

	result := tax.CalculateTax(gui.User, gui.Config)
	log.Printf("Result - %#v ", result) // TODO log debug

	var tax string = utils.ConvertInt64ToString(int64(result.Tax))
	var remainder string = utils.ConvertInt64ToString(int64(result.Remainder))
	var shares string = utils.ConvertInt64ToString(int64(result.Shares))

	// Set data in tax layout
	gui.Tax.Set(tax)
	gui.Remainder.Set(remainder)
	gui.Shares.Set(shares)

	// Set Tax details
	currency, _ := gui.Currency.Get()
	for index := 0; index < gui.labelsTrancheTaxes.Length(); index++ {
		var taxTranche string = utils.ConvertIntToString(int(result.TaxTranches[index].Tax))
		gui.labelsTrancheTaxes.SetValue(index, taxTranche+" "+currency)
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
	selectTheme.SetSelectedIndex(getThemeIndex(gui.Language.Code, gui.ThemeName))
	return container.NewHBox(
		widget.NewLabel(gui.Language.ThemeCode),
		selectTheme,
	)
}

// createSelectLanguage create select to change language
func (gui *GUI) createSelectLanguage() *fyne.Container {
	selectLanguage := widget.NewSelect(gui.Language.GetLanguages(), nil)
	selectLanguage.SetSelectedIndex(getLanguageIndex(gui.Language.Code))
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
	currency, _ := gui.Currency.Get()
	selectCurrency.SetSelected(currency)
	return container.NewHBox(
		widget.NewLabel(gui.Language.Currency),
		selectCurrency,
	)
}

// createHelpMenu create help item in toolbar to show about app
func (gui *GUI) createHelpMenu() *fyne.Menu {
	url, _ := url.Parse(config.APP_LINK)

	gui.labelsAbout = binding.NewStringList()
	gui.labelsAbout.Set(gui.Language.GetAbouts())
	var labels []binding.DataItem
	for index := range gui.Language.GetAbouts() {
		about, _ := gui.labelsAbout.GetItem(index)
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

// getThemeIndex get index to selectTheme in settings from language of the app
func getThemeIndex(langue, theme string) int {
	var themes map[string][]string = map[string][]string{
		"en": {"Dark", "Light"},
		"fr": {"Sombre", "Clair"},
	}

	var l []string = themes[langue]
	for index, v := range l {
		if v == theme {
			return index
		}
	}
	return -1
}

// getLanguageIndex get index to selectLanguage in settings from language of the app
func getLanguageIndex(langue string) int {
	switch langue {
	case settings.ENGLISH:
		return 0
	case settings.FRENCH:
		return 1
	default:
		return 0
	}
}
