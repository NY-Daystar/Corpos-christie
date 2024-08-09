package settings

import (
	"path/filepath"
	"testing"

	"go.uber.org/zap/zaptest"
)

// For testing
// $ cd gui/settings
// $ go test -v

var yaml Yaml

func init() {
	yaml = Yaml{
		Theme:         ThemeYaml{Dark: "dark", Light: "light"},
		Languages:     LanguageYaml{English: "en", French: "fr"},
		Abouts:        AboutYaml{Text1: "My text 1", Text2: "My text 2", Text3: "My text 3", Text4: "My Text 4"},
		TaxHeaders:    TaxHeadersYaml{Header1: "Header 1", Header2: "Header 2", Header3: "Header 3"},
		MaritalStatus: MaritalStatusYaml{Single: "Single", Couple: "Couple"},
	}
}

func TestGetDefaultCurrency(t *testing.T) {
	var expected = EURO
	currency := GetDefaultCurrency()

	if *currency != expected {
		t.Error("Currency not expected")
	}
}

func TestGetAllCurrency(t *testing.T) {
	var expected = 3
	currencies := GetCurrencies()

	if len(currencies) != expected {
		t.Error("Not amount on currencies available")
	}
}

func TestGetIcon(t *testing.T) {
	iconPath, _ := filepath.Abs("../resources/assets/logo.ico")
	icon := GetIcon(iconPath)

	if icon == nil {
		t.Error("Icon not available")
	}
}

func TestGetDefaultLanguage(t *testing.T) {
	var expected = ENGLISH
	language := GetDefaultLanguage()

	if *language != expected {
		t.Error("Language not expected")
	}
}

func TestGetLanguageCodeFromIndex(t *testing.T) {
	var expected = ENGLISH
	index := GetLanguageCodeFromIndex(0)

	if index != expected {
		t.Error("not index expected")
	}

	expected = FRENCH
	index = GetLanguageCodeFromIndex(1)

	if index != expected {
		t.Error("not index expected")
	}

	expected = SPANISH
	index = GetLanguageCodeFromIndex(2)

	if index != expected {
		t.Error("not index expected")
	}

	expected = GERMAN
	index = GetLanguageCodeFromIndex(3)

	if index != expected {
		t.Error("not index expected")
	}

	expected = ITALIAN
	index = GetLanguageCodeFromIndex(4)

	if index != expected {
		t.Error("not index expected")
	}
}

func TestGetLanguageIndexFromCode(t *testing.T) {
	var expected = 0
	index := GetLanguageIndexFromCode(ENGLISH)

	if index != expected {
		t.Error("not index expected")
	}

	expected = 1
	index = GetLanguageIndexFromCode(FRENCH)

	if index != expected {
		t.Error("not index expected")
	}

	expected = 2
	index = GetLanguageIndexFromCode(SPANISH)

	if index != expected {
		t.Error("not index expected")
	}

	expected = 3
	index = GetLanguageIndexFromCode(GERMAN)

	if index != expected {
		t.Error("not index expected")
	}

	expected = 4
	index = GetLanguageIndexFromCode(ITALIAN)

	if index != expected {
		t.Error("not index expected")
	}

	expected = 0
	index = GetLanguageIndexFromCode("NOT EXPECTED")

	if index != expected {
		t.Error("not index expected")
	}
}

func TestGetThemesLanguage(t *testing.T) {
	var expected = 2
	themes := yaml.GetThemes()

	if len(themes) != expected {
		t.Error("Themes are not load")
	}
}

func TestGetLanguagesValues(t *testing.T) {
	var expected = 5
	languages := yaml.GetLanguages()

	if len(languages) != expected {
		t.Error("Language are not load")
	}
}

func TestGetAboutsLanguage(t *testing.T) {
	var expected = 6
	abouts := yaml.GetAbouts()

	if len(abouts) != expected {
		t.Error("Themes are not load")
	}
}

func TestGetTaxHeadersLanguage(t *testing.T) {
	var expected = 5
	headers := yaml.GetTaxHeaders()

	if len(headers) != expected {
		t.Error("Themes are not load")
	}
}

func TestGetMaritalStatusLanguage(t *testing.T) {
	var expected = 2
	maritalStatus := yaml.GetMaritalStatus()

	if len(maritalStatus) != expected {
		t.Error("Themes are not load")
	}
}

func TestGetHistoryHeadersLanguage(t *testing.T) {
	var expected = 5
	headers := yaml.GetHistoryHeaders()

	if len(headers) != expected {
		t.Error("Themes are not load")
	}
}

func TestGetDefaultTheme(t *testing.T) {
	var expected = LIGHT
	theme := GetDefaultTheme()

	if theme != expected {
		t.Error("Theme not expected")
	}
}

func TestGetDefaultYear(t *testing.T) {
	var expected = "2024"
	year := GetDefaultYear()

	if *year != expected {
		t.Error("Year not expected")
	}
}

func TestLoadSettings(t *testing.T) {
	logger := zaptest.NewLogger(t)
	settings, err := Load(logger, "")

	if err != nil {
		t.Error("Should not have an error")
	}
	if *settings.Currency == "" {
		t.Error("Should have settings")
	}
}

func TestLoadWithWrongPathExpectedDefaultSettings(t *testing.T) {
	var expected = EURO
	logger := zaptest.NewLogger(t)
	settings, err := Load(logger, "wrong-path")

	if err != nil {
		t.Error("Should not have an error")
	}
	if *settings.Currency != expected {
		t.Error("Should have settings with default value")
	}
}

func TestSetSettings(t *testing.T) {
	logger := zaptest.NewLogger(t)

	settings, _ := Load(logger, "")

	// Set theme
	var expectedTheme = 1
	settings.Set("theme", expectedTheme)

	if settings.Theme != expectedTheme {
		t.Error("Theme not set")
	}

	// Set language
	var expectedLanguage = SPANISH
	settings.Set("language", expectedLanguage)

	if *settings.Language != expectedLanguage {
		t.Error("Language not set")
	}

	// Set currency
	var expectedCurrency = DOLLAR
	settings.Set("currency", expectedCurrency)

	if *settings.Currency != expectedCurrency {
		t.Error("Currency not set")
	}

	// Set year
	var expectedYear = "2021"
	settings.Set("year", expectedYear)

	if *settings.Year != expectedYear {
		t.Error("Year not set")
	}
}
