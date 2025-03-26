package upgrade

import (
	"testing"
)

func TestGetVersionTag(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"v1.2.3+build.123", "v1.2.3"},
		{"v0.0.1+abc123", "v0.0.1"},
		{"v2.3.4", "v2.3.4"},
		{"v3.0.0+20230301", "v3.0.0"},
		{"v4.5.6+meta+data", "v4.5.6"},
		{"", ""},
	}

	for _, test := range tests {
		result := GetVersionTag(test.input)
		if result != test.expected {
			t.Errorf("GetVersionTag(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestGetMajorMinorVersion(t *testing.T) {
	tests := []struct {
		name        string
		version     string
		wantMajor   int
		wantMinor   int
		expectError bool
	}{
		{
			name:        "valid version with v prefix",
			version:     "v1.24",
			wantMajor:   1,
			wantMinor:   24,
			expectError: false,
		},
		{
			name:        "valid version without v prefix",
			version:     "2.10",
			wantMajor:   2,
			wantMinor:   10,
			expectError: false,
		},
		{
			name:        "invalid version missing minor part",
			version:     "v3",
			wantMajor:   0,
			wantMinor:   0,
			expectError: true,
		},
		{
			name:        "invalid version with non-integer major",
			version:     "vX.10",
			wantMajor:   0,
			wantMinor:   0,
			expectError: true,
		},
		{
			name:        "invalid version with non-integer minor",
			version:     "2.Y",
			wantMajor:   0,
			wantMinor:   0,
			expectError: true,
		},
		{
			name:        "valid version with additional patch number",
			version:     "4.20.2",
			wantMajor:   4,
			wantMinor:   20,
			expectError: false,
		},
		{
			name:        "empty version string",
			version:     "",
			wantMajor:   0,
			wantMinor:   0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMajor, gotMinor, err := getMajorMinorVersion(tt.version)

			if (err != nil) != tt.expectError {
				t.Errorf("getMajorMinorVersion() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if gotMajor != tt.wantMajor {
				t.Errorf("getMajorMinorVersion() gotMajor = %v, want %v", gotMajor, tt.wantMajor)
			}
			if gotMinor != tt.wantMinor {
				t.Errorf("getMajorMinorVersion() gotMinor = %v, want %v", gotMinor, tt.wantMinor)
			}
		})
	}
}
