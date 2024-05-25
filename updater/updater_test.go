package updater

import (
	"testing"
)

// For testing
// $ cd updater
// $ go test -v

// Calculate tax for a single person with 30000 of income
func TestIsNewUpdateAvailable(t *testing.T) {
	tests := []struct {
		version  string
		expected bool
	}{
		{
			version:  "0.0",
			expected: true,
		},
		{
			version:  "1.0",
			expected: true,
		}, {
			version:  "100.0",
			expected: false,
		},
	}

	for _, testCase := range tests {
		t.Logf("Version: %v", testCase.version)
		result := IsNewUpdateAvailable(testCase.version)
		t.Logf("Res: %v", result)
		t.Logf("Expected %v", testCase.expected)
		if testCase.expected != result {
			t.Errorf("Test case failed for given input version:%s - expected:%v", testCase.version, testCase.expected)
		}
	}
}
