package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/NY-Daystar/corpos-christie/utils"
)

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
	list := widget.NewList(
		func() int { return len(view.Model.Histories) },
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

			var history = view.Model.Histories[i]
			var date = history.Date
			var income = utils.ConvertIntToString(history.Income)
			var couple = history.Couple
			var coupleText = history.IsInCouple
			var childrenNumber = utils.ConvertIntToString(history.Children)
			var icon = theme.DocumentIcon()

			children[0].(*widget.Label).SetText(date)
			children[2].(*widget.Label).SetText(income)
			children[4].(*widget.Label).SetText(coupleText)
			children[6].(*widget.Label).SetText(childrenNumber)
			children[8].(*widget.Button).SetIcon(icon)

			children[8].(*widget.Button).OnTapped = func() {
				view.recalculate(income, couple, childrenNumber)
			}
		})

	headers := container.NewHBox(
		widget.NewLabel(view.Model.Language.HistoryHeaders.Date),
		layout.NewSpacer(),
		widget.NewLabel(view.Model.Language.HistoryHeaders.Income),
		layout.NewSpacer(),
		widget.NewLabel(view.Model.Language.HistoryHeaders.Couple),
		layout.NewSpacer(),
		widget.NewLabel(view.Model.Language.HistoryHeaders.Children),
		layout.NewSpacer(),
		widget.NewLabel(view.Model.Language.HistoryHeaders.Actions),
		layout.NewSpacer(),
	)

	return container.NewBorder(
		headers, nil, nil, nil, list,
	)
}

// Go into tab taxes to recalculate
func (view HistoryLayout) recalculate(income string, couple bool, children string) {
	view.Tabs.SelectIndex(0)

	view.EntryIncome.SetText(income)
	view.SelectChildren.SetText(children)

	var option = 0
	if couple {
		option = 1
	}
	view.RadioStatus.SetSelected(view.RadioStatus.Options[option])

}

// No use for this layout
func (view HistoryLayout) setRightLayout() *fyne.Container {
	return nil
}
