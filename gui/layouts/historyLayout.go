package layouts

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"os"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/model"
	"github.com/NY-Daystar/corpos-christie/tax"
	"github.com/NY-Daystar/corpos-christie/utils"
	"go.uber.org/zap"
)

// Layout to display history tab
type HistoryLayout struct {
	MainLayout
}

// Set layout for history tab
func (view *HistoryLayout) SetLayout() *fyne.Container {
	return container.New(layout.NewStackLayout(),
		view.setLeftLayout(),
		view.setRightLayout(),
	)
}

// Create list for history
func (view *HistoryLayout) setLeftLayout() *fyne.Container {

	var initStructList = func() fyne.CanvasObject {
		return container.NewVBox(
			container.NewHBox(
				widget.NewLabel(""),
				layout.NewSpacer(),
				widget.NewLabel(""),
				layout.NewSpacer(),
				widget.NewLabel(""),
				layout.NewSpacer(),
				widget.NewLabel(""),
				layout.NewSpacer(),
				widget.NewButtonWithIcon("", nil, func() {}),
				widget.NewButtonWithIcon("", nil, func() {}),
				widget.NewButtonWithIcon("", nil, func() {}),
				layout.NewSpacer(),
			),
		)
	}

	var updateList = func(i widget.ListItemID, o fyne.CanvasObject) {
		line := o.(*fyne.Container).Objects[0].(*fyne.Container)
		children := line.Objects

		var history = view.Model.Histories[i]
		var date = history.Date
		var income = utils.ConvertIntToString(history.Income)
		var couple = history.Couple
		var coupleText = history.IsInCouple
		var childrenNumber = utils.ConvertIntToString(history.Children)
		var iconDoc = theme.DocumentIcon()
		var iconFile = theme.FileIcon()
		var iconMail = theme.MailSendIcon()

		children[0].(*widget.Label).SetText(date)
		children[2].(*widget.Label).SetText(income)
		children[4].(*widget.Label).SetText(coupleText)
		children[6].(*widget.Label).SetText(childrenNumber)
		children[8].(*widget.Button).SetIcon(iconDoc)
		children[9].(*widget.Button).SetIcon(iconFile)
		children[10].(*widget.Button).SetIcon(iconMail)

		children[8].(*widget.Button).OnTapped = func() {
			view.recalculate(income, couple, childrenNumber)
		}

		children[9].(*widget.Button).OnTapped = func() {
			folderChan := make(chan string)

			dialog.NewFolderOpen(func(folder fyne.ListableURI, err error) {
				if err != nil {
					dialog.ShowError(err, view.Window)
					view.Logger.Error("Dialog show folder error: %v", zap.String("error", err.Error()))
					return
				}
				if folder != nil {
					folderChan <- folder.Path()
				}
			}, view.Window).Show()

			go func() {
				for {
					var folderPath = <-folderChan
					var filename = "result.csv"
					var filePath = path.Join(folderPath, filename)

					view.exportCsv(filePath, income, couple, childrenNumber)

					dialog.ShowCustom(
						view.Model.Language.Export.ExportMessage,
						view.Model.Language.Close,
						container.NewHBox(
							widget.NewLabel(fmt.Sprintf("%s: ", view.Model.Language.Export.ExportMessage)),
							canvas.NewText(filePath, color.NRGBA{R: 218, G: 20, B: 51, A: 255}),
						),
						view.Window,
					)
				}
			}()
		}

		children[10].(*widget.Button).OnTapped = func() {
			view.showPopupMail(income, couple, childrenNumber)
		}
	}

	view.HistoryList = widget.NewList(
		func() int { return len(view.Model.Histories) },
		initStructList,
		updateList,
	)

	headers := container.NewHBox()

	for index, header := range view.Model.Language.GetHistoryHeaders() {
		view.Model.LabelsHistoryHeaders.Append(header)
		headerItem, _ := view.Model.LabelsHistoryHeaders.GetItem(index)
		var headerBind = binding.NewSprintf("%s", headerItem)
		headers.Add(widget.NewLabelWithData(headerBind))
		headers.Add(layout.NewSpacer())
	}

	globalActions := container.NewHBox(
		view.PurgeHistoryButton,
		view.ExportHistoryButton,
	)

	historyTable := container.NewBorder(
		headers, nil, nil, nil, view.HistoryList,
	)

	return container.NewBorder(
		globalActions, nil, nil, nil, historyTable,
	)
}

// Go into tab taxes to recalculate
func (view *HistoryLayout) recalculate(income string, couple bool, children string) {
	view.Tabs.SelectIndex(0)

	view.EntryIncome.SetText(income)
	view.SelectChildren.SetText(children)

	var option = 0
	if couple {
		option = 1
	}
	view.RadioStatus.SetSelected(view.RadioStatus.Options[option])
}

// Recalculate data in history to get tax
func (view *HistoryLayout) exportCsv(filePath, income string, couple bool, children string) error {
	incomeInt, _ := utils.ConvertStringToInt(income)
	childrenInt, _ := utils.ConvertStringToInt(children)
	view.Model.User = &model.User{
		Income:     incomeInt,
		IsInCouple: couple,
		Children:   childrenInt,
	}
	result := tax.CalculateTax(view.Model.User, view.Model.Config)

	var headers = []string{
		view.Model.Language.Year,
		view.Model.Language.HistoryHeaders.Income,
		view.Model.Language.Tax,
		view.Model.Language.Remainder,
		view.Model.Language.HistoryHeaders.Couple,
		view.Model.Language.HistoryHeaders.Children,
	}

	var year = utils.ConvertIntToString(view.Model.Config.Tax.Year)
	var tax = utils.ConvertInt64ToString(int64(result.Tax))
	var remainder = utils.ConvertInt64ToString(int64(result.Remainder))
	var coupleStr = ""
	if couple {
		coupleStr = view.Model.Language.Yes
	} else {
		coupleStr = view.Model.Language.No
	}

	var data = [][]string{headers, {year, income, tax, remainder, coupleStr, children}}
	file, _ := os.Create(filePath)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		writer.Write(value)
	}
	return file.Sync()
}

// Create popup with form to send mail
func (view *HistoryLayout) showPopupMail(income string, couple bool, children string) {
	view.MailPopup.SubjectEntry.Text = config.APP_NAME
	view.MailPopup.BodyEntry.Text = fmt.Sprintf("Hi, Here's your result for <em>%d</em>", view.Model.Config.Tax.Year)

	view.MailPopup.Income, _ = utils.ConvertStringToInt(income)
	view.MailPopup.IsInCouple = couple
	view.MailPopup.Children, _ = utils.ConvertStringToInt(children)

	var form = widget.NewForm(
		widget.NewFormItem(view.Model.Language.MailPopup.MailForm, view.MailPopup.EmailEntry),
		widget.NewFormItem(view.Model.Language.MailPopup.SubjectForm, view.MailPopup.SubjectEntry),
		widget.NewFormItem(view.Model.Language.MailPopup.BodyForm, view.MailPopup.BodyEntry),
	)

	var confirmationDialog = dialog.NewCustom(
		view.Model.Language.MailPopup.FormTitle,
		view.Model.Language.MailPopup.CloseForm,
		container.NewVBox(form,
			layout.NewSpacer(),
			view.MailPopup.SubmitButton,
		),
		view.Window,
	)
	confirmationDialog.Resize(fyne.NewSize(600, 400))
	confirmationDialog.Show()
}

// No use for this layout
func (view *HistoryLayout) setRightLayout() *fyne.Container {
	return container.NewStack()
}
