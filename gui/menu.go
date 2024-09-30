package gui

import (
	"fmt"
	"image/color"
	"io"
	"net/url"
	"os"
	"path"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/NY-Daystar/corpos-christie/config"
	"github.com/NY-Daystar/corpos-christie/settings"
	"github.com/NY-Daystar/corpos-christie/updater"
	"github.com/NY-Daystar/corpos-christie/utils"
	"go.uber.org/zap"
)

// GUIMenu represents menu of window application
type GUIMenu struct {
	Controller          *GUIController
	App                 fyne.App       // Fyne application
	Window              fyne.Window    // Fyne window
	MainMenu            *fyne.MainMenu // Fyne menu
	entryLogs           *widget.Entry  // Multilines entry
	copyClipboardButton *widget.Button // Button to copy logs
	saveButton          *widget.Button // Button to save logs into file
	deleteButton        *widget.Button // Button to delete logs
}

// NewMenu create main menu for window application
func NewMenu(controller *GUIController) *GUIMenu {
	return &GUIMenu{
		Controller: controller,
		App:        controller.View.App,
		Window:     controller.View.Window,
	}
}

func (menu *GUIMenu) Start() {
	menu.MainMenu = fyne.NewMainMenu(
		menu.createFileMenu(),
		menu.createHelpMenu(),
	)
	menu.Window.SetMainMenu(menu.MainMenu)
}

// createFileMenu create file item in toolbar to handle app settings
func (menu *GUIMenu) createFileMenu() *fyne.Menu {
	settingsMenuItem := fyne.NewMenuItem(menu.Controller.Model.Language.Settings, menu.showFileItem)
	quitMenuItem := fyne.NewMenuItem(menu.Controller.Model.Language.Quit, func() { menu.App.Quit() })
	quitMenuItem.IsQuit = true

	return fyne.NewMenu(
		menu.Controller.Model.Language.File,
		settingsMenuItem,
		quitMenuItem,
	)
}

// createHelpMenu create help item in toolbar to show about app
func (menu *GUIMenu) createHelpMenu() *fyne.Menu {
	helpMenu := fyne.NewMenu(
		menu.Controller.Model.Language.Help,
		fyne.NewMenuItem(menu.Controller.Model.Language.About, menu.showAboutItem),
		fyne.NewMenuItem(menu.Controller.Model.Language.Update, menu.showUpdateItem),
	)
	return helpMenu
}

// createAboutDialog create dialog box for about content
func (menu *GUIMenu) createAboutDialog() *fyne.Container {
	url, _ := url.Parse(config.APP_LINK)

	menu.Controller.Model.LabelsAbout.Set(menu.Controller.Model.Language.GetAbouts())
	var labels []binding.DataItem
	for index := range menu.Controller.Model.Language.GetAbouts() {
		about, _ := menu.Controller.Model.LabelsAbout.GetItem(index)
		labels = append(labels, about)
	}

	// Setup layouts with data
	firstLine := container.NewHBox(
		widget.NewLabelWithData(labels[0].(binding.String)),
		widget.NewLabel(config.APP_NAME),
		widget.NewLabelWithData(labels[1].(binding.String)),
	)
	secondLine := container.NewHBox(
		widget.NewLabelWithData(labels[2].(binding.String)),
		widget.NewHyperlink("GitHub Project", url),
		widget.NewLabelWithData(labels[3].(binding.String)),
	)
	thirdLine := widget.NewLabelWithData(labels[4].(binding.String))
	fourthLine := container.NewHBox(
		widget.NewLabel("Version:"),
		canvas.NewText(fmt.Sprintf("v%s", config.APP_VERSION), color.NRGBA{R: 218, G: 20, B: 51, A: 255}),
	)
	fifthLine := container.NewHBox(
		widget.NewLabelWithData(labels[5].(binding.String)),
		widget.NewLabel(config.APP_AUTHOR),
	)

	return container.NewVBox(
		firstLine,
		secondLine,
		thirdLine,
		fourthLine,
		fifthLine,
	)
}

