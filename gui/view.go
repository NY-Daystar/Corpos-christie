package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/layouts"
	"github.com/NY-Daystar/corpos-christie/gui/model"
	"github.com/NY-Daystar/corpos-christie/gui/settings"
	"github.com/NY-Daystar/corpos-christie/gui/widgets"
	"go.uber.org/zap"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Size and Position of application
const WIDTH = 1100
const HEIGHT = 540

// GUIView View of the application
type GUIView struct {
	Model  *model.GUIModel
	App    fyne.App    // Fyne application
	Window fyne.Window // Fyne window
	Logger *zap.Logger

	// Widgets
	Tabs           *container.AppTabs  // Tabs to handle layout
	EntryIncome    *widget.Entry       // Input Entry to set income
	RadioStatus    *widget.RadioGroup  // Input Radio buttons to get status
	SelectChildren *widget.SelectEntry // Input Select to know how children
	EntryRemainder *widget.Entry       // Input Entry to set remainder wished
}

// NewView instantiate view with existing model (data)
func NewView(model *model.GUIModel, logger *zap.Logger) *GUIView {
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
	view.App = app.New() // Launch Fyne App
	view.EntryIncome = widgets.CreateEntry()
	view.EntryIncome.SetPlaceHolder("30000")
	view.EntryRemainder = widgets.CreateEntry()
	view.EntryRemainder.SetPlaceHolder("30000")
	view.RadioStatus = widgets.CreateStatusRadio()
	view.SelectChildren = widgets.CreateChildrenSelect()
	view.SelectChildren.SetText("0")

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
	view.Window.SetContent(view.createAppTabs())
}

// createAppTabs Setup tabs for taxes and widget for income layout
func (view *GUIView) createAppTabs() *container.AppTabs {
	mainLayout := view.clone()

	views := []struct {
		name   string
		icon   fyne.Resource
		layout layouts.ViewLayout
	}{
		{
			name:   view.Model.Language.Tax,
			icon:   theme.AccountIcon(),
			layout: &layouts.TaxLayout{MainLayout: mainLayout},
		},
		{
			name:   view.Model.Language.ReverseTax,
			icon:   theme.ComputerIcon(),
			layout: &layouts.ReverseTaxLayout{MainLayout: mainLayout},
		},
		{
			name:   "History", // TODO language
			icon:   theme.FileIcon(),
			layout: &layouts.HistoryLayout{MainLayout: mainLayout},
		},
	}

	// Load Tabs
	view.Tabs = container.NewAppTabs()
	for _, item := range views {
		tab := container.NewTabItemWithIcon(item.name, item.icon, item.layout.SetLayout())
		view.Tabs.Append(tab)
	}

	view.Tabs.SetTabLocation(container.TabLocationTop)
	return view.Tabs
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

// clone create a copy of data in view for every layouts
func (view *GUIView) clone() layouts.MainLayout {
	return layouts.MainLayout{
		Model:          view.Model,
		App:            view.App,
		Window:         view.Window,
		Logger:         view.Logger,
		Tabs:           view.Tabs,
		EntryIncome:    view.EntryIncome,
		RadioStatus:    view.RadioStatus,
		SelectChildren: view.SelectChildren,
		EntryRemainder: view.EntryRemainder,
	}
}
