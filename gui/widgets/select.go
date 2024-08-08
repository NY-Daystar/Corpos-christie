package widgets

import (
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

// CreateChildrenSelect Create widget select for children
// Returns select in fyne object
func CreateChildrenSelectEntry(selected string) *widget.SelectEntry {
	var selector = widget.NewSelectEntry([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
	selector.Validator = validation.NewRegexp("^[0-9]{1,}$", "Not a number")
	if selected != "" {
		selector.SetText("0")
	}

	return selector
}

// CreateYearSelect Create widget select for year
// Returns select in fyne object
func CreateYearSelect(options []string, selected string) *widget.Select {
	var selector = widget.NewSelect(options, nil)
	if selected != "" {
		selector.SetSelected(selected)
	}

	return selector
}
