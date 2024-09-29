package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/NY-Daystar/corpos-christie/gui/model"
	"github.com/NY-Daystar/corpos-christie/gui/widgets"
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
	SaveButton     *widget.Button      // Button to save in history tax results
	SelectYear     *widget.Select      // Select to choose tax year

	EntryRemainder *widget.Entry // Input Entry to set remainder wished

	HistoryList         *widget.List   // items list in history
	PurgeHistoryButton  *widget.Button // Input button to purge history
	ExportHistoryButton *widget.Button // Input button to export all history

	MailPopup *widgets.MailPopup // Handle mail popup
}
