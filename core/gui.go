// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package core define the mode of the program console or gui
package core

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"net/url"
	"reflect"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/core/themes"
	"github.com/LucasNoga/corpos-christie/lib/utils"
	"github.com/LucasNoga/corpos-christie/tax"
	"github.com/LucasNoga/corpos-christie/user"
	"gopkg.in/yaml.v3"
)

// GUIMode represents the program parameters to launch in console mode the application
type GUIMode struct {
	config              *config.Config      // Config to use correctly the program
	user                *user.User          // User param to use program
	themeName           string              // name of the theme for fyne theme (Dark or Light)
	theme               themes.Theme        // Fyne theme for the application
	app                 fyne.App            // Fyne application
	window              fyne.Window         // Fyne window
	language            Yaml                // Yaml struct with all language data
	labelIncome         *widget.Label       // Label for income
	entryIncome         *widget.Entry       // Input Entry to set income
	labelStatus         *widget.Label       // Label for status
	radioStatus         *widget.RadioGroup  // Input Radio buttons to get status
	labelChildren       *widget.Label       // Label for children
	selectChildren      *widget.SelectEntry // Input Select to know how children
	labelTax            *widget.Label       // Label for tax
	labelTaxValue       *widget.Label       // Label for tax value
	labelRemainder      *widget.Label       // Label for remainder
	labelRemainderValue *widget.Label       // Label for remainder value
	labelShares         *widget.Label       // Label for shares
	labelSharesValue    *widget.Label       // Label for shares value
}

// LanguageYaml Yaml struct to get language data
type Yaml struct {
	code         string       // code of the language (fr, en, etc...)
	Theme        ThemeYaml    `yaml:"themes"`
	Languages    LanguageYaml `yaml:"languages"`
	File         string       `yaml:"file"`
	Settings     string       `yaml:"settings"`
	Income       string       `yaml:"income"`
	Status       string       `yaml:"status"`
	Children     string       `yaml:"children"`
	Tax          string       `yaml:"tax"`
	Remainder    string       `yaml:"remainder"`
	Share        string       `yaml:"share"`
	SaveTax      string       `yaml:"save_tax"`
	ThemeCode    string       `yaml:"theme"`
	LanguageCode string       `yaml:"language"`
	Help         string       `yaml:"help"`
	About        string       `yaml:"about"`
	Author       string       `yaml:"author"`
	Close        string       `yaml:"close"`
	Quit         string       `yaml:"quit"`
}

// ThemeYml Yaml struct for theme's app
type ThemeYaml struct {
	Dark  string `yaml:"dark"`
	Light string `yaml:"light"`
}

// Languages Yaml struct for theme's app
type LanguageYaml struct {
	English string `yaml:"english"`
	French  string `yaml:"french"`
}

// getThemes Parse ThemeYaml struct to get value of each field
func (t ThemeYaml) getThemes() []string {
	v := reflect.ValueOf(t)
	themes := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		themes[i] = v.Field(i).String()
	}
	return themes
}

// getLanguages Parse LanguagesXml struct to get value of each field
func (l LanguageYaml) getLanguages() []string {
	v := reflect.ValueOf(l)
	languages := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		languages[i] = v.Field(i).String()
	}
	return languages
}

