package widgets

import (
	"fyne.io/fyne/v2/widget"
)

// CreateStatusRadio Create widget radioGroup for marital status
// Returns radioGroup in fyne object
func CreateStatusRadio() *widget.RadioGroup {
	// TODO language through params
	var radio = widget.NewRadioGroup([]string{"Single", "Couple"}, nil)
	radio.SetSelected("Single")
	radio.Horizontal = true
	return radio
}
