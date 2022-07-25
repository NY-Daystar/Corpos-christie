package widgets

import (
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

// CreateIncomeEntry Create widget entry for income
// Returns entry in fyne object
func CreateIncomeEntry() *widget.Entry {
	var entry *widget.Entry = widget.NewEntry()
	entry.SetPlaceHolder("30000")
	entry.Validator = validation.NewRegexp("^[0-9]{1,}$", "Not a number") // TODO language through params
	return entry
}
