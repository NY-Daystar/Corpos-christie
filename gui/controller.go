package gui

import (
	"fmt"
	"os"

	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/settings"
	"github.com/NY-Daystar/corpos-christie/gui/themes"
	"github.com/NY-Daystar/corpos-christie/tax"
	"github.com/NY-Daystar/corpos-christie/utils"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// GuiController of MVC to handle event and data
type GUIController struct {
	Model  *GUIModel
	View   *GUIView
	Menu   *GUIMenu
	Logger *zap.Logger
}

// NewController instantiate new controller with data model and event to update view
func NewController(model *GUIModel, view *GUIView, logger *zap.Logger) *GUIController {
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
	controller.SetLanguage(*controller.Model.Settings.Language)
	controller.SetCurrency(*controller.Model.Settings.Currency)
	controller.SetYear(*controller.Model.Settings.Year)
}

// calculate Get values of gui to calculate tax
func (controller *GUIController) calculate() {
	controller.Model.User.Income = controller.getIncome()
	controller.Model.User.IsInCouple = controller.getStatus()
	controller.Model.User.Children = controller.getChildren()

	result := tax.CalculateTax(controller.Model.User, controller.Model.Config)
	controller.Logger.Sugar().Debugf("Result taxes %#v", result)

	var tax string = utils.ConvertInt64ToString(int64(result.Tax))
	var remainder string = utils.ConvertInt64ToString(int64(result.Remainder))
	var shares string = utils.ConvertInt64ToString(int64(result.Shares))

	// Set data in tax layout
	controller.Model.Tax.Set(tax)
	controller.Model.Remainder.Set(remainder)
	controller.Model.Shares.Set(shares)

	// Set Tax details
	currency, _ := controller.Model.Currency.Get()
	for index := 0; index < controller.Model.LabelsTrancheTaxes.Length(); index++ {
		var taxTranche string = utils.ConvertIntToString(int(result.TaxTranches[index].Tax))
		controller.Model.LabelsTrancheTaxes.SetValue(index, taxTranche+" "+currency)
	}
}

// getIncome Get value of widget entry
func (controller *GUIController) getIncome() int {
	intVal, err := utils.ConvertStringToInt(controller.View.EntryIncome.Text)
	if err != nil {
		return 0
	}
	return intVal
}

// getStatus Get value of widget radioGroup
func (controller *GUIController) getStatus() bool {
	return controller.View.RadioStatus.Selected == "Couple"
}

// getChildren get value of widget select
func (controller *GUIController) getChildren() int {
	children, err := utils.ConvertStringToInt(controller.View.SelectChildren.Entry.Text)
	if err != nil {
		return 0
	}
	return children
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
func (controller *GUIController) SetLanguage(code string) {
	controller.Logger.Info("Set language", zap.String("code", code))

	var languageFile string = fmt.Sprintf("%s/%s.yaml", config.LANGUAGES_PATH, code)
	controller.Logger.Debug("Load file for language", zap.String("file", languageFile))

	yamlFile, _ := os.ReadFile(languageFile)
	oldModelLanguage := controller.Model.Language
	err := yaml.Unmarshal(yamlFile, &controller.Model.Language)

	if err != nil {
		controller.Logger.Sugar().Fatalf("Unmarshal language file %s: %v", languageFile, err)
	}

	controller.Logger.Sugar().Debugf("Language Yaml %v", controller.Model.Language)
	controller.Model.Language.Code = code
	controller.Model.Settings.Set("language", code)
	controller.Model.Reload()

	controller.Logger.Debug("Renommage du menu", zap.String("file", languageFile))
	controller.Menu.Refresh(oldModelLanguage)
}

// SetCurrency change currency of the application
func (controller *GUIController) SetCurrency(currency string) {
	controller.Logger.Info("Set currency", zap.String("currency", currency))
	controller.Model.Currency.Set(currency)
	controller.Model.Settings.Set("currency", currency)
	controller.Model.Reload()
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
	controller.calculate() // Recalculate taxes
}