// createUpdateDialog create dialog box for updates
func (menu *GUIMenu) createUpdateDialog() *fyne.Container {
	check, _ := updater.IsNewUpdateAvailable(config.APP_VERSION)
	if !check {
		return container.NewVBox(
			container.NewHBox(
				widget.NewLabel("Pas de mise à jour"),
			),
		)
	}

	// Lancement de l'update avec progression
	// TODO a revoir le processus
	fmt.Printf("Demarrage de l'update\n")
	updater.StartUpdater()

	progress := widget.NewProgressBar()
	infinite := widget.NewProgressBarInfinite()

	// TODO mettre une popup dès le lancement de l'interface pour demander la mise à jour

	// TODO make a circular progress simualate with 5sec latences
	// TODO utiliser la demo : https://docs.fyne.io/started/demo.html
	// TODO la mettre dans le readme
	// TODO l'utiliser pour changer la barre de progression ou mettre un circular

	go func() {
		for i := 0.0; i <= 1.0; i += 0.1 {
			time.Sleep(time.Millisecond * 250)
			progress.SetValue(i)
		}
		infinite.Hide()
		if infinite.Hidden {
			fmt.Printf("Fin du check\n")
		}
	}()

	return container.NewVBox(
		container.NewHBox(widget.NewLabel("Vérification de mise à jour")),
		container.NewHBox(widget.NewLabel("Vérification de mise à jour")),
		container.NewVBox(progress, infinite),
	)
}

// Show dialog box for settings like change language, year, currency, etc...
func (menu *GUIMenu) showFileItem() {
	dialog.ShowCustom(menu.Controller.Model.Language.Settings, menu.Controller.Model.Language.Close,
		container.NewVBox(
			menu.createSelectTheme(),
			widget.NewSeparator(),
			menu.createSelectLanguage(),
			widget.NewSeparator(),
			menu.createSelectCurrency(),
			widget.NewSeparator(),
			menu.createLabelLogs(),
			widget.NewSeparator(),
		), menu.Window)
}

// Show dialog box about application (author, project and other)
func (menu *GUIMenu) showAboutItem() {
	dialog.ShowCustom(
		menu.Controller.Model.Language.About,
		menu.Controller.Model.Language.Close,
		menu.createAboutDialog(),
		menu.Window,
	)
}

// Show dialog box about update application
func (menu *GUIMenu) showUpdateItem() {
	fmt.Printf("Check update")
	dialog.ShowCustom(
		menu.Controller.Model.Language.Update,
		menu.Controller.Model.Language.Close,
		menu.createUpdateDialog(),
		menu.Window)
}

// Refresh change for each option in menu old language for new in model
func (menu *GUIMenu) Refresh(oldModelLanguage settings.Yaml) {
	if menu.Controller.Menu != nil && menu.Controller.Menu.MainMenu != nil {
		// for menuItem in level 1
		for _, item := range menu.Controller.Menu.MainMenu.Items {
			// For file option
			if item.Label == oldModelLanguage.File {
				item.Label = menu.Controller.Model.Language.File
			}
			// For help option
			if item.Label == oldModelLanguage.Help {
				item.Label = menu.Controller.Model.Language.Help
			}
		}

		// for menuItem 0 in level 2
		for _, item := range menu.Controller.Menu.MainMenu.Items[0].Items {
			// For settings option
			if item.Label == oldModelLanguage.Settings {
				item.Label = menu.Controller.Model.Language.Settings
			}
			// For quit option
			if item.Label == oldModelLanguage.Quit {
				item.Label = menu.Controller.Model.Language.Quit
			}
		}

		// for menuItem 1 in level 2
		for _, item := range menu.Controller.Menu.MainMenu.Items[1].Items {
			// For about option
			if item.Label == oldModelLanguage.About {
				item.Label = menu.Controller.Model.Language.About
			}
			if item.Label == oldModelLanguage.Update {
				item.Label = menu.Controller.Model.Language.Update
			}
		}
		menu.Controller.Menu.MainMenu.Refresh()
	}
}

// createSelectTheme create select to change theme
func (menu *GUIMenu) createSelectTheme() *fyne.Container {
	selectTheme := widget.NewSelect(menu.Controller.Model.Language.GetThemes(), nil)

	selectTheme.OnChanged = func(s string) {
		index := selectTheme.SelectedIndex()
		menu.Controller.SetTheme(index) // update model
	}

	selectTheme.SetSelectedIndex(menu.Controller.Model.Settings.Theme)

	return container.NewHBox(
		widget.NewLabel(menu.Controller.Model.Language.ThemeCode),
		selectTheme,
	)
}

