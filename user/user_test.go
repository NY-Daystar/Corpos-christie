package user

import (
	"log"
	"os"
	"testing"
)

// For testing
// $ cd user
// $ go test -v

// MoqStdIn Moq standard input to simulate user action
func MoqStdIn(value string) *os.File {
	content := []byte(value)
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	return tmpfile
}

func TestAskIncome(t *testing.T) {
	// Tricks to change Standard input
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	var expected = 32000
	var stringRef = "32000"
	os.Stdin = MoqStdIn(stringRef)

	var u User = User{}
	check, err := u.AskIncome()

	if !check {
		t.Errorf("AskIncome in error, err: %v", err)
	}
	if u.Income != expected {
		t.Errorf("AskIncome wrong value than expected %d - %d", u.Income, expected)
	}
}

func TestAskRemainder(t *testing.T) {
	// Tricks to change Standard input
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	var expected = 5254.
	var stringRef = "5254"
	os.Stdin = MoqStdIn(stringRef)

	var u User = User{}
	check, _ := u.AskRemainder()

	if !check {
		t.Error("AskRemainder in error")
	}
	if u.Remainder != expected {
		t.Errorf("AskRemainder wrong value than expected %f - %f", u.Remainder, expected)
	}
}

func TestIsInCoupleSayYes(t *testing.T) {
	// Tricks to change Standard input
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin
	var stringRef = "yes"
	os.Stdin = MoqStdIn(stringRef)

	var u User = User{}
	check, _ := u.AskIsInCouple()

	if !check {
		t.Error("AskIsInCouple in error")
	}
	if u.IsInCouple != true {
		t.Errorf("AskIsInCouple wrong value than expected %v", u.IsInCouple)
	}
}

func TestIsInCoupleSayNo(t *testing.T) {
	// Tricks to change Standard input
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin
	var stringRef = "no"
	os.Stdin = MoqStdIn(stringRef)

	var u User = User{}
	check, _ := u.AskIsInCouple()

	if check {
		t.Error("AskIsInCouple in error")
	}
	if u.IsInCouple {
		t.Errorf("AskIsInCouple wrong value than expected %v", u.IsInCouple)
	}
}

func TestIsInCoupleBadAnswer(t *testing.T) {
	// Tricks to change Standard input
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	var stringRef = "invalid answer"
	os.Stdin = MoqStdIn(stringRef)

	var u User = User{}
	_, err := u.AskIsInCouple()

	if err == nil {
		t.Errorf("AskIsInCouple has to have error, err: %v", err)
	}
}

func TestAskHasChildren(t *testing.T) {
	// Tricks to change Standard input
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	var expected = 3
	var stringRef = "3"
	os.Stdin = MoqStdIn(stringRef)

	var u User = User{}
	check, _ := u.AskHasChildren()

	if !check {
		t.Error("AskHasChildren in error")
	}
	if u.Children != expected {
		t.Errorf("AskHasChildren wrong value than expected %d - %d", u.Children, expected)
	}
}

func TestAskHasChildrenSkip(t *testing.T) {
	// Tricks to change Standard input
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	var expected = 0
	var stringRef = ""
	os.Stdin = MoqStdIn(stringRef)

	var u User = User{}
	check, _ := u.AskHasChildren()

	if !check {
		t.Error("AskHasChildren in error")
	}
	if u.Children != expected {
		t.Errorf("AskHasChildren wrong value than expected %d - %d", u.Children, expected)
	}
}

func TestAskTaxDetails(t *testing.T) {
	// Tricks to change Standard input
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	var stringRef = "no"
	os.Stdin = MoqStdIn(stringRef)

	var u User = User{}
	check, _ := u.AskTaxDetails()

	if check {
		t.Errorf("AskTaxDetails in error, err: %v", check)
	}
}

func TestAskTaxDetailsWrongAnswer(t *testing.T) {
	// Tricks to change Standard input
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	var stringRef = "wrong answer"
	os.Stdin = MoqStdIn(stringRef)

	var u User = User{}
	_, err := u.AskTaxDetails()

	if err == nil {
		t.Errorf("AskTaxDetails has to have error, err: %v", err)
	}
}

func TestAskRestart(t *testing.T) {
	// Tricks to change Standard input
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	var stringRef = "yes"
	os.Stdin = MoqStdIn(stringRef)

	var u User = User{}
	var check bool = u.AskRestart()

	if !check {
		t.Error("AskRestart in error")
	}
}

func TestIsIsolated(t *testing.T) {
	// Tricks to change Standard input
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	var stringRef = "32000"
	os.Stdin = MoqStdIn(stringRef)

	var u User = User{
		IsInCouple: false,
		Children:   1,
	}
	var check bool = u.IsIsolated()

	if !check {
		t.Error("user has to be isolated")
	}
}

func TestGetShares(t *testing.T) {
	var expected = 3.5
	var u User = User{Shares: 3.5}

	shares := u.GetShares()

	if shares != expected {
		t.Errorf("GetShares wrong value than expected %f - %f", shares, expected)
	}
}

func TestShow(t *testing.T) {
	var u User = User{
		Income:     50000,
		IsInCouple: true,
		Children:   5,
		Shares:     4.5,
		Tax:        6452,
		Remainder:  43548,
	}
	u.Show()
}
