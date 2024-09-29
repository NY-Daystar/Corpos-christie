package settings

// Handle the languages in GUI settings

import (
	"reflect"
)

// Enum for languages
const (
	FRENCH  string = "fr"
	ENGLISH string = "en"
	SPANISH string = "es"
	GERMAN  string = "ge"
	ITALIAN string = "it"
)

// Handle all data about language data
type Yaml struct {
	Code             string               // code of the language (fr, en, etc...)
	Theme            ThemeYaml            `yaml:"themes"`
	Languages        LanguageYaml         `yaml:"languages"`
	Abouts           AboutYaml            `yaml:"abouts"`
	SavePopup        SavePopupYaml        `yaml:"save_popup"`
	TaxHeaders       TaxHeadersYaml       `yaml:"tax_headers"`
	MaritalStatus    MaritalStatusYaml    `yaml:"status_list"`
	HistoryHeaders   HistoryHeadersYaml   `yaml:"history_headers"`
	PurgeHistory     PurgeHistoryYaml     `yaml:"purge_history"`
	LogsActions      LogsActionsYaml      `yaml:"logs_actions"`
	Export           ExportYaml           `yaml:"export"`
	MailPopup        MailPopupYaml        `yaml:"mail_popup"`
	ErrorsValidation ErrorsValidationYaml `yaml:"errors_validation"`
	Year             string               `yaml:"year"`
	Yes              string               `yaml:"yes"`
	No               string               `yaml:"no"`
	File             string               `yaml:"file"`
	Settings         string               `yaml:"settings"`
	Update           string               `yaml:"update"`
	Income           string               `yaml:"income"`
	Status           string               `yaml:"status"`
	Children         string               `yaml:"children"`
	Tax              string               `yaml:"tax"`
	ReverseTax       string               `yaml:"reverse_tax"`
	Remainder        string               `yaml:"remainder"`
	Result           string               `yaml:"result"`
	TotalTax         string               `yaml:"total_tax"`
	Share            string               `yaml:"share"`
	History          string               `yaml:"history"`
	Save             string               `yaml:"save"`
	ThemeCode        string               `yaml:"theme"`
	LanguageCode     string               `yaml:"language"`
	Currency         string               `yaml:"currency"`
	Logs             string               `yaml:"logs"`
	Help             string               `yaml:"help"`
	About            string               `yaml:"about"`
	Author           string               `yaml:"author"`
	Close            string               `yaml:"close"`
	Quit             string               `yaml:"quit"`
}

// Languages yaml struct for theme's app
type LanguageYaml struct {
	English string `yaml:"english"`
	French  string `yaml:"french"`
	Spanish string `yaml:"spanish"`
	German  string `yaml:"german"`
	Italian string `yaml:"italian"`
}

// About text yaml struct for theme's app
type AboutYaml struct {
	Text1 string `yaml:"text_1"`
	Text2 string `yaml:"text_2"`
	Text3 string `yaml:"text_3"`
	Text4 string `yaml:"text_4"`
	Text5 string `yaml:"text_5"`
	Text6 string `yaml:"text_6"`
}

// data for popup save
type SavePopupYaml struct {
	ConfirmedTitle   string `yaml:"confirmed_title"`
	ConfirmedMessage string `yaml:"confirmed_message"`
}

// data for mail popup
type MailPopupYaml struct {
	FormTitle      string `yaml:"form_title"`
	SubmitForm     string `yaml:"submit_form"`
	CloseForm      string `yaml:"close_form"`
	MailForm       string `yaml:"mail_form"`
	SubjectForm    string `yaml:"subject_form"`
	BodyForm       string `yaml:"body_form"`
	Error          string `yaml:"error"`
	Success        string `yaml:"success"`
	SuccessDetails string `yaml:"success_details"`
}

// Headers yaml for tax detail
type TaxHeadersYaml struct {
	Header1 string `yaml:"header_1"`
	Header2 string `yaml:"header_2"`
	Header3 string `yaml:"header_3"`
	Header4 string `yaml:"header_4"`
	Header5 string `yaml:"header_5"`
}

