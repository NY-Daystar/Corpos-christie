package widgets

import (
	"testing"
)

// For testing
// $ cd gui/widgets
// $ go test -v

func TestWidgetCreateEntry(t *testing.T) {
	var entry = CreateEntry("10000")

	if entry == nil {
		t.Errorf("No entry widget created")
	}
}

func TestWidgetCreateStatusRadio(t *testing.T) {
	var entry = CreateStatusRadio()

	if entry == nil {
		t.Errorf("No entry widget created")
	}
}

func TestWidgetCreateChildrenSelect(t *testing.T) {
	var selectEntry = CreateChildrenSelectEntry("")

	if selectEntry == nil {
		t.Errorf("No entry widget created")
	}
}

func TestWidgetCreateYearSelect(t *testing.T) {
	var selectEntry = CreateYearSelect([]string{"2022", "2023", "2024"}, "")

	if selectEntry == nil {
		t.Errorf("No entry widget created")
	}
}
