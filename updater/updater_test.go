package updater

import (
	"testing"
)

// For testing
// $ cd updater
// $ go test -v

func TestGetTag(t *testing.T) {
	tests := []struct {
		rawTag   string
		expected bool
	}{
		{
			rawTag:   "v1.2.3",
			expected: true,
		},
		{
			rawTag:   "v123.456.789",
			expected: true,
		},
		{
			rawTag:   "v12.34.56",
			expected: true,
		},
		{
			rawTag:   "v1.2",
			expected: false,
		},
		{
			rawTag:   "v1.2.3.4",
			expected: false,
		},
		{
			rawTag:   "va.b.c",
			expected: false,
		},
		{
			rawTag:   "v1.2a.3",
			expected: false,
		},
	}
	// Vérifier chaque chaîne pour voir si elle correspond à l'expression régulière
	for _, test := range tests {
		var tag = GetTag(test.rawTag)
		if tag == "" && test.expected == true {
			t.Errorf("Incompatible tag: %v", tag)
		}
	}
}
