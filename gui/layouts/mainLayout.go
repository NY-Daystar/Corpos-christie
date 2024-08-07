package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/NY-Daystar/corpos-christie/gui/model"
	"go.uber.org/zap"
)

// ViewLayout interface to define Tabs layouts in fyne
type ViewLayout interface {
	SetLayout() *fyne.Container
	setLeftLayout() *fyne.Container
	setRightLayout() *fyne.Container
}

// Clone of GuiView to avoid import cycle
type MainLayout struct {
	Model  *model.GUIModel
	App    fyne.App    // Fyne application
	Window fyne.Window // Fyne window
	Logger *zap.Logger

	// Widgets
	Tabs           *container.AppTabs  // Tabs to handle layout
	EntryIncome    *widget.Entry       // Input Entry to set income
	RadioStatus    *widget.RadioGroup  // Input Radio buttons to get status
	SelectChildren *widget.SelectEntry // Input Select to know how children
	EntryRemainder *widget.Entry       // Input Entry to set remainder wished
	SaveButton     *widget.Button      // Button to save in history tax results
}
