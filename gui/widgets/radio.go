package widgets

import (
	"fyne.io/fyne/v2/widget"
)

// CreateStatusRadio Create widget radioGroup for marital status
// Returns radioGroup in fyne object
func CreateStatusRadio() *widget.RadioGroup {
	radio := widget.NewRadioGroup(nil, nil)
	radio.Horizontal = true
	return radio
}
