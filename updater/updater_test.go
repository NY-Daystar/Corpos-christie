package updater

import (
	"testing"
)

// For testing
// $ cd updater
// $ go test -v

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
		//var w = errors.New("no release")
		if err != nil && err.Error() == "rate limiter reached" {
			t.Logf("Unable to download the release")
			continue
		}
		if testCase.expected != result {
			t.Errorf("Test case failed for given input version:%s - expected:%v", testCase.version, testCase.expected)
		}
	}
}
