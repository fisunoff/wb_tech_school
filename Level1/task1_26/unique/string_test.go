package unique

import "testing"

func TestCheckString(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"abcd", true},
		{"Привет", true},
		{"", true},
		{"abCdefAaf", false},
		{"aabcd", false},
		{"Aa", false},
		{"Привет, мир", false},
		{"Рр", false},
		{"Hello", false},
	}

	for _, tt := range tests {
		if got := CheckString(tt.input); got != tt.want {
			t.Errorf("CheckString(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
