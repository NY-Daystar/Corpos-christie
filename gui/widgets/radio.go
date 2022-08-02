package widgets

import (
	"fyne.io/fyne/v2/widget"
)

// CreateStatusRadio Create widget radioGroup for marital status
// Returns radioGroup in fyne object
func CreateStatusRadio() *widget.RadioGroup {
	var radio *widget.RadioGroup = widget.NewRadioGroup([]string{"Single", "Couple"}, nil) // TODO language through params
	radio.SetSelected("Single")
	radio.Horizontal = true
	return radio
}
