package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// TODO corriger les issues deepsource
// Layout to display history tab
type HistoryLayout struct {
	MainLayout
}

// Set layout for history tab
func (view HistoryLayout) SetLayout() *fyne.Container {
	return container.New(layout.NewGridLayout(1),
		view.setLeftLayout(),
	)
}

// Create list for history
func (view HistoryLayout) setLeftLayout() *fyne.Container {
	// TODO load data from file
	data := []struct {
		Date     string
		Income   string
		Couple   string
		Children string
		Icon     fyne.Resource
	}{
		{"2023-06-01", "40000", "Yes", "0", theme.DocumentIcon()},
		{"2023-06-02", "50000", "No", "1", theme.DocumentIcon()},
		{"2023-06-03", "70000", "No", "2", theme.DocumentIcon()},
	}

	list := widget.NewList(
		func() int { return len(data) },
		func() fyne.CanvasObject {
			dateLabel := widget.NewLabel("")
			num1Label := widget.NewLabel("")
			num2Label := widget.NewLabel("")
			num3Label := widget.NewLabel("")
			iconButton := widget.NewButtonWithIcon("", nil, func() {})
			return container.NewVBox(
				container.NewHBox(
					dateLabel,
					layout.NewSpacer(),
					num1Label,
					layout.NewSpacer(),
					num2Label,
					layout.NewSpacer(),
					num3Label,
					layout.NewSpacer(),
					iconButton,
					layout.NewSpacer(),
				),
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			line := o.(*fyne.Container).Objects[0].(*fyne.Container)
			children := line.Objects
			children[0].(*widget.Label).SetText(data[i].Date)
			children[2].(*widget.Label).SetText(data[i].Income)
			children[4].(*widget.Label).SetText(data[i].Couple)
			children[6].(*widget.Label).SetText(data[i].Children)
			children[8].(*widget.Button).SetIcon(data[i].Icon)

			children[8].(*widget.Button).OnTapped = func() {
				view.recalculate(data[i].Income)
			}
		})

	headers := container.NewHBox(
		widget.NewLabel("Date"), // TODO language
		layout.NewSpacer(),
		widget.NewLabel("Income"), // TODO language
		layout.NewSpacer(),
		widget.NewLabel("Couple"), // TODO language
		layout.NewSpacer(),
		widget.NewLabel("Children"), // TODO language
		layout.NewSpacer(),
		widget.NewLabel("Actions"), // TODO language
		layout.NewSpacer(),
	)

	return container.NewBorder(
		headers, nil, nil, nil, list,
	)
}

// Go into tab taxes to recalculate
func (view HistoryLayout) recalculate(income string) {
	view.Tabs.SelectIndex(0)
	view.EntryIncome.SetText(income)
}

// No use for this layout
func (view HistoryLayout) setRightLayout() *fyne.Container {
	return nil
}
