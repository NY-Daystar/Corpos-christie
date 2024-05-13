package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/settings"
	"github.com/NY-Daystar/corpos-christie/gui/widgets"
	"github.com/NY-Daystar/corpos-christie/utils"
	"go.uber.org/zap"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Size and Position of application
const WIDTH = 1100
const HEIGHT = 540

// GUIView View of the application
type GUIView struct {
	Model  *GUIModel
	App    fyne.App    // Fyne application
	Window fyne.Window // Fyne window
	Logger *zap.Logger

	// Widgets
	EntryIncome    *widget.Entry       // Input Entry to set income
	RadioStatus    *widget.RadioGroup  // Input Radio buttons to get status
	SelectChildren *widget.SelectEntry // Input Select to know how children
}

// NewView instantiate view with existing model (data)
func NewView(model *GUIModel, logger *zap.Logger) *GUIView {
	view := &GUIView{
		Model:  model,
		Logger: logger,
	}

	view.prepare() // Init Fyne component to avoid error

	view.Logger.Info("Launch view")
	return view
}

// prepare initialize Fyne application and components to avoid error
func (view *GUIView) prepare() {
	// Setup Fyne App
	view.App = app.New() // Launch Fyne App
	view.EntryIncome = widgets.CreateIncomeEntry()
	view.RadioStatus = widgets.CreateStatusRadio()
	view.SelectChildren = widgets.CreateChildrenSelect()

	// Setup Fyne window
	view.Window = view.App.NewWindow(config.APP_NAME)
	view.Logger.Info("Load window", zap.Int("height", HEIGHT), zap.Int("width", WIDTH))

	view.Window.Resize(fyne.NewSize(WIDTH, HEIGHT))
	view.Window.CenterOnScreen()

	// Create layouts and widgets
	view.setLayouts()
	view.Logger.Info("Layout and widgets loaded")

	view.setIcon()
}

// Start display view of the application layout and menu
func (view *GUIView) Start(controller *GUIController) {
	view.Model.Reload()
	view.Window.ShowAndRun()
}

// setIcon create icon for GUI application
func (view *GUIView) setIcon() {
	var iconName string = "logo.ico"
	var iconPath string = fmt.Sprintf("%s/%s", config.ASSETS_PATH, iconName)
	var icon fyne.Resource = settings.GetIcon(iconPath)
	view.Window.SetIcon(icon)
	view.Logger.Info("Load icon", zap.String("name", iconName), zap.String("path", iconPath))
}

// setLayouts Setup components/widget in the window
func (view *GUIView) setLayouts() {
	content := container.New(layout.NewGridLayout(2),
		view.createLayoutForm(),
		view.createLayoutTax(),
	)
	view.Window.SetContent(content)
}

// createLayoutForm Setup left side of window
func (view *GUIView) createLayoutForm() *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		view.createLayoutIncome(),
		view.createLayoutStatus(),
		view.createLayoutChildren(),
		// gui.createLayoutSave(), // TODO saveExcel
	)
}

// createLayoutIncome Setup layouts and widget for income layout
func (view *GUIView) createLayoutIncome() *fyne.Container {
	view.Model.LabelIncome = binding.BindString(&view.Model.Language.Income)
	return container.New(
		layout.NewFormLayout(),
		widget.NewLabelWithData(view.Model.LabelIncome),
		view.EntryIncome,
	)
}

// createLayoutStatus Setup layouts and widget for income layout
func (view *GUIView) createLayoutStatus() *fyne.Container {
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
func (view *GUIView) createLayoutChildren() *fyne.Container {
	view.Model.LabelChildren = binding.BindString(&view.Model.Language.Children)
	return container.NewHBox(
		widget.NewLabelWithData(view.Model.LabelChildren),
		container.New(
			layout.NewVBoxLayout(),
			view.SelectChildren,
		),
	)
}

// TODO saveExcel
// createLayoutSave Setup layouts and widget for save button layout
// func (view *GUIView) createLayoutSave() *fyne.Container {
// 	gui.Model.buttonSave = widget.NewButton(gui.Model.Language.Save, func() {
// 		gui.Model.Calculate()
// 		gui.Logger.Info("Save Taxes")
// 		// TODO Export taxes data in csv and/or pdf
// 	})
// 	return container.NewHBox(gui.buttonSave)
// }

// createLayoutTax Setup right side of window
func (view *GUIView) createLayoutTax() *fyne.Container {
	return container.New(
		layout.NewVBoxLayout(),
		view.createLayoutTaxResult(),
		container.NewVBox(widget.NewLabel(""), widget.NewSeparator(), widget.NewLabel("")),
		view.createLayoutTaxDetails(),
	)
}

// createLayoutTaxResult Setup right top side of window
func (view *GUIView) createLayoutTaxResult() *fyne.Container {
	view.Model.LabelTax = binding.BindString(&view.Model.Language.Tax)
	view.Model.Tax = binding.NewString()

	view.Model.LabelShares = binding.BindString(&view.Model.Language.Share)
	view.Model.Shares = binding.NewString()

	view.Model.LabelRemainder = binding.BindString(&view.Model.Language.Remainder)
	view.Model.Remainder = binding.NewString()

	return container.New(layout.NewGridLayout(3),
		widget.NewLabelWithData(view.Model.LabelTax),
		widget.NewLabelWithData(view.Model.Tax),
		widget.NewLabelWithData(view.Model.Currency),

		widget.NewLabelWithData(view.Model.LabelShares),
		widget.NewLabelWithData(view.Model.Shares),
		widget.NewLabelWithData(view.Model.Currency),

		widget.NewLabelWithData(view.Model.LabelRemainder),
		widget.NewLabelWithData(view.Model.Remainder),
		widget.NewLabelWithData(view.Model.Currency),
	)
}

// createLayoutTax Setup right bottom side of window
func (view *GUIView) createLayoutTaxDetails() *fyne.Container {
	var trancheNumber int = 5 // TODO put in constants

	// Add header columns in grid
	grid := container.New(layout.NewGridLayout(trancheNumber))

	for index, header := range view.Model.Language.GetTaxHeaders() {
		view.Model.LabelsTaxHeaders.Append(header)
		h, _ := view.Model.LabelsTaxHeaders.GetItem(index)
		grid.Add(widget.NewLabelWithData(h.(binding.String)))
	}

	// Add Tranche rows in grid
	for index := 0; index < view.Model.LabelsTrancheTaxes.Length(); index++ {
		minItem, _ := view.Model.LabelsMinTranche.GetItem(index)
		maxItem, _ := view.Model.LabelsMaxTranche.GetItem(index)
		taxItem, _ := view.Model.LabelsTrancheTaxes.GetItem(index)
		var rate string = view.Model.Config.Tax.Tranches[index].Rate

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
