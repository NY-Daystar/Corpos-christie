// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package core define the mode of the program console or gui
package core

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/LucasNoga/corpos-christie/config"
	"github.com/LucasNoga/corpos-christie/tax"
	"github.com/LucasNoga/corpos-christie/user"
)

// GUIMode represents the program parameters to launch in console mode the application
type GUIMode struct {
	config *config.Config // Config to use correctly the program
	user   *user.User     // User param to use program
}

// start launch core program in GUI Mode
func (mode GUIMode) start() {
	app := app.New()

	window := app.NewWindow("Corpos-Christie")

	// Set Icon
	r, _ := fyne.LoadResourceFromPath("./assets/logo.ico")
	window.SetIcon(r)

	// Size and Position
	window.Resize(fyne.NewSize(760, 480))
	window.CenterOnScreen()

	// Main menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { app.Quit() }),
	)

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel("Welcome to Gopher, a simple Desktop app created in Go with Fyne."),
				widget.NewLabel("Version: v0.1"),
				widget.NewLabel("Author: Aurélie Vache"),
			), window)
		}))
	window.SetMainMenu(fyne.NewMainMenu(
		fileMenu,
		helpMenu,
	))

	// Body
	labelIncome := widget.NewLabel("Enter your income")
	entryIncome := widget.NewEntry()
	entryIncome.SetPlaceHolder("Enter Income...")
	incomeLayout := container.New(layout.NewFormLayout(), labelIncome, entryIncome)

	labelStatus := widget.NewLabel("Marital status")
	radioStatus := widget.NewRadioGroup([]string{"Single", "Couple"}, func(value string) {})
	radioStatus.Horizontal = true
	statusLayout := container.NewHBox(labelStatus, container.New(layout.NewVBoxLayout(), radioStatus))

	labelChildren := widget.NewLabel("Children ? ")
	comboChildren := widget.NewSelect([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, func(value string) {})
	comboChildren.SetSelectedIndex(0)
	childrenLayout := container.NewHBox(labelChildren, container.New(layout.NewVBoxLayout(), comboChildren))

	button := widget.NewButton("Calculate Tax", func() {
		intVal, _ := strconv.Atoi(entryIncome.Text) // TODO Handle Error
		mode.user.Income = intVal
		if radioStatus.Selected == "Couple" {
			mode.user.IsInCouple = true
		}
		children, _ := strconv.Atoi(comboChildren.Selected) // TODO Handle Error
		mode.user.Children = children
		result := tax.CalculateTax(mode.user, mode.config)
		log.Printf("Result - %#v ", result)
	})
	launcherLayout := container.NewHBox(button)

	form := container.New(layout.NewVBoxLayout(), incomeLayout, statusLayout, childrenLayout, launcherLayout)

	content := container.New(layout.NewGridLayout(2), form)

	window.SetContent(content)
	window.ShowAndRun()
}
