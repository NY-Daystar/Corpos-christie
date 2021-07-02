// Handle program command
package core

import (
	"testing"

	"github.com/LucasNoga/corpos-christie/lib/colors"
)

// For testing
// $ cd core
// $ go test -v

// Test select mode when gui params
func TestSelectModeWithGUIParams(t *testing.T) {
	var expectedValue = "GUI"
	var args []string = []string{"main.go", "--gui"}

	var mode string = selectMode(args)
	t.Logf("Function result:\t%s", mode)

	if mode != expectedValue {
		t.Errorf("Expected that the Mode '%v' should be equal to %v", colors.Red(expectedValue), colors.Red(mode))
	}
}

// Test select mode when console params
func TestSelectModeWithConsoleParams(t *testing.T) {
	var expectedValue = "console"
	var args []string = []string{"main.go", "--console"}

	var mode string = selectMode(args)
	t.Logf("Function result:\t%s", mode)

	if mode != expectedValue {
		t.Errorf("Expected that the Mode '%v' should be equal to %v", colors.Red(expectedValue), colors.Red(mode))
	}
}

// Test select mode when no params default GUI
func TestSelectModeWithNoParams(t *testing.T) {
	var expectedValue = "GUI"
	var args []string = []string{}

	var mode string = selectMode(args)
	t.Logf("Function result:\t%s", mode)

	if mode != expectedValue {
		t.Errorf("Expected that the Mode '%v' should be equal to %v", colors.Red(expectedValue), colors.Red(mode))
	}
}
