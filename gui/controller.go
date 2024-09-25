package gui

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"image/color"
	"os"
	"path"
	"sort"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/gui/model"
	"github.com/NY-Daystar/corpos-christie/gui/themes"
	"github.com/NY-Daystar/corpos-christie/helper"
	"github.com/NY-Daystar/corpos-christie/settings"
	"github.com/NY-Daystar/corpos-christie/tax"
	"github.com/NY-Daystar/corpos-christie/user"
	"github.com/NY-Daystar/corpos-christie/utils"
	"go.uber.org/zap"
)

// GuiController of MVC to handle event and data
type GUIController struct {
	Model  *model.GUIModel
	View   *GUIView
	Menu   *GUIMenu
	Logger *zap.Logger
}

// NewController instantiate new controller with data model and event to update view
func NewController(model *model.GUIModel, view *GUIView, logger *zap.Logger) *GUIController {
	var controller *GUIController = &GUIController{
		Model:  model,
		View:   view,
		Logger: logger,
	}

	controller.prepare()

	controller.Logger.Info("Launch controller")
	return controller
}

// prepare Set the events/trigger of gui widgets
// set menu for application
func (controller *GUIController) prepare() {
	controller.Menu = NewMenu(controller)
	controller.setAppSettings()
	controller.loadHistory()

	controller.View.EntryIncome.OnChanged = func(input string) {
		if controller.canCalculate() {
			controller.calculate()
		}
	}
	controller.View.RadioStatus.OnChanged = func(input string) {
		if controller.canCalculate() {
			controller.calculate()
		}
	}
	controller.View.SelectChildren.OnChanged = func(input string) {
		if controller.canCalculate() {
			controller.calculate()
		}
	}
	controller.View.EntryRemainder.OnChanged = func(input string) {
		controller.reverseCalculate()
	}
	controller.View.SaveButton.OnTapped = func() {
		controller.save()
		dialog.ShowInformation(
			controller.Model.Language.SavePopup.ConfirmedTitle,
			controller.Model.Language.SavePopup.ConfirmedMessage,
			controller.View.Window,
		)
	}

	controller.View.SelectYear.OnChanged = func(year string) {
		controller.SetYear(year)
	}

	controller.View.PurgeHistoryButton.OnTapped = controller.purgeHistory
	controller.View.ExportHistoryButton.OnTapped = controller.exportAllHistory

	controller.View.MailPopup.SubmitButton.OnTapped = controller.prepareMail

	// Handle tabs selections
	// READ: https://github.com/fyne-io/fyne/issues/3466
	controller.View.Tabs.OnSelected = func(item *container.TabItem) {
		var index = controller.View.Tabs.SelectedIndex()
		controller.loadHistory()
		controller.Logger.Sugar().Infof("Change tab index: %v - %v", "value", index, item.Text)
	}

	controller.Logger.Info("Events loaded")

	controller.Logger.Info("Menu is set")
	controller.Menu.Start()
}

// setAppSettings get and configure app settings and synchronizing model and view
func (controller *GUIController) setAppSettings() {
	controller.View.Logger.Info("Settings loaded",
		zap.Int("theme", controller.Model.Settings.Theme),
		zap.String("language", *controller.Model.Settings.Language),
		zap.String("theme", *controller.Model.Settings.Currency),
		zap.String("year", *controller.Model.Settings.Year),
	)

	controller.SetTheme(controller.Model.Settings.Theme)
	controller.SetLanguage(settings.GetLanguageIndexFromCode(*controller.Model.Settings.Language))
	controller.SetCurrency(*controller.Model.Settings.Currency)
	controller.SetYear(*controller.Model.Settings.Year)
	controller.Model.Reload()
}

// Verify if we can calculate
func (controller *GUIController) canCalculate() bool {
	if controller.View.GetIncome() < config.MIN_INCOME {
		controller.View.SaveButton.Disable()
		return false
	}
	controller.View.SaveButton.Enable()
	return true
}

// calculate Get values of gui to calculate tax
func (controller *GUIController) calculate() {
	controller.Model.User.Income = controller.View.GetIncome()
	controller.Model.User.IsInCouple = controller.View.IsCoupleSelected()
	controller.Model.User.Children = controller.View.GetChildren()

	result := tax.CalculateTax(controller.Model.User, controller.Model.Config)
	controller.Logger.Sugar().Debugf("Result taxes %#v", result)

	// Set data in tax layout
	controller.Model.Tax.Set(utils.ConvertInt64ToString(int64(result.Tax)))
	controller.Model.Remainder.Set(utils.ConvertInt64ToString(int64(result.Remainder)))
	controller.Model.Shares.Set(utils.ConvertInt64ToString(int64(result.Shares)))

	// Set Tax details
	for index := 0; index < controller.Model.LabelsTrancheTaxes.Length(); index++ {
		var taxTranche string = utils.ConvertIntToString(int(result.TaxTranches[index].Tax))
		controller.Model.LabelsTrancheTaxes.SetValue(index, taxTranche)
	}
}

