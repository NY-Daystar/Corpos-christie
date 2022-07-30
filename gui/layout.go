// Copyright 2016 The corpos-christie author
// Licensed under GPLv3.

// Package gui defines component and script to launch gui application
package gui

import (
	"image/color"
	"log"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/LucasNoga/corpos-christie/gui/widgets"
	"github.com/LucasNoga/corpos-christie/lib/utils"
)

// setLayouts Setup components/widget in the window
func (gui *GUI) setLayouts() {

	// TODO make a line separator between setLayoutForm and setLayoutTax

	content := container.New(
		layout.NewGridLayout(2),
		gui.createLayoutForm(),
		gui.createLayoutTax())

	gui.Window.SetContent(content)
}

// createLayoutForm Setup left side of window
func (gui *GUI) createLayoutForm() *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		gui.createLayoutIncome(),
		gui.createLayoutStatus(),
		gui.createLayoutChildren(),
		gui.createLayoutSave(),
	)
}

// createLayoutIncome Setup layouts and widget for income layout
func (gui *GUI) createLayoutIncome() *fyne.Container {
	gui.entryIncome = widgets.CreateIncomeEntry()
	gui.labelIncome = widget.NewLabel(gui.Language.Income)
	return container.New(
		layout.NewFormLayout(),
		gui.labelIncome,
		gui.entryIncome,
	)
}

// createLayoutStatus Setup layouts and widget for income layout
func (gui *GUI) createLayoutStatus() *fyne.Container {
	gui.radioStatus = widgets.CreateStatusRadio()
	gui.labelStatus = widget.NewLabel(gui.Language.Status)
	return container.NewHBox(
		gui.labelStatus,
		container.New(
			layout.NewVBoxLayout(),
			gui.radioStatus,
		),
	)

}

// createLayoutChildren Setup layouts and widget for income layout
func (gui *GUI) createLayoutChildren() *fyne.Container {
	gui.selectChildren = widgets.CreateChildrenSelect()
	gui.labelChildren = widget.NewLabel(gui.Language.Children)
	return container.NewHBox(
		gui.labelChildren,
		container.New(
			layout.NewVBoxLayout(),
			gui.selectChildren,
		),
	)
}

// createLayoutSave Setup layouts and widget for save button layout
func (gui *GUI) createLayoutSave() *fyne.Container {
	gui.buttonSave = widget.NewButton(gui.Language.Save, func() {
		gui.calculate()
		log.Printf("Save Tax") // TODO log debug save
	})
	return container.NewHBox(gui.buttonSave)
}

// createLayoutTax Setup right side of window
func (gui *GUI) createLayoutTax() *fyne.Container {
	// TODO make a line separator between createLayoutTaxResult and createLayoutTaxDetails
	return container.New(
		layout.NewVBoxLayout(),
		gui.createLayoutTaxResult(),
		gui.createLayoutTaxDetails(),
	)
}

// createLayoutTaxResult Setup right top side of window
func (gui *GUI) createLayoutTaxResult() *fyne.Container {
	gui.labelTax = widget.NewLabel(gui.Language.Tax)
	gui.labelTaxValue = widget.NewLabel("")
	gui.labelRemainder = widget.NewLabel(gui.Language.Remainder)
	gui.labelRemainderValue = widget.NewLabel("")
	gui.labelShares = widget.NewLabel(gui.Language.Share)
	gui.labelSharesValue = widget.NewLabel("")
	return container.New(
		layout.NewGridLayout(3),
		gui.labelTax,
		gui.labelTaxValue,
		widget.NewLabelWithData(gui.Currency),

		gui.labelShares,
		gui.labelSharesValue,
		widget.NewLabelWithData(gui.Currency),

		gui.labelRemainder,
		gui.labelRemainderValue,
		widget.NewLabelWithData(gui.Currency),
	)

}

// createLayoutTax Setup right bottom side of window
func (gui *GUI) createLayoutTaxDetails() *fyne.Container {
	var trancheNumber int = 5
	gui.labelsTrancheTaxes = widgets.CreateTrancheTaxesLabels(trancheNumber)

	var headers []string = []string{"TRANCHE", "MIN", "MAX", "RATE", "TAX"} // TODO language taxDetails
	grid := container.New(layout.NewGridLayout(len(headers)))

	// Headers Rows
	for _, label := range headers {
		grid.Add(widget.NewLabel(label))
	}

	// Tranche rows
	for index := range gui.labelsTrancheTaxes {
		grid.Add(widget.NewLabel("Tranche " + utils.ConvertIntToString(index+1)))
		var min string = utils.ConvertIntToString(gui.Config.Tax.Tranches[index].Min) + " €"
		grid.Add(widget.NewLabel(min))

		var max string = utils.ConvertIntToString(gui.Config.Tax.Tranches[index].Max) + " €"
		if gui.Config.Tax.Tranches[index].Max == math.MaxInt64 {
			max = "-"
		}
		grid.Add(widget.NewLabel(max))

		var rate string = gui.Config.Tax.Tranches[index].Rate
		grid.Add(widget.NewLabel(rate))

		grid.Add(gui.labelsTrancheTaxes[index])
	}

	//  TODO create border border := container.NewBorder()
	btn_color := canvas.NewRectangle(color.NRGBA{R: 25, G: 25, B: 25, A: 255})
	return container.New(
		layout.NewMaxLayout(),
		btn_color,
		grid,
	)
}
