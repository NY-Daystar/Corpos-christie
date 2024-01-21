// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package gui defines component and script to launch gui application
package gui

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/widgets"
	"github.com/NY-Daystar/corpos-christie/utils"
)

// setLayouts Setup components/widget in the window
func (gui *GUI) setLayouts() {
	content := container.New(layout.NewGridLayout(2),
		gui.createLayoutForm(),
		gui.createLayoutTax(),
	)
	gui.Window.SetContent(content)
}

// createLayoutForm Setup left side of window
func (gui *GUI) createLayoutForm() *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		gui.createLayoutIncome(),
		gui.createLayoutStatus(),
		gui.createLayoutChildren(),
		// gui.createLayoutSave(), // TODO
	)
}

// createLayoutIncome Setup layouts and widget for income layout
func (gui *GUI) createLayoutIncome() *fyne.Container {
	gui.entryIncome = widgets.CreateIncomeEntry()
	gui.labelIncome = binding.BindString(&gui.Language.Income)
	return container.New(
		layout.NewFormLayout(),
		widget.NewLabelWithData(gui.labelIncome),
		gui.entryIncome,
	)
}

// createLayoutStatus Setup layouts and widget for income layout
func (gui *GUI) createLayoutStatus() *fyne.Container {
	gui.radioStatus = widgets.CreateStatusRadio()
	gui.labelStatus = binding.BindString(&gui.Language.Status)
	// gui.labelStatus = widget.NewLabel(gui.Language.Status)
	return container.NewHBox(
		widget.NewLabelWithData(gui.labelStatus),
		container.New(
			layout.NewVBoxLayout(),
			gui.radioStatus,
		),
	)

}

// createLayoutChildren Setup layouts and widget for income layout
func (gui *GUI) createLayoutChildren() *fyne.Container {
	gui.selectChildren = widgets.CreateChildrenSelect()
	gui.labelChildren = binding.BindString(&gui.Language.Children)
	return container.NewHBox(
		widget.NewLabelWithData(gui.labelChildren),
		container.New(
			layout.NewVBoxLayout(),
			gui.selectChildren,
		),
	)
}

// createLayoutSave Setup layouts and widget for save button layout
// func (gui *GUI) createLayoutSave() *fyne.Container {
// 	gui.buttonSave = widget.NewButton(gui.Language.Save, func() {
// 		gui.calculate()
// 		gui.Logger.Info("Save Taxes")
// 		// TODO Export taxes data in csv and/or pdf
// 	})
// 	return container.NewHBox(gui.buttonSave)
// }

// createLayoutTax Setup right side of window
func (gui *GUI) createLayoutTax() *fyne.Container {
	return container.New(
		layout.NewVBoxLayout(),
		gui.createLayoutTaxResult(),
		container.NewVBox(widget.NewLabel(""), widget.NewSeparator(), widget.NewLabel("")),
		gui.createLayoutTaxDetails(),
	)
}

// createLayoutTaxResult Setup right top side of window
func (gui *GUI) createLayoutTaxResult() *fyne.Container {
	gui.labelTax = binding.BindString(&gui.Language.Tax)
	gui.Tax = binding.NewString()

	gui.labelShares = binding.BindString(&gui.Language.Share)
	gui.Shares = binding.NewString()

	gui.labelRemainder = binding.BindString(&gui.Language.Remainder)
	gui.Remainder = binding.NewString()

	return container.New(layout.NewGridLayout(3),
		widget.NewLabelWithData(gui.labelTax),
		widget.NewLabelWithData(gui.Tax),
		widget.NewLabelWithData(gui.Currency),

		widget.NewLabelWithData(gui.labelShares),
		widget.NewLabelWithData(gui.Shares),
		widget.NewLabelWithData(gui.Currency),

		widget.NewLabelWithData(gui.labelRemainder),
		widget.NewLabelWithData(gui.Remainder),
		widget.NewLabelWithData(gui.Currency),
	)
}

// createLayoutTax Setup right bottom side of window
func (gui *GUI) createLayoutTaxDetails() *fyne.Container {
	var trancheNumber int = 5
	currency, _ := gui.Currency.Get()

	// Add header columns in grid
	grid := container.New(layout.NewGridLayout(trancheNumber))

	gui.labelsTaxHeaders = binding.NewStringList()
	for index, header := range gui.Language.GetTaxHeaders() {
		gui.labelsTaxHeaders.Append(header)
		h, _ := gui.labelsTaxHeaders.GetItem(index)
		grid.Add(widget.NewLabelWithData(h.(binding.String)))
	}

	// Setup binding for min, max and taxes columns
	gui.labelsMinTranche = binding.BindStringList(createMinTrancheLabels(currency, gui.Config.Tax.Tranches))
	gui.labelsMaxTranche = binding.BindStringList(createMaxTrancheLabels(currency, gui.Config.Tax.Tranches))
	gui.labelsTrancheTaxes = binding.BindStringList(createTrancheTaxesLabels(trancheNumber, currency))

	// Add Tranche rows in grid
	for index := 0; index < gui.labelsTrancheTaxes.Length(); index++ {
		minItem, _ := gui.labelsMinTranche.GetItem(index)
		maxItem, _ := gui.labelsMaxTranche.GetItem(index)
		taxItem, _ := gui.labelsTrancheTaxes.GetItem(index)
		var rate string = gui.Config.Tax.Tranches[index].Rate

		grid.Add(widget.NewLabel("Tranche " + utils.ConvertIntToString(index+1)))
		grid.Add(widget.NewLabelWithData(minItem.(binding.String)))
		grid.Add(widget.NewLabelWithData(maxItem.(binding.String)))
		grid.Add(widget.NewLabel(rate))
		grid.Add(widget.NewLabelWithData(taxItem.(binding.String)))
	}

	return container.New(
		layout.NewMaxLayout(),
		grid,
	)
}

// CreateTrancheLabels create widgets labels for tranche taxes value into an array
// Create number of tranche with currency value
// Returns Array of label widget in fyne object
func createTrancheTaxesLabels(number int, currency string) *[]string {
	var labels []string = make([]string, 0, number)

	for i := 1; i <= number; i++ {
		labels = append(labels, "0"+" "+currency)
	}
	return &labels
}

// createMinTrancheLabels create string from config.Tranche to create binding
// Returns Array string with min tranches value
func createMinTrancheLabels(currency string, tranches []config.Tranche) *[]string {
	var labels []string = make([]string, 0, len(tranches))

	for _, tranche := range tranches {
		var min string = utils.ConvertIntToString(tranche.Min) + " " + currency
		labels = append(labels, min)
	}

	return &labels
}

// createMaxTrancheLabels create string from config.Tranche to create binding
// Returns Array string with max tranches value
func createMaxTrancheLabels(currency string, tranches []config.Tranche) *[]string {
	var labels []string = make([]string, 0, len(tranches))

	for _, tranche := range tranches {
		var max = utils.ConvertIntToString(tranche.Max) + " " + currency
		if tranche.Max == math.MaxInt64 {
			max = "-"
		}
		labels = append(labels, max)
	}
	return &labels
}
