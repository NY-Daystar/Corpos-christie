package widgets

import (
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

// CreateIncomeEntry Create widget entry for income
// Returns entry in fyne object
func CreateEntry(placeholder string) *widget.Entry {
	var entry = widget.NewEntry()
	entry.PlaceHolder = placeholder
	entry.Validator = validation.NewRegexp("^[0-9]{1,}$", "Not a number")
	return entry
}
