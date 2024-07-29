package model

import (
	"fmt"
	"math"
	"os"

	"fyne.io/fyne/v2/data/binding"
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/settings"
	"github.com/NY-Daystar/corpos-christie/gui/themes"
	"github.com/NY-Daystar/corpos-christie/user"
	"github.com/NY-Daystar/corpos-christie/utils"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// Enum for type of tranche
const (
	MIN   string = "MIN"
	MAX   string = "MAX"
	RATE  string = "RATE"
	VALUE string = "VALUE"
)

// GUIModel data of the application
type GUIModel struct {
	Config    *config.Config // Config to use correctly the program
	User      *user.User     // User params to use program
	Logger    *zap.Logger    // Logger
	Histories []History      // List of tax history saved

	// Settings
	Settings settings.Settings // Settings of the app
	Theme    themes.Theme      // Fyne theme for the application
	Language settings.Yaml     // Yaml struct with all language data
	Currency binding.String    // Currency to display
	Year     binding.String    // Year of tax calculation based on config

	// buttonSave *widget.Button // Label for save button

	// Bindings
	Income             binding.String     // Bind for income value
	Tax                binding.String     // Bind for tax value
	Remainder          binding.String     // Bind for remainder value
	Shares             binding.String     // Bind for shares value
	LabelShares        binding.String     // Bind for shares label
	LabelYear          binding.String     // Bind for year label
	LabelIncome        binding.String     // Bind for income label
	LabelStatus        binding.String     // Bind for status label
	LabelChildren      binding.String     // Bind for children label
	LabelTax           binding.String     // Bind for tax label
	LabelRemainder     binding.String     // Bind for remainder label
	LabelsAbout        binding.StringList // List of label in about modal
	LabelsTaxHeaders   binding.StringList // List of label for tax details headers
	LabelsMinTranche   binding.StringList // List of labels for min tranche in grid
	LabelsMaxTranche   binding.StringList // List of labels for max tranche in grid
	LabelsRateTranche  binding.StringList // List of labels for rate tranche in grid
	LabelsTrancheTaxes binding.StringList // List of tranches tax label results
}

// NewModel: instantiate data for the application
func NewModel(config *config.Config, user *user.User, logger *zap.Logger) *GUIModel {
	model := GUIModel{
		Config: config,
		User:   user,
		Logger: logger,
	}

	model.configure()
	model.prepare()

	model.Logger.Info("Launch model")
	return &model
}

// Set settings of model like language, currency and other
func (model *GUIModel) configure() {
	model.Settings, _ = settings.Load(model.Logger, "")
	var code = *model.Settings.Language

	model.LoadLanguage(code)

	// Refactoring model with language
	model.Logger.Sugar().Debugf("Language Yaml %v", model.Language)
	model.Language.Code = code
	model.Settings.Set("language", code)

	// Set currency
	model.Currency = binding.BindString(model.Settings.Currency)

	// Set tax year
	model.Year = binding.BindString(model.Settings.Year)
	model.Config.Tax.Year = utils.ConvertBindStringToInt(model.Year)
}

// Init data and binding for GUI
func (model *GUIModel) prepare() {
	model.LabelIncome = binding.NewString()
	model.Income = binding.NewString()
	model.LabelStatus = binding.NewString()
	model.LabelChildren = binding.NewString()
	model.LabelYear = binding.NewString()
	model.LabelTax = binding.NewString()
	model.Tax = binding.NewString()
	model.LabelRemainder = binding.NewString()
	model.Remainder = binding.NewString()
	model.LabelShares = binding.NewString()
	model.Shares = binding.NewString()
	model.LabelsAbout = binding.NewStringList()
	model.LabelsTaxHeaders = binding.NewStringList()

	// Setup binding for min, max and taxes columns
	model.LabelsMinTranche = binding.BindStringList(model.createTrancheLabels(MIN))
	model.LabelsMaxTranche = binding.BindStringList(model.createTrancheLabels(MAX))
	model.LabelsRateTranche = binding.BindStringList(model.createTrancheLabels(RATE))
	model.LabelsTrancheTaxes = binding.BindStringList(model.createTrancheLabels(VALUE))
}

// CreateTrancheLabels create widgets labels for each data of tranche taxes (min, max, rate, taxValue)
// Convert this value into an array
// Returns Array of label widget in fyne object
func (model *GUIModel) createTrancheLabels(enumTranche string) *[]string {
	var tranches = model.Config.Tax.Tranches
	var labels = make([]string, 0, len(tranches))

	// To handle `min` tranche
	if enumTranche == MIN {
		for _, tranche := range tranches {
			var min string = utils.ConvertIntToString(tranche.Min)
			labels = append(labels, min)
		}

		// To handle `max` tranche
	} else if enumTranche == MAX {
		for _, tranche := range tranches {
			var max = utils.ConvertIntToString(tranche.Max)
			if tranche.Max == math.MaxInt64 {
				max = "-"
			}
			labels = append(labels, max)
		}
		// To handle `rate` tranche
	} else if enumTranche == RATE {
		for _, tranche := range tranches {
			var rate = utils.ConvertIntToString(tranche.Max)
			labels = append(labels, rate)
		}
		// To handle `value` of tranche
	} else if enumTranche == VALUE {
		for i := 1; i <= len(tranches); i++ {
			labels = append(labels, "0")
		}
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
	model.LabelYear.Set(model.Language.Year)
	model.LabelRemainder.Set(model.Language.Remainder)
	model.LabelShares.Set(model.Language.Share)
	// Handle widget
	// gui.buttonSave.SetText(gui.Language.Save) // TODO saveExcel

	// Reload about content
	model.LabelsAbout.Set(model.Language.GetAbouts())

	// Reload header tax details
	model.LabelsTaxHeaders.Set(model.Language.GetTaxHeaders())

	// Reload grid min tranches
	var minList []string
	for index := 0; index < model.LabelsMinTranche.Length(); index++ {
		var min = utils.ConvertIntToString(model.Config.Tax.Tranches[index].Min)
		minList = append(minList, min)
	}
	model.LabelsMinTranche.Set(minList)

	// Reload grid max tranches
	var maxList []string
	for index := 0; index < model.LabelsMaxTranche.Length(); index++ {
		var max = utils.ConvertIntToString(model.Config.Tax.Tranches[index].Max)
		if model.Config.Tax.Tranches[index].Max == math.MaxInt64 {
			max = "-"
		}
		maxList = append(maxList, max)
	}
	model.LabelsMaxTranche.Set(maxList)

	// Reload rate tranches
	var rateList []string
	for index := 0; index < model.LabelsRateTranche.Length(); index++ {
		var rate = utils.ConvertIntToString(model.Config.Tax.Tranches[index].Rate)
		rateList = append(rateList, rate)
	}
	model.LabelsRateTranche.Set(rateList)
}

// readLanguage Load into model data language
func (model *GUIModel) LoadLanguage(code string) {
	var languageFile = fmt.Sprintf("%s/%s.yaml", config.LANGUAGES_PATH, code)
	model.Logger.Info("Configure settings with code language", zap.String("file", languageFile), zap.String("code", code))

	yamlFile, _ := os.ReadFile(languageFile)
	err := yaml.Unmarshal(yamlFile, &model.Language)

	if err != nil {
		model.Logger.Sugar().Fatalf("Unmarshal language file %s: %v", languageFile, err)
	}
}