// Marital status for radio buttons
type MaritalStatusYaml struct {
	Single string `yaml:"single"`
	Couple string `yaml:"couple"`
}

// Headers yaml for history table
type HistoryHeadersYaml struct {
	Date     string `yaml:"header_1"`
	Income   string `yaml:"header_2"`
	Couple   string `yaml:"header_3"`
	Children string `yaml:"header_4"`
	Actions  string `yaml:"header_5"`
}

// Dialog to purge data
type PurgeHistoryYaml struct {
	ConfirmTitle   string `yaml:"confirm_title"`
	Confirm        string `yaml:"confirm"`
	ConfirmedTitle string `yaml:"confirmed_title"`
	Confirmed      string `yaml:"confirmed"`
}

// Dialog to control logs data
type LogsActionsYaml struct {
	ClipboardAction  string `yaml:"clipboard_action"`
	ClipboardSuccess string `yaml:"clipboard_success"`
	SaveAction       string `yaml:"save_action"`
	SaveSuccess      string `yaml:"save_success"`
	DeleteAction     string `yaml:"delete_action"`
	DeleteSuccess    string `yaml:"delete_success"`
}

// Data for export dialog box
type ExportYaml struct {
	ExportTitle   string `yaml:"export_title"`
	ExportMessage string `yaml:"export_message"`
}

// errors list
type ErrorsValidationYaml struct {
	NaN            string `yaml:"nan"`
	NotEnough      string `yaml:"not_enough"`
	InvalidMail    string `yaml:"invalid_mail"`
	InvalidBody    string `yaml:"invalid_body"`
	InvalidSubject string `yaml:"invalid_subject"`
}

// getLanguages return all languages
func getLanguages() []string {
	return []string{ENGLISH, FRENCH, SPANISH, GERMAN, ITALIAN}
}

// GetDefaultLanguage get value of last language selected (fr, en)
func GetDefaultLanguage() *string {
	var lang = ENGLISH
	return &lang
}

// GetLanguageIndex get index to select language in settings from language of the app
func GetLanguageIndexFromCode(language string) int {
	for index, value := range getLanguages() {
		if value == language {
			return index
		}
	}
	return 0
}

// GetLanguageIndex get index to selectLanguage in settings from language of the app
func GetLanguageCodeFromIndex(index int) string {
	return getLanguages()[index]
}

// GetThemes parse ThemeYaml struct to get value of each field
func (yaml *Yaml) GetThemes() []string {
	v := reflect.ValueOf(yaml.Theme)
	themes := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		themes[i] = v.Field(i).String()
	}
	return themes
}

// GetLanguages parse LanguageYaml struct to get value of each field
func (yaml *Yaml) GetLanguages() []string {
	v := reflect.ValueOf(yaml.Languages)
	languages := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		languages[i] = v.Field(i).String()
	}
	return languages
}

// GetAbouts parse AboutYaml struct to get value of each field
func (yaml *Yaml) GetAbouts() []string {
	v := reflect.ValueOf(yaml.Abouts)
	abouts := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		abouts[i] = v.Field(i).String()
	}
	return abouts
}

// GetTaxHeaders parse TaxHeadersYaml struct to get value of each field
func (yaml *Yaml) GetTaxHeaders() []string {
	v := reflect.ValueOf(yaml.TaxHeaders)
	headers := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		headers[i] = v.Field(i).String()
	}
	return headers
}

// GetMaritalStatus parse MaritalStatusYaml struct to get value of each field
func (yaml *Yaml) GetMaritalStatus() []string {
	v := reflect.ValueOf(yaml.MaritalStatus)
	status := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		status[i] = v.Field(i).String()
	}
	return status
}

// GetHistoryHeaders parse HistoryHeaders struct to get value of each field
func (yaml *Yaml) GetHistoryHeaders() []string {
	v := reflect.ValueOf(yaml.HistoryHeaders)
	headers := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		headers[i] = v.Field(i).String()
	}
	return headers
}
