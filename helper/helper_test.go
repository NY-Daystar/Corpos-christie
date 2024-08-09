package helper

import (
	"testing"

	"github.com/NY-Daystar/corpos-christie/settings"
)

// For testing
// $ cd helper
// $ go test -v

// Test Smtp
func TestSmtp(t *testing.T) {
	// Arrange
	var config = &settings.Smtp{}

	// Act
	var smtpClient = NewSMTP(config)

	// Assert
	if smtpClient == nil {
		t.Errorf("No smtp Client created")
	}
}

// Test mail
func TestMail(t *testing.T) {
	// Arrange
	var from = "sender@mail.com"
	var to = "receiver@mail.com"
	var subject = "Unit test"
	var body = "Unit test in corpos christie application"

	// Act
	var mail = NewMail(from, to, subject, body)

	// Assert
	if mail == nil {
		t.Errorf("No mail object created")
	}
}
