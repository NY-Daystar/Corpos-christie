package widgets

import (
	"fyne.io/fyne/v2/widget"
)

// CreateTrancheLabels create widgets labels for tranche taxes value into an array
// Returns Array of label widget in fyne object
func CreateTrancheTaxesLabels(number int) []*widget.Label {
	var labels []*widget.Label = make([]*widget.Label, 0, number)
	for i := 1; i <= number; i++ {
		labels = append(labels, widget.NewLabel("0 €")) // TODO devise
	}
	return labels
}
