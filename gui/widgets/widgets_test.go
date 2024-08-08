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
	var entry = CreateChildrenSelectEntry()

	if entry == nil {
		t.Errorf("No entry widget created")
	}
}
