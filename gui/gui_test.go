package gui

import (
	"testing"

	"fyne.io/fyne/v2"
	"github.com/NY-Daystar/corpos-christie/gui/settings"
	"github.com/NY-Daystar/corpos-christie/utils/colors"
)

// For testing
// $ cd gui
// $ go test -v

// Test to know if file icon is reachable
func TestIconAccess(t *testing.T) {
	var iconPath string = "../resources/assets/logo.ico"
	var icon fyne.Resource = settings.GetIcon(iconPath)

	if icon == nil {
		t.Errorf("Expected icon loaded with path '%v'", colors.Red(iconPath))
	}
}

// Test if language data in english and fresh can be load
func TestLoadLanguageData(t *testing.T) {
	var lang string = "en"
	var expected *string = &lang

	var language = settings.GetDefaultLanguage()
	t.Logf("Function result:\t%+v", *language)

	if *language != *expected {
		t.Errorf("Expected that the Mode '%v' should be equal to %v", colors.Red(*expected), colors.Red(*language))
	}
}

// Test if currency data can be loaded
func TestLoadCurrencyData(t *testing.T) {
	var expected int = 3
	var anotherExpected = "$"

	var currencies = settings.GetCurrencies()
	t.Logf("Function result:\t%+v", currencies)

	if len(currencies) != expected {
		t.Errorf("Expected that the currency number '%v' should be equal to %v", colors.Red(expected), colors.Red(len(currencies)))
	}

	if currencies[1] != anotherExpected {
		t.Errorf("Expected that the currency '%v' should be equal to %v", colors.Red(anotherExpected), colors.Red(currencies[1]))
	}
}
