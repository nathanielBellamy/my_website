package env

import "testing"

func TestIsProd(t *testing.T) {
	tests := []struct {
		mode     string
		expected bool
	}{
		{"prod", true},
		{"remotedev", false},
		{"localhost", false},
		{"unknown", false},
		{"", false},
	}

	for _, test := range tests {
		t.Run(test.mode, func(t *testing.T) {
			if result := IsProd(test.mode); result != test.expected {
				t.Errorf("IsProd(%q) = %v; want %v", test.mode, result, test.expected)
			}
		})
	}
}

func TestIsLocalhost(t *testing.T) {
	tests := []struct {
		mode     string
		expected bool
	}{
		{"prod", false},
		{"remotedev", false},
		{"localhost", true},
		{"unknown", false},
		{"", true},
	}

	for _, test := range tests {
		t.Run(test.mode, func(t *testing.T) {
			if result := IsLocalhost(test.mode); result != test.expected {
				t.Errorf("IsLocalhost(%q) = %v; want %v", test.mode, result, test.expected)
			}
		})
	}
}

func TestIsRemotedev(t *testing.T) {
	tests := []struct {
		mode     string
		expected bool
	}{
		{"prod", false},
		{"remotedev", true},
		{"localhost", false},
		{"unknown", false},
		{"", false},
	}

	for _, test := range tests {
		t.Run(test.mode, func(t *testing.T) {
			if result := IsRemotedev(test.mode); result != test.expected {
				t.Errorf("IsRemotedev(%q) = %v; want %v", test.mode, result, test.expected)
			}
		})
	}
}
