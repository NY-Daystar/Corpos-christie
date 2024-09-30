package model

import (
	"os"
	"testing"
)

// For testing
// $ cd gui/model
// $ go test -v

// MoqStdIn Moq standard input to simulate user action
func MoqStdIn(value string) *os.File {
	content := []byte(value)
	tmpfile, _ := os.CreateTemp("", "example")

	defer os.Remove(tmpfile.Name()) // clean up

	tmpfile.Write(content)

	tmpfile.Seek(0, 0)

	return tmpfile
}

func TestIsIsolated(t *testing.T) {
	// Tricks to change Standard input
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	var stringRef = "32000"
	os.Stdin = MoqStdIn(stringRef)

	var u = User{
		IsInCouple: false,
		Children:   1,
	}
	var check = u.IsIsolated()

	if !check {
		t.Error("user has to be isolated")
	}
}
