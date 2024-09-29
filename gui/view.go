package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/layouts"
	"github.com/NY-Daystar/corpos-christie/gui/model"
	"github.com/NY-Daystar/corpos-christie/gui/widgets"
	"github.com/NY-Daystar/corpos-christie/settings"
	"github.com/NY-Daystar/corpos-christie/utils"
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
	Tabs *container.AppTabs // Tabs to handle layout

	// Tab tax
	EntryIncome    *widget.Entry       // Input Entry to set income
	RadioStatus    *widget.RadioGroup  // Input Radio buttons to get status
	SelectChildren *widget.SelectEntry // Input Select to know how children
	SaveButton     *widget.Button      // Input button to save excel
	SelectYear     *widget.Select      // Select to choose tax year

	// Tab reverse tax
	EntryRemainder *widget.Entry // Input Entry to set remainder wished

	// Tab history
	HistoryList         *widget.List   // items list in history
	PurgeHistoryButton  *widget.Button // Input button to purge history
	ExportHistoryButton *widget.Button // Input button to export all history

	MailPopup *widgets.MailPopup // Handle mail popup
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

	view.EntryIncome = widgets.CreateEntry(utils.ConvertIntToString(config.MIN_INCOME), view.Model.Language.ErrorsValidation)
	view.EntryRemainder = widgets.CreateEntry(utils.ConvertIntToString(config.MIN_INCOME), view.Model.Language.ErrorsValidation)
	view.RadioStatus = widgets.CreateStatusRadio()
	view.SelectChildren = widgets.CreateChildrenSelectEntry("0")
	view.SaveButton = widgets.CreateButtonLabel(view.Model.Language.Save)

	year, _ := view.Model.Year.Get()
	view.SelectYear = widgets.CreateYearSelect(settings.GetYears(view.Model.Config), year)

	view.PurgeHistoryButton = widgets.CreateButtonIcon(theme.DeleteIcon())
	view.ExportHistoryButton = widgets.CreateButtonIcon(theme.FileImageIcon())

	view.HistoryList = &widget.List{}

	view.MailPopup = widgets.CreateMailPopup(&view.Model.Language)

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
func (view *GUIView) Start(controller *GUIController, path string) {
	view.Model.Reload()
	view.showNewVersion(path)
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
	view.Tabs = container.NewAppTabs()

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
			name:   view.Model.Language.History,
			icon:   theme.FileIcon(),
			layout: &layouts.HistoryLayout{MainLayout: mainLayout},
		},
	}

	// Load Tabs
	for _, item := range views {
		tab := container.NewTabItemWithIcon(item.name, item.icon, item.layout.SetLayout())
		view.Tabs.Append(tab)
	}

	view.Tabs.SetTabLocation(container.TabLocationTop)
	return view.Tabs
}

// showNewVersion display popup to show new version
func (view *GUIView) showNewVersion(path string) {
	execPath, err := utils.GetExecutablePath()
	if err != nil {
		view.Logger.Warn("%v", zap.Error(err))
	}

	if path != "" {
		updateDialog := dialog.NewCustom(
			view.Model.Language.UpdatePopup.Title,
			view.Model.Language.Close,
			widget.NewLabel(
				fmt.Sprintf(
					`%v '%v' %v
%v: %v

1. %v
2. %v : %v
3. %v
`,
					view.Model.Language.UpdatePopup.Text1,
					config.APP_NAME,
					view.Model.Language.UpdatePopup.Text2,
					view.Model.Language.UpdatePopup.Text3,
					path,
					view.Model.Language.UpdatePopup.Text4,
					view.Model.Language.UpdatePopup.Text5,
					execPath,
					view.Model.Language.UpdatePopup.Text6,
				),
			),
			view.Window,
		)
		updateDialog.Resize(fyne.NewSize(400, 300))
		updateDialog.Show()
	}
}

// getIncome Get value of widget entry of income
func (view *GUIView) GetIncome() int {
	intVal, err := utils.ConvertStringToInt(view.EntryIncome.Text)
	if err != nil {
		return 0
	}
	return intVal
}

// IsCoupleSelected Get value of widget radioGroup
// returns 1 if it's couple, 0 if single
func (view *GUIView) IsCoupleSelected() bool {
	return utils.FindIndex(view.RadioStatus.Options, view.RadioStatus.Selected) == 1
}

// getChildren get value of widget select
func (view *GUIView) GetChildren() int {
	children, err := utils.ConvertStringToInt(view.SelectChildren.Entry.Text)
	if err != nil {
		return 0
	}
	return children
}

// getRemainder Get value of widget entry of taxes
func (view *GUIView) GetRemainder() float64 {
	intVal, err := utils.ConvertStringToFloat64(view.EntryRemainder.Text)
	if err != nil {
		return 0
	}
	return intVal
}

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
		SaveButton:     view.SaveButton,
		SelectYear:     view.SelectYear,

		EntryRemainder: view.EntryRemainder,

		HistoryList:         view.HistoryList,
		PurgeHistoryButton:  view.PurgeHistoryButton,
		ExportHistoryButton: view.ExportHistoryButton,

		MailPopup: view.MailPopup,
	}
}
