package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/NY-Daystar/corpos-christie/utils"
)

// Layout when tax tab is selected
type ReverseTaxLayout struct {
	MainLayout
}

// Set layout for tax tab
func (view ReverseTaxLayout) SetLayout() *fyne.Container {
	return container.New(layout.NewGridLayout(2),
		view.setLeftLayout(),
		view.setRightLayout(),
	)
}

// Load form for tax tab
func (view ReverseTaxLayout) setLeftLayout() *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		view.createLayoutRemainder(),
		view.createLayoutStatus(),
		view.createLayoutChildren(),
	)
}

// Load result for tax tab
func (view ReverseTaxLayout) setRightLayout() *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		view.createLayoutTaxYear(),
		view.createLayoutTaxResult(),
		container.NewVBox(widget.NewLabel(""), widget.NewSeparator(), widget.NewLabel("")),
		view.createLayoutTaxDetails(),
	)
}

// createLayoutRemainder Setup layouts and widget for income layout
func (view *ReverseTaxLayout) createLayoutRemainder() *fyne.Container {
	return container.New(
		layout.NewFormLayout(),
		widget.NewLabelWithData(binding.NewSprintf("%s (%s)", view.Model.LabelRemainder, view.Model.Currency)),
		view.EntryRemainder,
	)
}

// createLayoutStatus Setup layouts and widget for income layout
func (view *ReverseTaxLayout) createLayoutStatus() *fyne.Container {
	view.Model.LabelStatus = binding.BindString(&view.Model.Language.Status)
	return container.NewHBox(
		widget.NewLabelWithData(view.Model.LabelStatus),
		container.New(
			layout.NewVBoxLayout(),
			view.RadioStatus,
		),
	)
}

// createLayoutChildren Setup layouts and widget for income layout
func (view *ReverseTaxLayout) createLayoutChildren() *fyne.Container {
	view.Model.LabelChildren = binding.BindString(&view.Model.Language.Children)
	return container.NewHBox(
		widget.NewLabelWithData(view.Model.LabelChildren),
		container.New(
			layout.NewVBoxLayout(),
			view.SelectChildren,
		),
	)
}

// createLayoutTaxResult Setup right top side of window
func (view *ReverseTaxLayout) createLayoutTaxYear() *fyne.Container {
	return container.New(layout.NewGridLayout(8),
		widget.NewLabel(""),
		widget.NewLabel(""),
		widget.NewLabel(""),
		widget.NewLabel(""),
		widget.NewLabel(""),
		widget.NewLabel(""),
		widget.NewLabelWithData(view.Model.LabelYear),
		widget.NewLabelWithData(view.Model.Year),
	)
}

// createLayoutTaxResult Setup right top side of window
func (view *ReverseTaxLayout) createLayoutTaxResult() *fyne.Container {
	var taxBind = binding.NewSprintf("%s (%s)", view.Model.Tax, view.Model.Currency)
	var incomeBind = binding.NewSprintf("%s (%s)", view.Model.Income, view.Model.Currency)

	return container.New(layout.NewGridLayout(3),
		widget.NewLabelWithData(view.Model.LabelTax),
		widget.NewLabelWithData(taxBind),
		widget.NewLabel(""),

		widget.NewLabelWithData(view.Model.LabelIncome),
		widget.NewLabelWithData(incomeBind),
		widget.NewLabel(""),

		widget.NewLabelWithData(view.Model.LabelShares),
		widget.NewLabelWithData(view.Model.Shares),
		widget.NewLabel(""),
	)
}

// createLayoutTax Setup right bottom side of window
func (view *ReverseTaxLayout) createLayoutTaxDetails() *fyne.Container {
	var trancheNumber int = view.Model.LabelsTrancheTaxes.Length()

	// Add header columns in grid
	grid := container.New(layout.NewGridLayout(trancheNumber))

	for index, header := range view.Model.Language.GetTaxHeaders() {
		view.Model.LabelsTaxHeaders.Append(header)
		headerItem, _ := view.Model.LabelsTaxHeaders.GetItem(index)
		var headerBind = binding.NewSprintf("%s", headerItem)
		grid.Add(widget.NewLabelWithData(headerBind))
	}

	// Add Tranche rows in grid
	for index := 0; index < view.Model.LabelsTrancheTaxes.Length(); index++ {
		minItem, _ := view.Model.LabelsMinTranche.GetItem(index)
		maxItem, _ := view.Model.LabelsMaxTranche.GetItem(index)
		rateItem, _ := view.Model.LabelsRateTranche.GetItem(index)
		taxItem, _ := view.Model.LabelsTrancheTaxes.GetItem(index)

		minBind := binding.NewSprintf("%s %s", minItem, view.Model.Currency)
		maxBind := binding.NewSprintf("%s %s", maxItem, view.Model.Currency)
		rateBind := binding.NewSprintf("%s %%", rateItem)
		taxBind := binding.NewSprintf("%s %s", taxItem, view.Model.Currency)

		grid.Add(widget.NewLabel("Tranche " + utils.ConvertIntToString(index+1)))
		grid.Add(widget.NewLabelWithData(minBind))
		grid.Add(widget.NewLabelWithData(maxBind))
		grid.Add(widget.NewLabelWithData(rateBind))
		grid.Add(widget.NewLabelWithData(taxBind))
	}

	return container.New(
		layout.NewStackLayout(),
		grid,
	)
}
