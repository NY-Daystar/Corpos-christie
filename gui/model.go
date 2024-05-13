package gui

import (
	"math"

	"fyne.io/fyne/v2/data/binding"
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/settings"
	"github.com/NY-Daystar/corpos-christie/gui/themes"
	"github.com/NY-Daystar/corpos-christie/user"
	"github.com/NY-Daystar/corpos-christie/utils"
	"go.uber.org/zap"
)

// GUIModel data of the application
type GUIModel struct {
	Config *config.Config // Config to use correctly the program
	User   *user.User     // User params to use program
	Logger *zap.Logger

	// Settings
	Settings settings.Settings // Settings of the app
	Theme    themes.Theme      // Fyne theme for the application
	Language settings.Yaml     // Yaml struct with all language data
	Currency binding.String    // Currency to display

	// buttonSave *widget.Button // Label for save button

	// Bindings
	Tax                binding.String     // Bind for tax value
	Remainder          binding.String     // Bind for remainder value
	Shares             binding.String     // Bind for shares value
	LabelShares        binding.String     // Bind for shares label
	LabelIncome        binding.String     // Bind for income label
	LabelStatus        binding.String     // Bind for status label
	LabelChildren      binding.String     // Bind for children label
	LabelTax           binding.String     // Bind for tax label
	LabelRemainder     binding.String     // Bind for remainder label
	LabelsAbout        binding.StringList // List of label in about modal
	LabelsTaxHeaders   binding.StringList // List of label for tax details headers
	LabelsMinTranche   binding.StringList // List of labels for min tranche in grid
	LabelsMaxTranche   binding.StringList // List of labels for max tranche in grid
	LabelsTrancheTaxes binding.StringList // List of tranches tax label
}

// NewModel: instantiate data for the application
func NewModel(config *config.Config, user *user.User, logger *zap.Logger) *GUIModel {
	model := GUIModel{
		Config: config,
		User:   user,
		Logger: logger,
	}

	model.prepare()

	model.Logger.Info("Launch model")
	return &model
}

// prepare init data and binding
func (model *GUIModel) prepare() {
	model.Settings, _ = settings.Load(model.Logger)
	model.Currency = binding.BindString(&model.Settings.Currency)

	model.LabelsAbout = binding.NewStringList()
	model.LabelsTaxHeaders = binding.NewStringList()

	// Setup binding for min, max and taxes columns
	model.LabelsMinTranche = binding.BindStringList(model.createMinTrancheLabels())
	model.LabelsMaxTranche = binding.BindStringList(model.createMaxTrancheLabels())
	var trancheNumber int = 5 // TODO put in constants
	model.LabelsTrancheTaxes = binding.BindStringList(model.createTrancheTaxesLabels(trancheNumber))
}

// TODO UTILS
// getLanguageIndex get index to selectLanguage in settings from language of the app
func (model *GUIModel) GetLanguageIndex(langue string) int {
	switch langue {
	case settings.ENGLISH:
		return 0
	case settings.FRENCH:
		return 1
	default:
		return 0
	}
}

// CreateTrancheLabels create widgets labels for tranche taxes value into an array
// Create number of tranche with currency value
// Returns Array of label widget in fyne object
func (model *GUIModel) createTrancheTaxesLabels(number int) *[]string {
	currency, _ := model.Currency.Get()

	var labels []string = make([]string, 0, number)

	for i := 1; i <= number; i++ {
		labels = append(labels, "0"+" "+currency)
	}
	return &labels
}

// TODO merge with createMaxTrancheLabels
// createMinTrancheLabels create string from config.Tranche to create binding
// Returns Array string with min tranches value
func (model *GUIModel) createMinTrancheLabels() *[]string {
	var tranches []config.Tranche = model.Config.Tax.Tranches
	currency, _ := model.Currency.Get()
	var labels []string = make([]string, 0, len(tranches))

	for _, tranche := range tranches {
		var min string = utils.ConvertIntToString(tranche.Min) + " " + currency
		labels = append(labels, min)
	}

	return &labels
}

// TODO merge with createMinTrancheLabels
// createMaxTrancheLabels create string from config.Tranche to create binding
// Returns Array string with max tranches value
func (model *GUIModel) createMaxTrancheLabels() *[]string {
	var tranches []config.Tranche = model.Config.Tax.Tranches
	currency, _ := model.Currency.Get()
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

// reload Refresh widget who needed specially when language changed
func (model *GUIModel) Reload() {
	// Simple data bind
	model.LabelIncome.Set(model.Language.Income)
	model.LabelStatus.Set(model.Language.Status)
	model.LabelChildren.Set(model.Language.Children)
	model.LabelTax.Set(model.Language.Tax)
	model.LabelRemainder.Set(model.Language.Remainder)
	model.LabelShares.Set(model.Language.Share)

	// Handle widget
	// gui.buttonSave.SetText(gui.Language.Save) // TODO saveExcel

	// Reload about content
	model.LabelsAbout.Set(model.Language.GetAbouts())

	// Reload header tax details
	model.LabelsTaxHeaders.Set(model.Language.GetTaxHeaders())

	// Reload grid header
	model.LabelsTrancheTaxes.Set(*model.createTrancheTaxesLabels(model.LabelsTrancheTaxes.Length()))

	// Reload grid min tranches
	var minList []string
	currency, _ := model.Currency.Get()
	for index := 0; index < model.LabelsMinTranche.Length(); index++ {
		var min string = utils.ConvertIntToString(model.Config.Tax.Tranches[index].Min) + " " + currency
		minList = append(minList, min)
	}
	model.LabelsMinTranche.Set(minList)

	// Reload grid max tranches
	var maxList []string
	for index := 0; index < model.LabelsMaxTranche.Length(); index++ {
		var max string = utils.ConvertIntToString(model.Config.Tax.Tranches[index].Max) + " " + currency
		if model.Config.Tax.Tranches[index].Max == math.MaxInt64 {
			max = "-"
		}
		maxList = append(maxList, max)
	}
	model.LabelsMaxTranche.Set(maxList)
}