// start Launch core program in GUI Mode
func (gui GUIMode) start() {
	gui.app = app.New()
	gui.window = gui.app.NewWindow(config.APP_NAME)

	// Set Theme
	var theme string = gui.getTheme()
	gui.setTheme(theme)

	// Set Language
	var language string = gui.getLanguage()
	gui.setLanguage(language)

	// Set Icon
	var icon fyne.Resource = gui.getIcon()
	gui.window.SetIcon(icon)

	// Size and Position
	gui.window.Resize(fyne.NewSize(760, 480))
	gui.window.CenterOnScreen()

	gui.window.SetMainMenu(gui.setMenu())

	gui.setEntryIncome()
	gui.setRadioStatus()
	gui.setSelectChildren()

	// Handle Events
	gui.setEvents()

	// Layout income
	gui.labelIncome = widget.NewLabel(gui.language.Income)
	incomeLayout := container.New(layout.NewFormLayout(), gui.labelIncome, gui.entryIncome)

	// Layout status
	gui.labelStatus = widget.NewLabel(gui.language.Status)
	statusLayout := container.NewHBox(gui.labelStatus, container.New(layout.NewVBoxLayout(), gui.radioStatus))

	// Layout children
	gui.labelChildren = widget.NewLabel(gui.language.Children)
	childrenLayout := container.NewHBox(gui.labelChildren, container.New(layout.NewVBoxLayout(), gui.selectChildren))

	// Layout tax results
	gui.labelTax = widget.NewLabel(gui.language.Tax)
	gui.labelTaxValue = widget.NewLabel("")
	gui.labelRemainder = widget.NewLabel(gui.language.Remainder)
	gui.labelRemainderValue = widget.NewLabel("")
	gui.labelShares = widget.NewLabel(gui.language.Share)
	gui.labelSharesValue = widget.NewLabel("")
	taxResultLayout := container.New(layout.NewGridLayout(3),
		gui.labelTax,
		gui.labelTaxValue,
		widget.NewLabel("€"), // TODO ajouter les labels €,$,£ sur la partie droite

		gui.labelShares,
		gui.labelSharesValue,
		widget.NewLabel("€"), // TODO ajouter les labels €,$,£ sur la partie droite

		gui.labelRemainder,
		gui.labelRemainderValue,
		widget.NewLabel("€"), // TODO ajouter les labels €,$,£ sur la partie droite
	)
	// Layout button
	button := widget.NewButton(gui.language.SaveTax, func() {
		gui.calculate()
		log.Printf("Save Tax") // TODO debug // TODO language
	})
	launcherLayout := container.NewHBox(button)

	formLayout := container.New(layout.NewVBoxLayout(), incomeLayout, statusLayout, childrenLayout, launcherLayout)
	taxLayout := container.New(layout.NewVBoxLayout(), taxResultLayout, taxDetailLayout)

	// separator := widget.NewSeparator()

	content := container.New(layout.NewGridLayout(2), formLayout, taxLayout)

	gui.window.SetContent(content)
	gui.window.ShowAndRun()
}

// SetMenu Create mainMenu for window
func (gui *GUIMode) setMenu() *fyne.MainMenu {

	selectTheme := widget.NewSelect(gui.language.Theme.getThemes(), func(val string) {
		gui.setTheme(val)
		// TODO save data in .settings
	})
	selectTheme.SetSelected(gui.themeName)

	selectLanguage := widget.NewSelect(gui.language.Languages.getLanguages(), nil)
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

		gui.setLanguage(language)
		gui.reload()
		// TODO save data in .settings

	}

	fileMenu := fyne.NewMenu(gui.language.File,
		fyne.NewMenuItem(gui.language.Settings, func() {
			dialog.ShowCustom(gui.language.Settings, gui.language.Close, container.NewVBox(
				container.NewHBox(
					widget.NewLabel(gui.language.ThemeCode),
					selectTheme,
				),
				widget.NewSeparator(),
				container.NewHBox(
					widget.NewLabel(gui.language.LanguageCode),
					selectLanguage,
				),
			), gui.window)
		}),
		fyne.NewMenuItem(gui.language.Quit, func() { gui.app.Quit() }),
	)

	url, _ := url.Parse(config.APP_LINK)

	helpMenu := fyne.NewMenu(gui.language.Help,
		fyne.NewMenuItem(gui.language.About, func() {
			dialog.ShowCustom(gui.language.About, gui.language.Close, container.NewVBox(
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
				widget.NewLabel(fmt.Sprintf("%s: %s", gui.language.Author, config.APP_AUTHOR)),
			), gui.window)
		}))
	return fyne.NewMainMenu(fileMenu, helpMenu)
}

// setEntryIncome Create widget entry for income
func (gui *GUIMode) setEntryIncome() {
	gui.entryIncome = widget.NewEntry()
	gui.entryIncome.SetPlaceHolder("30000")
	gui.entryIncome.Validator = validation.NewRegexp("^[0-9]{1,}$", "Not a number") // TODO language
}