// createSelectLanguage create select to change language
func (menu *GUIMenu) createSelectLanguage() *fyne.Container {
	selectLanguage := widget.NewSelect(menu.Controller.Model.Language.GetLanguages(), nil)
	selectLanguage.SetSelectedIndex(settings.GetLanguageIndexFromCode(menu.Controller.Model.Language.Code))
	selectLanguage.OnChanged = func(s string) {
		menu.Controller.SetLanguage(selectLanguage.SelectedIndex())
	}

	return container.NewHBox(
		widget.NewLabel(menu.Controller.Model.Language.LanguageCode),
		selectLanguage,
	)
}

// createSelectCurrency create select to change currency
func (menu *GUIMenu) createSelectCurrency() *fyne.Container {
	selectCurrency := widget.NewSelect(
		settings.GetCurrencies(),
		menu.Controller.SetCurrency,
	)

	currency, _ := menu.Controller.Model.Currency.Get()
	selectCurrency.SetSelected(currency)
	return container.NewHBox(
		widget.NewLabel(menu.Controller.Model.Language.Currency),
		selectCurrency,
	)
}

// createLabelLogs create label to show logs
func (menu *GUIMenu) createLabelLogs() *fyne.Container {
	return container.NewHBox(
		widget.NewLabel(menu.Controller.Model.Language.Logs),
		widget.NewButton(utils.GetLogsFile(), menu.showLogsDialog),
	)
}

// Show dialog box for Logs to show them, copy them or save them
func (menu *GUIMenu) showLogsDialog() {
	menu.copyClipboardButton = widget.NewButton(menu.Controller.Model.Language.LogsActions.ClipboardAction, menu.copyClipboard)
	menu.saveButton = widget.NewButton(menu.Controller.Model.Language.LogsActions.SaveAction, menu.saveLogs)
	menu.deleteButton = widget.NewButton(menu.Controller.Model.Language.LogsActions.DeleteAction, menu.purgeLogs)

	lines, err := utils.ReadFileLastNLines(utils.GetLogsFile(), 500)
	if err != nil {
		menu.Controller.View.Logger.Sugar().Errorf("Failed to read file: %v", err)
	}

	menu.entryLogs = widget.NewMultiLineEntry()
	menu.entryLogs.SetText(string(lines))
	scroll := container.NewScroll(menu.entryLogs)

	// NewSpacer push to the right buttons
	buttonContainer := container.NewHBox(layout.NewSpacer(), menu.copyClipboardButton, menu.saveButton, menu.deleteButton)

	// Init container
	contentContainer := container.NewBorder(
		buttonContainer,
		nil,
		nil,
		nil,
		scroll)

	customDialog := dialog.NewCustom(
		"Logs",
		menu.Controller.Model.Language.Close,
		contentContainer,
		menu.Window)

	customDialog.Resize(fyne.NewSize(1400, 900))

	customDialog.Show()
}

// Copy logs menu in clipBoard
func (menu *GUIMenu) copyClipboard() {
	var content = menu.entryLogs.Text
	menu.Window.Clipboard().SetContent(content)

	menu.copyClipboardButton.SetIcon(theme.ConfirmIcon())
	menu.copyClipboardButton.SetText(menu.Controller.Model.Language.LogsActions.ClipboardSuccess)

}

// Save logs into a file
func (menu *GUIMenu) saveLogs() {
	folderChan := make(chan string)

	dialog.NewFolderOpen(func(folder fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, menu.Window)
			menu.Controller.View.Logger.Error("Dialog show folder error: %v", zap.String("error", err.Error()))
			return
		}
		if folder != nil {
			folderChan <- folder.Path()
		}
	}, menu.Window).Show()

	go func() {
		for {
			var folderPath = <-folderChan
			var filename = "corpos-christie.json"
			var filePath = path.Join(folderPath, filename)

			in, err := os.Open(utils.GetLogsFile())
			if err != nil {
				return
			}
			defer in.Close()
			out, err := os.Create(filePath)
			if err != nil {
				return
			}
			defer func() {
				cerr := out.Close()
				if err == nil {
					err = cerr
				}
			}()
			if _, err = io.Copy(out, in); err != nil {
				return
			}
			err = out.Sync()

			menu.saveButton.SetIcon(theme.ConfirmIcon())
			menu.saveButton.SetText(menu.Controller.Model.Language.LogsActions.SaveSuccess)
		}
	}()
}

// Erase content of logs
func (menu *GUIMenu) purgeLogs() {
	os.Truncate(utils.GetLogsFile(), 0)
	menu.deleteButton.SetIcon(theme.ConfirmIcon())
	menu.deleteButton.SetText(menu.Controller.Model.Language.LogsActions.DeleteSuccess)
}
