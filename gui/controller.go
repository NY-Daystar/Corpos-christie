package gui

import (
	"fyne.io/fyne/v2/container"
	"github.com/NY-Daystar/corpos-christie/gui/model"
	"github.com/NY-Daystar/corpos-christie/gui/settings"
	"github.com/NY-Daystar/corpos-christie/gui/themes"
	"github.com/NY-Daystar/corpos-christie/tax"
	"github.com/NY-Daystar/corpos-christie/utils"
	"go.uber.org/zap"
)

// GuiController of MVC to handle event and data
type GUIController struct {
	Model  *model.GUIModel
	View   *GUIView
	Menu   *GUIMenu
	Logger *zap.Logger
}

// NewController instantiate new controller with data model and event to update view
func NewController(model *model.GUIModel, view *GUIView, logger *zap.Logger) *GUIController {
	var controller *GUIController = &GUIController{
		Model:  model,
		View:   view,
		Logger: logger,
	}

	controller.prepare()

	controller.Logger.Info("Launch controller")
	return controller
}

// prepare Set the events/trigger of gui widgets
// set menu for application
func (controller *GUIController) prepare() {
	controller.Menu = NewMenu(controller)
	controller.setAppSettings()

	controller.View.EntryIncome.OnChanged = func(input string) {
		controller.calculate()
	}
	controller.View.RadioStatus.OnChanged = func(input string) {
		controller.calculate()
	}
	controller.View.SelectChildren.OnChanged = func(input string) {
		controller.calculate()
	}
	controller.View.EntryRemainder.OnChanged = func(input string) {
		controller.reverseCalculate()
	}

	// READ:  https://github.com/fyne-io/fyne/issues/3466
	controller.View.Tabs.OnSelected = func(item *container.TabItem) {
		var index = controller.View.Tabs.SelectedIndex()
		controller.Logger.Sugar().Infof("Change tab index: %v - %v", "value", index, item.Text)
	}

	controller.Logger.Info("Events loaded")

	controller.Logger.Info("Menu is set")
	controller.Menu.Start()
}

// setAppSettings get and configure app settings and synchronizing model and view
func (controller *GUIController) setAppSettings() {
	controller.View.Logger.Info("Settings loaded",
		zap.Int("theme", controller.Model.Settings.Theme),
		zap.String("language", *controller.Model.Settings.Language),
		zap.String("theme", *controller.Model.Settings.Currency),
		zap.String("year", *controller.Model.Settings.Year),
	)

	controller.SetTheme(controller.Model.Settings.Theme)
	controller.SetLanguage(settings.GetLanguageIndexFromCode(*controller.Model.Settings.Language))
	controller.SetCurrency(*controller.Model.Settings.Currency)
	controller.SetYear(*controller.Model.Settings.Year)
	controller.Model.Reload()
}

// calculate Get values of gui to calculate tax
func (controller *GUIController) calculate() {
	controller.Model.User.Income = controller.getIncome()
	// If income to low no need to calculate
	if controller.Model.User.Income < 10000 {
		return
	}
	controller.Model.User.IsInCouple = controller.IsCoupleSelected()
	controller.Model.User.Children = controller.getChildren()

	result := tax.CalculateTax(controller.Model.User, controller.Model.Config)
	controller.Logger.Sugar().Debugf("Result taxes %#v", result)

	// Set data in tax layout
	controller.Model.Tax.Set(utils.ConvertInt64ToString(int64(result.Tax)))
	controller.Model.Remainder.Set(utils.ConvertInt64ToString(int64(result.Remainder)))
	controller.Model.Shares.Set(utils.ConvertInt64ToString(int64(result.Shares)))

	// Set Tax details
	for index := 0; index < controller.Model.LabelsTrancheTaxes.Length(); index++ {
		var taxTranche string = utils.ConvertIntToString(int(result.TaxTranches[index].Tax))
		controller.Model.LabelsTrancheTaxes.SetValue(index, taxTranche)
	}

	// TODO sauvegarder dans un fichier data.json
}