// setRadioStatus Create widget radioGroup for marital status
func (gui *GUIMode) setRadioStatus() {
	gui.radioStatus = widget.NewRadioGroup([]string{"Single", "Couple"}, nil) // TODO language
	gui.radioStatus.SetSelected("Single")
	gui.radioStatus.Horizontal = true
}

// setComboChildren Create widget select for children
func (gui *GUIMode) setSelectChildren() {
	gui.selectChildren = widget.NewSelectEntry([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
	gui.selectChildren.SetText("0")
	gui.selectChildren.Validator = validation.NewRegexp("^[0-9]{1,}$", "Not a number") // TODO language
}

// setEvents Set the events/trigger of gui widgets
func (gui *GUIMode) setEvents() {
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
func (gui *GUIMode) getIncome() int {
	intVal, err := strconv.Atoi(gui.entryIncome.Text)
	if err != nil {
		return 0
	}
	return intVal
}

// getStatus Get value of widget radioGroup
func (gui *GUIMode) getStatus() bool {
	return gui.radioStatus.Selected == "Couple" // TODO language ou mettre un id
}

// getChildren get value of widget select
func (gui *GUIMode) getChildren() int {
	children, err := strconv.Atoi(gui.selectChildren.Entry.Text)
	if err != nil {
		return 0
	}
	return children
}

// GetIcon Load icon file to show in window
func (gui *GUIMode) getIcon() fyne.Resource {
	var iconName string = "logo.ico"
	var iconPath string = fmt.Sprintf("%s/%s", config.ASSETS_PATH, iconName)
	icon, _ := fyne.LoadResourceFromPath(iconPath)
	// TODO log debug to show icon loaded
	return icon
}

// getTheme Get value of last theme selected
func (gui *GUIMode) getTheme() string {
	// TODO get value from .setting file
	// TODO log debug to show change theme
	return "Dark"
}

// setTheme Change Theme of the application
func (gui *GUIMode) setTheme(theme string) {
	// TODO log debug to show change theme
	gui.themeName = theme
	if gui.themeName == themes.DARK {
		gui.theme = themes.DarkTheme{}
	} else {
		gui.theme = themes.LightTheme{}
	}
	gui.app.Settings().SetTheme(gui.theme)
}

// getLanguage Get value of last language selected (fr, en)
func (gui *GUIMode) getLanguage() string {
	// TODO get value from .setting file
	// TODO log debug to show change language
	return "en"
}

// setLanguage Change language of the application
func (gui *GUIMode) setLanguage(code string) {
	var language Yaml = Yaml{code: code}
	var languageFile string = fmt.Sprintf("%s/%s.yaml", config.LANGUAGES_PATH, language.code)
	yamlFile, _ := ioutil.ReadFile(languageFile)

	err := yaml.Unmarshal(yamlFile, &gui.language)
	if err != nil {
		log.Fatalf("Unmarshal language file %s: %v", languageFile, err)
	}

	log.Printf("Debug languages: %+v", gui.language) // TODO log debug to show change theme
}

// reload Refresh widget who needed specially when language changed
func (gui *GUIMode) reload() {
	gui.labelIncome.SetText(gui.language.Income)
	gui.labelStatus.SetText(gui.language.Status)
	gui.labelChildren.SetText(gui.language.Children)
	gui.labelTax.SetText(gui.language.Tax)
	gui.labelRemainder.SetText(gui.language.Remainder)
	gui.labelShares.SetText(gui.language.Share)
}

// calculate Get values of gui to calculate tax
func (gui *GUIMode) calculate() {
	gui.user.Income = gui.getIncome()
	gui.user.IsInCouple = gui.getStatus()
	gui.user.Children = gui.getChildren()
	result := tax.CalculateTax(gui.user, gui.config)
	log.Printf("Result - %#v ", result)

	var taxValue string = utils.ConvertInt64ToString(int64(result.Tax))
	var remainderValue string = utils.ConvertInt64ToString(int64(result.Remainder))
	var shareValue string = utils.ConvertInt64ToString(int64(result.Shares))

	// Set data in tax layout
	gui.labelTaxValue.SetText(taxValue)
	gui.labelRemainderValue.SetText(remainderValue)
	gui.labelSharesValue.SetText(shareValue)

}
