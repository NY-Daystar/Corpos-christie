// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package core define the mode of the program console or gui
package core

import (
	"fmt"
	"image/color"
	"log"
	"net/url"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/core/themes"
	"github.com/LucasNoga/corpos-christie/tax"
	"github.com/LucasNoga/corpos-christie/user"
)

// GUIMode represents the program parameters to launch in console mode the application
type GUIMode struct {
	config         *config.Config // Config to use correctly the program
	user           *user.User     // User param to use program
	entryIncome    *widget.Entry
	radioStatus    *widget.RadioGroup
	selectChildren *widget.Select
}

// start launch core program in GUI Mode
func (gui GUIMode) start() {
	app := app.New()

	window := app.NewWindow("Corpos-Christie")
	app.Settings().SetTheme(themes.LightTheme{})
	// Set Icon
	r, _ := fyne.LoadResourceFromPath("./assets/logo.ico")
	window.SetIcon(r)

	// Size and Position
	window.Resize(fyne.NewSize(760, 480))
	window.CenterOnScreen()

	window.SetMainMenu(gui.setMenu(app, window))

	gui.setEntryIncome()
	gui.setRadioStatus()
	gui.setSelectChildren()

	gui.setEvents()
	// Layout income
	labelIncome := widget.NewLabel("Enter your income")
	incomeLayout := container.New(layout.NewFormLayout(), labelIncome, gui.entryIncome)

	// Layout status
	labelStatus := widget.NewLabel("Marital status")
	statusLayout := container.NewHBox(labelStatus, container.New(layout.NewVBoxLayout(), gui.radioStatus))

	// Layout children
	labelChildren := widget.NewLabel("Children ? ")

	childrenLayout := container.NewHBox(labelChildren, container.New(layout.NewVBoxLayout(), gui.selectChildren))

	button := widget.NewButton("Calculate Tax", func() {
		result := gui.calculate()
		log.Printf("Result - %#v ", result)
	})
	launcherLayout := container.NewHBox(button)

	form := container.New(layout.NewVBoxLayout(), incomeLayout, statusLayout, childrenLayout, launcherLayout)

	content := container.New(layout.NewGridLayout(2), form)

	window.SetContent(content)
	window.ShowAndRun()
}

// SetMenu create mainMenu for window
// TODO gerer la config
func (g *GUIMode) setMenu(app fyne.App, window fyne.Window) *fyne.MainMenu {
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Settings", func() {
			dialog.ShowCustom("Settings", "Close", container.NewVBox(
				container.NewHBox(

					widget.NewLabel("Theme"),
					widget.NewSelect([]string{"Dark", "Light"}, func(s string) {
						log.Printf("Theme %s", s)
						// TODO change theme
					}),
				),
				widget.NewSeparator(),
				container.NewHBox(
					widget.NewLabel("Languages"),
					widget.NewSelect([]string{"FR", "EN"}, func(s string) {
						log.Printf("Languages %s", s)
						// TODO change language + refresh
					}),
				),
			), window)
		}),
		fyne.NewMenuItem("Quit", func() { app.Quit() }),
	)

	url, _ := url.Parse(config.APP_LINK)

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel(fmt.Sprintf("Welcome to %s, a Desktop app to calculate your tax in France.", config.APP_NAME)),
				container.NewHBox(
					widget.NewLabel("This"),
					widget.NewHyperlink("GitHub Project", url),
					widget.NewLabel("is open-source."),
				),
				widget.NewLabel("Developped in Go with Fyne."),
				container.NewHBox(
					widget.NewLabel("Version:"),
					canvas.NewText(fmt.Sprintf("v%s", config.APP_VERSION), color.NRGBA{R: 218, G: 20, B: 51, A: 255}),
				),
				widget.NewLabel(fmt.Sprintf("Author: %s", config.APP_AUTHOR)),
			), window)
		}))
	return fyne.NewMainMenu(fileMenu, helpMenu)
}

// setEntryIncome create widget entry for income
func (gui *GUIMode) setEntryIncome() {
	gui.entryIncome = widget.NewEntry()
	gui.entryIncome.SetPlaceHolder("30000")
}

// setRadioStatus create widget radioGroup for marital status
func (gui *GUIMode) setRadioStatus() {
	gui.radioStatus = widget.NewRadioGroup([]string{"Single", "Couple"}, nil)
	gui.radioStatus.SetSelected("Single")
	gui.radioStatus.Horizontal = true
}

// setComboChildren create widget select for children
func (gui *GUIMode) setSelectChildren() {
	gui.selectChildren = widget.NewSelect([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, nil)
	gui.selectChildren.SetSelectedIndex(0)
}

// setEvents set the events/trigger of gui widgets
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

// getIncome get value of widget entry
func (gui *GUIMode) getIncome() int {
	intVal, err := strconv.Atoi(gui.entryIncome.Text)
	if err != nil {
		return 0
	}
	return intVal
}

// getStatus get value of widget radioGroup
func (gui *GUIMode) getStatus() bool {
	return gui.radioStatus.Selected == "Couple"
}

// getChildren get value of widget select
func (gui *GUIMode) getChildren() int {
	children, err := strconv.Atoi(gui.selectChildren.Selected)
	if err != nil {
		return 0
	}
	return children
}

// calculate get values of gui to calculate tax
func (gui *GUIMode) calculate() tax.Result {
	gui.user.Income = gui.getIncome()
	gui.user.IsInCouple = gui.getStatus()
	gui.user.Children = gui.getChildren()
	result := tax.CalculateTax(gui.user, gui.config)
	log.Printf("Result - %#v ", result)
	return result
	// TODO insted of return set a new widget to show values in EAST part of window
}