// reverseCalculate Get values of gui to calculate income with taxes
func (controller *GUIController) reverseCalculate() {
	controller.Model.User.Remainder = controller.getRemainder()

	result := tax.CalculateReverseTax(controller.Model.User, controller.Model.Config)
	controller.Logger.Sugar().Debugf("Reverse taxes %#v", result)

	// Set data in tax layout
	controller.Model.Income.Set(utils.ConvertInt64ToString(int64(result.Income)))

	// Set taxes value
	var w = utils.ConvertInt64ToString(int64(result.Income))
	controller.View.EntryIncome.SetText(w)

	// Set Tax details
	for index := 0; index < controller.Model.LabelsTrancheTaxes.Length(); index++ {
		var taxTranche string = utils.ConvertIntToString(int(result.TaxTranches[index].Tax))
		controller.Model.LabelsTrancheTaxes.SetValue(index, taxTranche)
	}
}

// getIncome Get value of widget entry of income
func (controller *GUIController) getIncome() int {
	intVal, err := utils.ConvertStringToInt(controller.View.EntryIncome.Text)
	if err != nil {
		return 0
	}
	return intVal
}

// IsCoupleSelected Get value of widget radioGroup
// returns 1 if it's couple, 0 if single
func (controller *GUIController) IsCoupleSelected() bool {
	return utils.FindIndex(controller.View.RadioStatus.Options, controller.View.RadioStatus.Selected) == 1
}

// getChildren get value of widget select
func (controller *GUIController) getChildren() int {
	children, err := utils.ConvertStringToInt(controller.View.SelectChildren.Entry.Text)
	if err != nil {
		return 0
	}
	return children
}

// getRemainder Get value of widget entry of taxes
func (controller *GUIController) getRemainder() float64 {
	intVal, err := utils.ConvertStringToFloat64(controller.View.EntryRemainder.Text)
	if err != nil {
		return 0
	}
	return intVal
}

// SetTheme change theme of the application
// (if param = 0 then dark if 1 then light)
func (controller *GUIController) SetTheme(theme int) {
	var t themes.Theme
	if theme == settings.DARK {
		t = themes.DarkTheme{}
	} else {
		t = themes.LightTheme{}
	}
	controller.Logger.Info("Set theme", zap.Int("theme", theme))
	controller.View.App.Settings().SetTheme(t)
	controller.Model.Settings.Set("theme", theme) // Update model
}

// SetLanguage change language of the application
func (controller *GUIController) SetLanguage(index int) {
	code := settings.GetLanguageCodeFromIndex(index)
	oldModelLanguage := controller.Model.Language

	controller.Model.LoadLanguage(code)

	// Refactoring model with language
	controller.Logger.Sugar().Debugf("Language Yaml %v", controller.Model.Language)
	controller.Model.Language.Code = code
	controller.Model.Settings.Set("language", code)
	controller.Model.Reload()

	// Rename radioStatus
	controller.View.RadioStatus.Options = controller.Model.Language.GetMaritalStatus()
	controller.View.RadioStatus.SetSelected(controller.View.RadioStatus.Options[0])
	controller.View.RadioStatus.Refresh()

	// Refreshing menu
	controller.Menu.Refresh(oldModelLanguage)
}

// SetCurrency change currency of the application
func (controller *GUIController) SetCurrency(currency string) {
	controller.Logger.Info("Set currency", zap.String("currency", currency))
	controller.Model.Currency.Set(currency)
	controller.Model.Settings.Set("currency", currency)
}

// SetYear change Year to calculate taxes of the application
func (controller *GUIController) SetYear(year string) {
	controller.Logger.Info("Set year", zap.String("year", year))
	if year == "" {
		year = *settings.GetDefaultYear()
	}
	controller.Model.Year.Set(year)
	controller.Model.Settings.Set("year", year)
	controller.Model.Config.ChangeTax(utils.ConvertBindStringToInt(controller.Model.Year))
	controller.Model.Reload()
	controller.calculate()
}
