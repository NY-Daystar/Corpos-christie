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
		},
		{
			version:  "2.1.0",
			expected: true,
		},
		{
			version:  "100.0",
			expected: false,
		},
	}

	for _, testCase := range tests {
		t.Logf("Version: %v", testCase.version)
		result, err := IsNewUpdateAvailable(testCase.version)
		t.Logf("Res: %v", result)
		t.Logf("Err: %v", err)
		t.Logf("Expected %v", testCase.expected)
		if err != nil && err.Error() == "rate limiter reached" {
			t.Logf("Unable to download the release")
			continue
		}
		if testCase.expected != result {
			t.Errorf("Test case failed for given input version:%s - expected:%v", testCase.version, testCase.expected)
		}
	}
}
