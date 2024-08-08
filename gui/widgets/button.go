package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// CreateButtonLabel Create widget button with label name
// Returns button in fyne object
func CreateButtonLabel(label string) *widget.Button {
	return widget.NewButton(label, nil)
}

// CreateButtonIcon Create widget button with an icon and no label
// Returns button in fyne object
func CreateButtonIcon(icon fyne.Resource) *widget.Button {
	return widget.NewButtonWithIcon("", icon, nil)
}
