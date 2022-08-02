package widgets

import (
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

// CreateChildrenSelect Create widget select for children
// Returns select in fyne object
func CreateChildrenSelect() *widget.SelectEntry {
	var sel *widget.SelectEntry = widget.NewSelectEntry([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
	sel.SetText("0")
	sel.Validator = validation.NewRegexp("^[0-9]{1,}$", "Not a number")
	return sel
}