// reverseCalculate Get values of gui to calculate income with taxes
func (controller *GUIController) reverseCalculate() {
	controller.Model.User.Remainder = controller.View.GetRemainder()

	result := tax.CalculateReverseTax(controller.Model.User, controller.Model.Config)
	controller.Logger.Sugar().Debugf("Reverse taxes %#v", result)

	// Set data in tax layout
	controller.Model.Income.Set(utils.ConvertInt64ToString(int64(result.Income)))

	// Set taxes value
	var w = utils.ConvertInt64ToString(int64(result.Income))
	controller.View.EntryIncome.SetText(w)

	// Set Tax details
	for index := 0; index < controller.Model.LabelsTrancheTaxes.Length(); index++ {
		var taxTranche string = utils.ConvertIntToString(int(result.TaxTranches[index].Tax))
		controller.Model.LabelsTrancheTaxes.SetValue(index, taxTranche)
	}
}

// Delete history file and refresh list
func (controller *GUIController) purgeHistory() {
	dialog.NewConfirm(
		controller.View.Model.Language.PurgeHistory.ConfirmTitle,
		controller.View.Model.Language.PurgeHistory.Confirm,
		func(response bool) {
			if response {
				utils.DeleteFile(utils.GetHistoryFile())
				controller.Model.Histories = []model.History{}
				controller.View.HistoryList.Refresh()
				dialog.ShowInformation(
					controller.View.Model.Language.PurgeHistory.ConfirmedTitle,
					controller.View.Model.Language.PurgeHistory.Confirmed,
					controller.View.Window,
				)
			}
		},
		controller.View.Window,
	).Show()
}

// Button to delete history file and refresh list
func (controller *GUIController) exportAllHistory() {
	folderChan := make(chan string)

	dialog.NewFolderOpen(func(folder fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, controller.View.Window)
			controller.Logger.Error("Dialog error: %v", zap.String("error", err.Error()))
			return
		}
		if folder != nil {
			folderChan <- folder.Path()
		}
	}, controller.View.Window).Show()

	go func() error {
		for {
			var folderPath = <-folderChan
			var filename = "result.csv"
			var filePath = path.Join(folderPath, filename)

			var headers = []string{
				controller.Model.Language.Year,
				controller.Model.Language.HistoryHeaders.Income,
				controller.Model.Language.Tax,
				controller.Model.Language.Remainder,
				controller.Model.Language.HistoryHeaders.Couple,
				controller.Model.Language.HistoryHeaders.Children,
			}

			var data = [][]string{headers}

			for _, history := range controller.Model.Histories {

				var user = &user.User{
					Income:     history.Income,
					IsInCouple: history.Couple,
					Children:   history.Children,
				}
				result := tax.CalculateTax(user, controller.Model.Config)

				var year = utils.ConvertIntToString(controller.Model.Config.Tax.Year)
				var tax = utils.ConvertInt64ToString(int64(result.Tax))
				var remainder = utils.ConvertInt64ToString(int64(result.Remainder))
				var coupleStr = ""
				if history.Couple {
					coupleStr = controller.Model.Language.Yes
				} else {
					coupleStr = controller.Model.Language.No
				}

				var row = []string{
					year,
					utils.ConvertIntToString(history.Income),
					tax,
					remainder,
					coupleStr,
					utils.ConvertIntToString(history.Children),
				}
				data = append(data, row)

			}

			file, _ := os.Create(filePath)

			writer := csv.NewWriter(file)

			for _, value := range data {
				writer.Write(value)
			}

			dialog.ShowCustom(
				controller.Model.Language.Export.ExportMessage,
				controller.Model.Language.Close,
				container.NewHBox(
					widget.NewLabel(fmt.Sprintf("%s: ", controller.Model.Language.Export.ExportMessage)),
					canvas.NewText(filePath, color.NRGBA{R: 218, G: 20, B: 51, A: 255}),
				),
				controller.View.Window,
			)
			writer.Flush()
			file.Sync()
		}
	}()
}

// verify data to send mail
func (controller *GUIController) prepareMail() {
	if err := controller.View.MailPopup.EmailEntry.Validate(); err != nil {
		dialog.ShowError(err, controller.View.Window)
		return
	}

	if err := controller.View.MailPopup.SubjectEntry.Validate(); err != nil {
		dialog.ShowError(err, controller.View.Window)
		return
	}

	if err := controller.View.MailPopup.BodyEntry.Validate(); err != nil {
		dialog.ShowError(err, controller.View.Window)
		return
	}

	controller.Model.User = &user.User{
		Income:     controller.View.MailPopup.Income,
		IsInCouple: controller.View.MailPopup.IsInCouple,
		Children:   controller.View.MailPopup.Children,
	}

	var body = helper.FormatMail(controller.Model.User, controller.Model.Config, controller.Model.Settings, controller.Model.Language, controller.View.MailPopup)
	var subject = controller.View.MailPopup.SubjectEntry.Text
	controller.sendMail(subject, body)
}

