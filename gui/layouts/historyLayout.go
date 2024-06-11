package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// TODO to develop
type HistoryLayout struct {
	MainLayout
}

// Set layout for tax tab
func (view HistoryLayout) SetLayout() *fyne.Container {
	return container.New(layout.NewGridLayout(2),
		view.setLeftLayout(),
		view.setRightLayout(),
	)
}

// TODO to develop
func (view HistoryLayout) setLeftLayout() *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		widget.NewLabel("LEFT LAYOUT"),
	)
}

// TODO to develop
func (view HistoryLayout) setRightLayout() *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		widget.NewLabel("RIGHT LAYOUT"),
	)
}
