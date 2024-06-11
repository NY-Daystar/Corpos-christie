package gui

import (
	"testing"

	"fyne.io/fyne/v2"
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/model"
	"github.com/NY-Daystar/corpos-christie/gui/settings"
	"github.com/NY-Daystar/corpos-christie/user"
	"github.com/NY-Daystar/corpos-christie/utils/colors"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

// For testing
// $ cd gui
// $ go test -v

// Test GUI components (model, view, controller)
func TestGUIComponents(t *testing.T) {
	var user = &user.User{}
	var logger *zap.Logger = zaptest.NewLogger(t)
	var cfg *config.Config = config.New()

	var model = model.NewModel(cfg, user, logger)
	var view = NewView(model, logger)
	var controller = NewController(model, view, logger)
	var menu = NewMenu(controller)

	// model.Reload()
	t.Logf("Gui Model %+v", model)
	t.Logf("Gui View %+v", view)
	t.Logf("Gui Controller %#v", controller)

	if model == nil {
		t.Errorf("No model loaded")
	}
	if view == nil {
		t.Errorf("No view loaded")
	}
	if controller == nil {
		t.Errorf("No controller loaded")
	}
	if menu == nil {
		t.Errorf("No menu loaded")
	}
	menu.ShowFileItem()
	menu.ShowAboutItem()
	menu.ShowUpdateItem()

	var yaml = settings.Yaml{
		Theme:         settings.ThemeYaml{Dark: "dark", Light: "light"},
		Languages:     settings.LanguageYaml{English: "en", French: "fr"},
		Abouts:        settings.AboutYaml{Text1: "My text 1", Text2: "My text 2", Text3: "My text 3", Text4: "My Text 4"},
		TaxHeaders:    settings.TaxHeadersYaml{Header1: "Header 1", Header2: "Header 2", Header3: "Header 3"},
		MaritalStatus: settings.MaritalStatusYaml{Single: "Single", Couple: "Couple"},
	}
	menu.Refresh(yaml)

	if view.EntryIncome == nil || view.RadioStatus == nil || view.SelectChildren == nil {
		t.Errorf("No widgets loaded")
	}
	view.EntryIncome.SetText("45000")
	view.RadioStatus.SetSelected("Couple")
	view.SelectChildren.SetText("0")
}

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
	var lang = "en"
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
