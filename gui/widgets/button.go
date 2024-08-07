package widgets

import (
	"fyne.io/fyne/v2/widget"
)

// CreateSaveButton Create widget button to save
// Returns button in fyne object
func CreateSaveButton(s string) *widget.Button {
	var btn = widget.NewButton(s, nil)
	return btn
}
