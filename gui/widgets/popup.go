package widgets

import (
	"errors"

	"fyne.io/fyne/v2/widget"
	"github.com/NY-Daystar/corpos-christie/settings"
	"github.com/NY-Daystar/corpos-christie/utils"
)

// MailPopup view popup to send mail
type MailPopup struct {
	EmailEntry   *widget.Entry  // Email address to send mail
	SubjectEntry *widget.Entry  // Subject of mail
	BodyEntry    *widget.Entry  // Body of the mail
	SubmitButton *widget.Button // Button to confirm the sending
	Username     string         // Username to receive the mail
	Income       int            // Income get in history
	IsInCouple   bool           // Couple or no in history
	Children     int            // Children get in history
}

// CreateMailPopup Create widgets for mail popup
func CreateMailPopup(language settings.Yaml) *MailPopup {
	var emailEntry = widget.NewEntry()
	emailEntry.Validator = func(input string) error {
		return checkEmail(input, language.ErrorsValidation)
	}

	var subjectEntry = widget.NewEntry()
	subjectEntry.Validator = func(input string) error {
		return checkSubject(input, language.ErrorsValidation)
	}

	var bodyEntry = widget.NewMultiLineEntry()
	bodyEntry.Validator = func(input string) error {
		return checkBody(input, language.ErrorsValidation)
	}

	var submitButton = widget.NewButton(language.MailPopup.SubmitForm, nil)

	return &MailPopup{
		EmailEntry:   emailEntry,
		SubjectEntry: subjectEntry,
		BodyEntry:    bodyEntry,
		SubmitButton: submitButton,
	}
}

// mail validation control
func checkEmail(input string, language settings.ErrorsValidationYaml) error {
	if !utils.IsValidEmail(input) {
		return errors.New(language.InvalidMail)
	}
	return nil
}

// subject validation control
func checkSubject(input string, language settings.ErrorsValidationYaml) error {
	if input == "" {
		return errors.New(language.InvalidSubject)
	}
	return nil
}

// body validation control
func checkBody(input string, language settings.ErrorsValidationYaml) error {
	if input == "" {
		return errors.New(language.InvalidBody)
	}
	return nil
}
