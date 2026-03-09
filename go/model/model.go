package model

import (
	"fmt"
	"math"
	"os"

	"corpos-christie/go/config"
	"corpos-christie/go/settings"
	"corpos-christie/go/utils"

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
	User      *User          // User params to use program
	Logger    *zap.Logger    // Logger
	Histories []History      // List of tax history saved

	// Settings
	Settings settings.Settings // Settings of the app
	Theme    string            // Theme for the application
	Language settings.Yaml     // Yaml struct with all language data
	Currency string            // Currency to display
	Year     string            // Year of tax calculation based on config

	// buttonSave *widget.Button // Label for save button

	// Bindings
	Income               string   // Bind for income value
	Tax                  string   // Bind for tax value
	Remainder            string   // Bind for remainder value
	Shares               string   // Bind for shares value
	LabelShares          string   // Bind for shares label
	LabelYear            string   // Bind for year label
	LabelIncome          string   // Bind for income label
	LabelStatus          string   // Bind for status label
	LabelChildren        string   // Bind for children label
	LabelTax             string   // Bind for tax label
	LabelRemainder       string   // Bind for remainder label
	LabelsAbout          []string // List of label in about modal
	LabelsTaxHeaders     []string // List of label for tax details headers
	LabelsMinTranche     []string // List of labels for min tranche in grid
	LabelsMaxTranche     []string // List of labels for max tranche in grid
	LabelsRateTranche    []string // List of labels for rate tranche in grid
	LabelsTrancheTaxes   []string // List of tranches tax label results
	LabelsHistoryHeaders []string // List of labels for history
}

// NewModel: instantiate data for the application
func NewModel(config *config.Config, user *User, logger *zap.Logger) *GUIModel {
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
	model.Currency = *model.Settings.Currency

	// Set tax year
	model.Year = *model.Settings.Year
	model.Config.Tax.Year, _ = utils.ConvertStringToInt(model.Year)
}

// Init data and binding for GUI
func (model *GUIModel) prepare() {
	// Setup binding for min, max and taxes columns
	model.LabelsMinTranche = *model.createTrancheLabels(MIN)
	model.LabelsMaxTranche = *model.createTrancheLabels(MAX)
	model.LabelsRateTranche = *model.createTrancheLabels(RATE)
	model.LabelsTrancheTaxes = *model.createTrancheLabels(VALUE)
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
			var min = utils.ConvertIntToString(tranche.Min)
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
	model.LabelIncome = model.Language.Income
	model.LabelStatus = model.Language.Status
	model.LabelChildren = model.Language.Children
	model.LabelTax = model.Language.Tax
	model.LabelYear = model.Language.Year
	model.LabelRemainder = model.Language.Remainder
	model.LabelShares = model.Language.Share

	// Reload List Binding string
	model.LabelsAbout = model.Language.GetAbouts()
	model.LabelsTaxHeaders = model.Language.GetTaxHeaders()
	model.LabelsHistoryHeaders = model.Language.GetHistoryHeaders()

	// Reload grid min tranches
	var minList []string
	for index := 0; index < len(model.LabelsMinTranche); index++ {
		var min = utils.ConvertIntToString(model.Config.Tax.Tranches[index].Min)
		minList = append(minList, min)
	}
	model.LabelsMinTranche = minList

	// Reload grid max tranches
	var maxList []string
	for index := 0; index < len(model.LabelsMaxTranche); index++ {
		var max = utils.ConvertIntToString(model.Config.Tax.Tranches[index].Max)
		if model.Config.Tax.Tranches[index].Max == math.MaxInt64 {
			max = "-"
		}
		maxList = append(maxList, max)
	}
	model.LabelsMaxTranche = maxList

	// Reload rate tranches
	var rateList []string
	for index := 0; index < len(model.LabelsRateTranche); index++ {
		var rate = utils.ConvertIntToString(model.Config.Tax.Tranches[index].Rate)
		rateList = append(rateList, rate)
	}
	model.LabelsRateTranche = rateList
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
