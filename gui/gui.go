// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package gui defines component and script to launch gui application
package gui

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/gui/settings"
	"github.com/LucasNoga/corpos-christie/gui/themes"
	"github.com/LucasNoga/corpos-christie/gui/widgets"
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

// TODO gerer les selecteurs de langues et de theme en fonction de la langue

// TODO faire un gui_test.go
// - Tester l'icon voir si on a accès au fichier et si il existe
// - Tester les languages voir si on a des valeurs ou pas
// - Tester les themes voir si on a des valeurs ou pas

// Start Launch GUI application
func (gui GUI) Start() {
	gui.App = app.New()
	gui.Window = gui.App.NewWindow(config.APP_NAME)

	// Set Theme
	var theme string = settings.GetTheme()
	gui.SetTheme(theme)

	// Set Language
	var language string = settings.GetLanguage()
	gui.SetLanguage(language)

	// Set Icon
	var icon fyne.Resource = settings.GetIcon()
	gui.Window.SetIcon(icon)

	// Size and Position
	gui.Window.Resize(fyne.NewSize(760, 480))
	gui.Window.CenterOnScreen()

	// Set menu
	var menu *fyne.MainMenu = gui.SetMenu()
	gui.Window.SetMainMenu(menu)

	gui.entryIncome = widgets.CreateIncomeEntry()
	gui.radioStatus = widgets.CreateStatusRadio()
	gui.selectChildren = widgets.CreateChildrenSelect()

	// Handle Events
	gui.setEvents()

	// Layout income
	gui.labelIncome = widget.NewLabel(gui.Language.Income)
	incomeLayout := container.New(layout.NewFormLayout(), gui.labelIncome, gui.entryIncome)

	// Layout status
	gui.labelStatus = widget.NewLabel(gui.Language.Status)
	statusLayout := container.NewHBox(gui.labelStatus, container.New(layout.NewVBoxLayout(), gui.radioStatus))

	// Layout children
	gui.labelChildren = widget.NewLabel(gui.Language.Children)
	childrenLayout := container.NewHBox(gui.labelChildren, container.New(layout.NewVBoxLayout(), gui.selectChildren))

	// Layout tax results
	gui.labelTax = widget.NewLabel(gui.Language.Tax)
	gui.labelTaxValue = widget.NewLabel("")
	gui.labelRemainder = widget.NewLabel(gui.Language.Remainder)
	gui.labelRemainderValue = widget.NewLabel("")
	gui.labelShares = widget.NewLabel(gui.Language.Share)
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

	// TODO faire un separator

	// TODO faire un tableau
	t := widget.NewLabel("TABLE")
	taxDetailLayout := container.NewHBox(t)
	// +-----------+-----------+-----------------------+-----------+--------+
	// |  TRANCHE  |    MIN    |          MAX          |   RATE    |  TAX   |
	// +-----------+-----------+-----------------------+-----------+--------+
	// | Tranche 1 | 0 €       | 10225 €               | 0 %       | 0 €    |
	// | Tranche 2 | 10226 €   | 26070 €               | 11 %      | 635 €  |
	// | Tranche 3 | 26071 €   | 74545 €               | 30 %      | 0 €    |
	// | Tranche 4 | 74546 €   | 160336 €              | 41 %      | 0 €    |
	// | Tranche 5 | 160337 €  | 9223372036854775807 € | 45 %      | 0 €    |
	// +-----------+-----------+-----------------------+-----------+--------+
	// |  RESULT   | REMAINDER |        38412 €        | TOTAL TAX | 1588 € |
	// +-----------+-----------+-----------------------+-----------+--------+

	// Layout buttons
	button := widget.NewButton(gui.Language.SaveTax, func() {
		gui.calculate()
		log.Printf("Save Tax") // TODO debug // TODO language
	})
	// TODO faire un boutton reset
	launcherLayout := container.NewHBox(button)

	formLayout := container.New(layout.NewVBoxLayout(), incomeLayout, statusLayout, childrenLayout, launcherLayout)
	taxLayout := container.New(layout.NewVBoxLayout(), taxResultLayout, taxDetailLayout)

	// TODO make a line separator
	// TODO check canvas
	// separator := widget.NewSeparator()

	// content := container.New(layout.NewGridLayout(3), formLayout, separator, taxLayout)
	content := container.New(layout.NewGridLayout(2), formLayout, taxLayout)

	gui.Window.SetContent(content)
	gui.Window.ShowAndRun()
}

// SetTheme Change Theme of the application
func (gui *GUI) SetTheme(theme string) {
	// TODO log debug to show change theme
	var t themes.Theme
	if theme == themes.DARK {
		t = themes.DarkTheme{}
	} else {
		t = themes.LightTheme{}
	}
	gui.App.Settings().SetTheme(t)
}

// SetLanguage Change language of the application
// Returns Yaml structure with language label
func (gui *GUI) SetLanguage(code string) {
	var language settings.Yaml = settings.Yaml{Code: code}
	var languageFile string = fmt.Sprintf("%s/%s.yaml", config.LANGUAGES_PATH, language.Code)
	yamlFile, _ := ioutil.ReadFile(languageFile)

	err := yaml.Unmarshal(yamlFile, &gui.Language)
	if err != nil {
		log.Fatalf("Unmarshal language file %s: %v", languageFile, err)
	}

	log.Printf("Debug languages: %+v", language) // TODO log debug to show change theme

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
	return gui.radioStatus.Selected == "Couple" // TODO language ou mettre un id
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
	gui.labelIncome.SetText(gui.Language.Income)
	gui.labelStatus.SetText(gui.Language.Status)
	gui.labelChildren.SetText(gui.Language.Children)
	gui.labelTax.SetText(gui.Language.Tax)
	gui.labelRemainder.SetText(gui.Language.Remainder)
	gui.labelShares.SetText(gui.Language.Share)
}

// calculate Get values of gui to calculate tax
func (gui *GUI) calculate() {
	gui.User.Income = gui.getIncome()
	gui.User.IsInCouple = gui.getStatus()
	gui.User.Children = gui.getChildren()
	result := tax.CalculateTax(gui.User, gui.Config)
	log.Printf("Result - %#v ", result)

	var taxValue string = utils.ConvertInt64ToString(int64(result.Tax))
	var remainderValue string = utils.ConvertInt64ToString(int64(result.Remainder))
	var shareValue string = utils.ConvertInt64ToString(int64(result.Shares))

	// Set data in tax layout
	gui.labelTaxValue.SetText(taxValue)
	gui.labelRemainderValue.SetText(remainderValue)
	gui.labelSharesValue.SetText(shareValue)

}
