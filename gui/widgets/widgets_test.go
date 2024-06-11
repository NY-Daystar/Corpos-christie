package widgets

import (
	"testing"
)

// For testing
// $ cd gui/widgets
// $ go test -v

func TestWidgetCreateEntry(t *testing.T) {
	var entry = CreateEntry()

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
	var entry = CreateChildrenSelect()

	if entry == nil {
		t.Errorf("No entry widget created")
	}
}
