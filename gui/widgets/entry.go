package widgets

import (
	"errors"
	"fmt"

	"fyne.io/fyne/v2/widget"
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/settings"
	"github.com/NY-Daystar/corpos-christie/utils"
)

// CreateEntry Create widget entry for income
// Returns entry in fyne object
func CreateEntry(placeholder string, errors settings.ErrorsValidationYaml) *widget.Entry {
	var entry = widget.NewEntry()
	entry.PlaceHolder = placeholder

	entry.Validator = func(input string) error {
		return checkEntry(input, errors)
	}
	return entry
}

// entry validation control
func checkEntry(input string, language settings.ErrorsValidationYaml) error {
	inputInt, err := utils.ConvertStringToInt(input)
	if err != nil {
		return errors.New(language.NaN)
	}

	if inputInt < config.MIN_INCOME {
		return fmt.Errorf("%s (> %d)", language.NotEnough, config.MIN_INCOME)
	}
	return nil
}