// send mail of history
func (controller *GUIController) sendMail(subject, body string) {
	var from = controller.Model.Settings.Smtp.User
	var to = controller.View.MailPopup.EmailEntry.Text

	controller.Logger.Info("Send mail", zap.String("From", from), zap.String("To", to), zap.String("subject", subject), zap.String("body", body))

	var mail = helper.NewMail(from, to, subject, body)

	var smtpConfig = controller.Model.Settings.Smtp

	controller.Logger.Info("SMTP Server", zap.String("host", smtpConfig.Host), zap.Int("Port", smtpConfig.Port), zap.String("User", smtpConfig.User), zap.String("Password", smtpConfig.Password))

	var smtpClient = helper.NewSMTP(smtpConfig)
	var err = smtpClient.DialAndSend(mail)
	if err != nil {
		dialog.ShowError(fmt.Errorf("%s: %w", controller.Model.Language.MailPopup.Error, err), controller.View.Window)
	} else {
		dialog.ShowInformation(controller.Model.Language.MailPopup.Success, controller.Model.Language.MailPopup.SuccessDetails, controller.View.Window)
	}
}

// SetTheme change theme of the application
// (if param = 0 then dark if 1 then light)
func (controller *GUIController) SetTheme(theme int) {
	var t themes.Theme
	if theme == settings.DARK {
		t = themes.DarkTheme{}
	} else {
		t = themes.LightTheme{}
	}
	controller.Logger.Info("Set theme", zap.Int("theme", theme))
	controller.View.App.Settings().SetTheme(t)
	controller.Model.Settings.Set("theme", theme) // Update model
}

// SetLanguage change language of the application
func (controller *GUIController) SetLanguage(index int) {
	code := settings.GetLanguageCodeFromIndex(index)
	oldModelLanguage := controller.Model.Language

	controller.Model.LoadLanguage(code)

	// Refactoring model with language
	controller.Logger.Sugar().Debugf("Language Yaml %v", controller.Model.Language)
	controller.Model.Language.Code = code
	controller.Model.Settings.Set("language", code)
	controller.Model.Reload()

	// Rename radioStatus
	controller.View.RadioStatus.Options = controller.Model.Language.GetMaritalStatus()
	controller.View.RadioStatus.SetSelected(controller.View.RadioStatus.Options[0])
	controller.View.RadioStatus.Refresh()

	// Refreshing menu
	controller.Menu.Refresh(oldModelLanguage)
}

// SetCurrency change currency of the application
func (controller *GUIController) SetCurrency(currency string) {
	controller.Logger.Info("Set currency", zap.String("currency", currency))
	controller.Model.Currency.Set(currency)
	controller.Model.Settings.Set("currency", currency)
}

// SetYear change Year to calculate taxes of the application
func (controller *GUIController) SetYear(year string) {
	controller.Logger.Info("Set year", zap.String("year", year))
	if year == "" {
		year = *settings.GetDefaultYear()
	}
	controller.Model.Year.Set(year)
	controller.Model.Settings.Set("year", year)
	controller.Model.Config.ChangeTax(utils.ConvertBindStringToInt(controller.Model.Year))
	controller.Model.Reload()
	controller.calculate()
}

// Save calculation in history file
func (controller *GUIController) save() {
	// Open history file
	var filePath = utils.GetHistoryFile()

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		controller.Logger.Error("Dialog save error: %v", zap.String("error", err.Error()))
		return
	}
	defer f.Close()

	// Format data to save
	var dateTime = time.Now().Format("2006-01-02 15:04:05")
	var history = &model.History{
		Date:     dateTime,
		Income:   controller.Model.User.Income,
		Couple:   controller.Model.User.IsInCouple,
		Children: controller.Model.User.Children,
	}
	byteArray, _ := json.Marshal(history)

	// Saving data
	if _, err := fmt.Fprintf(f, "%s\n", byteArray); err != nil {
		controller.Logger.Error("save error file: %v", zap.String("error", err.Error()))
	}
}

// Load history from file
func (controller *GUIController) loadHistory() {
	var lines = utils.GetHistory(utils.GetHistoryFile())

	var histories = make([]model.History, 0, len(lines))
	for _, line := range lines {
		var history = model.History{}
		json.Unmarshal([]byte(line), &history)
		if history.Couple {
			history.IsInCouple = controller.Model.Language.Yes
		} else {
			history.IsInCouple = controller.Model.Language.No
		}
		histories = append(histories, history)
	}

	// Sort by date
	sort.Slice(histories, func(i, j int) bool {
		timeI, _ := time.Parse("2006-01-02 15:05:05", histories[i].Date)
		timeJ, _ := time.Parse("2006-01-02 15:05:05", histories[j].Date)
		return timeI.After(timeJ)
	})

	controller.Model.Histories = histories
}
