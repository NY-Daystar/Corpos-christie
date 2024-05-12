// gui/view/main.go
package view

import (
	"fmt"
	"image/color"
	"net/url"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/model"
	"github.com/NY-Daystar/corpos-christie/gui/settings"
	"github.com/NY-Daystar/corpos-christie/gui/themes"
	"github.com/NY-Daystar/corpos-christie/gui/widgets"
	"github.com/NY-Daystar/corpos-christie/utils"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

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
	Model  *model.GUIModel
	App    fyne.App    // Fyne application
	Window fyne.Window // Fyne window
	Logger *zap.Logger

	// Widgets
	EntryIncome    *widget.Entry       // Input Entry to set income
	RadioStatus    *widget.RadioGroup  // Input Radio buttons to get status
	SelectChildren *widget.SelectEntry // Input Select to know how children
}

// NewView instantiate view with existing model (data)
func NewView(model *model.GUIModel, logger *zap.Logger) *GUIView {
	view := &GUIView{
		Model:  model,
		Logger: logger,
	}

	view.App = app.New() // Launch Fyne App

	view.EntryIncome = widgets.CreateIncomeEntry()
	view.RadioStatus = widgets.CreateStatusRadio()
	view.SelectChildren = widgets.CreateChildrenSelect()

	view.Logger.Info("Launch view")
	return view
}

// Start setup view of the application layout and menu
func (view *GUIView) Start() {
	view.Window = view.App.NewWindow(config.APP_NAME)
	view.Logger.Info("Load window", zap.Int("height", HEIGHT), zap.Int("width", WIDTH))

	view.setAppSettings()

	view.Window.Resize(fyne.NewSize(WIDTH, HEIGHT))
	view.Window.CenterOnScreen()

	// Set menu
	var menu *fyne.MainMenu = view.setMenu()
	view.Window.SetMainMenu(menu)

	// Create layouts and widgets
	view.setLayouts()
	view.Logger.Info("Layout and widgets loaded")

	view.setIcon()

	view.display()
}

// setSettings get and configure app settings
func (view *GUIView) setAppSettings() {
	view.Model.Settings, _ = settings.Load(view.Logger)

	view.Logger.Info("Settings loaded",
		zap.Int("theme", view.Model.Settings.Theme),
		zap.String("language", view.Model.Settings.Language),
		zap.String("theme", view.Model.Settings.Currency),
	)

	view.setTheme(view.Model.Settings.Theme)
	view.setLanguage(view.Model.Settings.Language)
	view.Model.Currency = binding.BindString(&view.Model.Settings.Currency)
}

// display Show and run the application
func (view *GUIView) display() {
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
		// gui.createLayoutSave(), // TODO
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
	// TODO mettre les labels des status couple/celibataire avec la conversion de langue
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

// TODO
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
	var trancheNumber int = 5
	currency, _ := view.Model.Currency.Get()

	// Add header columns in grid
	grid := container.New(layout.NewGridLayout(trancheNumber))

	view.Model.LabelsTaxHeaders = binding.NewStringList()
	for index, header := range view.Model.Language.GetTaxHeaders() {
		view.Model.LabelsTaxHeaders.Append(header)
		h, _ := view.Model.LabelsTaxHeaders.GetItem(index)
		grid.Add(widget.NewLabelWithData(h.(binding.String)))
	}

	// Setup binding for min, max and taxes columns
	// TODO appeler ces methodes depuis le model en laissant ces methode private
	view.Model.LabelsMinTranche = binding.BindStringList(view.Model.CreateMinTrancheLabels(currency, view.Model.Config.Tax.Tranches))
	view.Model.LabelsMaxTranche = binding.BindStringList(view.Model.CreateMaxTrancheLabels(currency, view.Model.Config.Tax.Tranches))
	view.Model.LabelsTrancheTaxes = binding.BindStringList(view.Model.CreateTrancheTaxesLabels(trancheNumber, currency))

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

// TODO MENU
// TODO faire un package menu avec un structure GuiMenu
// createMenu create mainMenu for window
func (view *GUIView) setMenu() *fyne.MainMenu {
	return fyne.NewMainMenu(
		view.createFileMenu(),
		view.createHelpMenu(),
	)
}

// TODO MENU
// createFileMenu create file item in toolbar to handle app settings
func (view *GUIView) createFileMenu() *fyne.Menu {
	fileMenu := fyne.NewMenu(view.Model.Language.File,
		fyne.NewMenuItem(view.Model.Language.Settings, func() {
			dialog.ShowCustom(view.Model.Language.Settings, view.Model.Language.Close,
				container.NewVBox(
					view.createSelectTheme(),
					widget.NewSeparator(),
					view.createSelectLanguage(),
					widget.NewSeparator(),
					view.createSelectCurrency(),
					widget.NewSeparator(),
					view.createLabelLogs(),
				), view.Window)
		}),
		fyne.NewMenuItem(view.Model.Language.Quit, func() { view.App.Quit() }),
	)
	return fileMenu
}

// TODO MENU
// createHelpMenu create help item in toolbar to show about app
func (view *GUIView) createHelpMenu() *fyne.Menu {
	url, _ := url.Parse(config.APP_LINK)

	view.Model.LabelsAbout = binding.NewStringList()
	view.Model.LabelsAbout.Set(view.Model.Language.GetAbouts())
	var labels []binding.DataItem
	for index := range view.Model.Language.GetAbouts() {
		about, _ := view.Model.LabelsAbout.GetItem(index)
		labels = append(labels, about)
	}

	// Setup layouts with data
	firstLine := container.NewHBox(
		widget.NewLabelWithData(labels[0].(binding.String)),
		widget.NewLabel(config.APP_NAME),
		widget.NewLabelWithData(labels[1].(binding.String)),
	)
	secondLine := container.NewHBox(
		widget.NewLabelWithData(labels[2].(binding.String)),
		widget.NewHyperlink("GitHub Project", url),
		widget.NewLabelWithData(labels[3].(binding.String)),
	)
	thirdLine := widget.NewLabelWithData(labels[4].(binding.String))
	fourthLine := container.NewHBox(
		widget.NewLabel("Version:"),
		canvas.NewText(fmt.Sprintf("v%s", config.APP_VERSION), color.NRGBA{R: 218, G: 20, B: 51, A: 255}),
	)
	fifthLine := widget.NewLabel(fmt.Sprintf("%s: %s", view.Model.Language.Author, config.APP_AUTHOR))

	helpMenu := fyne.NewMenu(view.Model.Language.Help,
		fyne.NewMenuItem(view.Model.Language.About, func() {
			dialog.ShowCustom(view.Model.Language.About, view.Model.Language.Close,
				container.NewVBox(
					firstLine,
					secondLine,
					thirdLine,
					fourthLine,
					fifthLine,
				), view.Window)
		}))
	return helpMenu
}

// TODO MENU
// createSelectTheme create select to change theme
func (view *GUIView) createSelectTheme() *fyne.Container {
	selectTheme := widget.NewSelect(view.Model.Language.GetThemes(), nil)

	selectTheme.OnChanged = func(s string) {
		index := selectTheme.SelectedIndex()
		view.setTheme(index)
		view.Model.Settings.Set("theme", index) // TODO a mettre dans setTheme // Update model
	}
	selectTheme.SetSelectedIndex(view.Model.Settings.Theme)
	return container.NewHBox(
		widget.NewLabel(view.Model.Language.ThemeCode),
		selectTheme,
	)
}

// TODO MENU
// createSelectLanguage create select to change language
func (view *GUIView) createSelectLanguage() *fyne.Container {
	selectLanguage := widget.NewSelect(view.Model.Language.GetLanguages(), nil)
	selectLanguage.SetSelectedIndex(view.Model.GetLanguageIndex(view.Model.Language.Code))
	selectLanguage.OnChanged = func(s string) {
		index := selectLanguage.SelectedIndex()
		var getLanguage = func() string {
			switch index {
			case 0:
				return settings.ENGLISH
			case 1:
				return settings.FRENCH
			case 2:
				return settings.SPANISH
			case 3:
				return settings.GERMAN
			case 4:
				return settings.ITALIAN
			default:
				return settings.ENGLISH
			}
		}

		language := getLanguage()
		view.setLanguage(language)
		view.Model.Settings.Set("language", language) // TODO a mettre dans setLanguage // Update model
		view.Model.Reload()
	}

	return container.NewHBox(
		widget.NewLabel(view.Model.Language.LanguageCode),
		selectLanguage,
	)
}

// TODO MENU
// createSelectCurrency create select to change currency
func (view *GUIView) createSelectCurrency() *fyne.Container {
	selectCurrency := widget.NewSelect(settings.GetCurrencies(), func(currency string) {
		view.setCurrency(currency)
		view.Model.Settings.Set("currency", currency) // TODO a mettre dans setCurrency // Update model
		view.Model.Reload()
	})
	currency, _ := view.Model.Currency.Get()
	selectCurrency.SetSelected(currency)
	return container.NewHBox(
		widget.NewLabel(view.Model.Language.Currency),
		selectCurrency,
	)
}

// TODO MENU
// createLabelLogs create label to show logs
func (view *GUIView) createLabelLogs() *fyne.Container {
	return container.NewHBox(
		widget.NewLabel(view.Model.Language.Logs),
		widget.NewLabel(config.LOGS_PATH),
	)
}

// TODO controller
// setTheme change theme of the application
// (if param = 0 then dark if 1 then light)
func (view *GUIView) setTheme(theme int) {
	var t themes.Theme
	if theme == settings.DARK {
		t = themes.DarkTheme{}
	} else {
		t = themes.LightTheme{}
	}
	view.Logger.Info("Set theme", zap.Int("theme", theme))
	view.App.Settings().SetTheme(t)
}

// TODO controller
// setLanguage change language of the application
func (view *GUIView) setLanguage(code string) {
	view.Logger.Info("Set language", zap.String("code", code))

	var languageFile string = fmt.Sprintf("%s/%s.yaml", config.LANGUAGES_PATH, code)
	view.Logger.Debug("Load file for language", zap.String("file", languageFile))

	yamlFile, _ := os.ReadFile(languageFile)
	err := yaml.Unmarshal(yamlFile, &view.Model.Language)

	view.Model.Language.Code = code

	if err != nil {
		view.Logger.Sugar().Fatalf("Unmarshal language file %s: %v", languageFile, err)
	}

	view.Logger.Sugar().Debugf("Language Yaml %v", view.Model.Language)
}

// TODO controller
// setCurrency change language of the application
func (view *GUIView) setCurrency(currency string) {
	view.Logger.Info("Set currency", zap.String("currency", currency))
	view.Model.Currency.Set(currency)
}
