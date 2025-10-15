package unpacker

import "testing"

func TestUnpack(t *testing.T) {
	type testCase struct {
		name           string
		input          string
		expectedOutput string
		expectError    bool
	}

	tests := []testCase{
		{
			name:           "Standard case",
			input:          "a4bc2d5e",
			expectedOutput: "aaaabccddddde",
			expectError:    false,
		},
		{
			name:           "No digits",
			input:          "abcd",
			expectedOutput: "abcd",
			expectError:    false,
		},
		{
			name:           "Invalid string (starts with digit)",
			input:          "45",
			expectedOutput: "",
			expectError:    true,
		},
		{
			name:           "Empty string",
			input:          "",
			expectedOutput: "",
			expectError:    false,
		},
		{
			name:           "Escape digits",
			input:          "qwe\\4\\5",
			expectedOutput: "qwe45",
			expectError:    false,
		},
		{
			name:           "Escape digit then repeat",
			input:          "qwe\\45",
			expectedOutput: "qwe44444",
			expectError:    false,
		},
		{
			name:           "Escape backslash",
			input:          "qwe\\\\5",
			expectedOutput: "qwe\\\\\\\\\\", // qwe + 5 обратных слэшей
			expectError:    false,
		},
		{
			name:           "Invalid string (ends with escape)",
			input:          "abc\\",
			expectedOutput: "",
			expectError:    true,
		},
		{
			name:           "Single characters",
			input:          "a",
			expectedOutput: "a",
			expectError:    false,
		},
		{
			name:           "String with only one repeating char",
			input:          "a3",
			expectedOutput: "aaa",
			expectError:    false,
		},
		{
			name:           "Multi-digit repetition",
			input:          "a12b1",
			expectedOutput: "aaaaaaaaaaaab",
			expectError:    false,
		},
		{
			name:           "Unicode support",
			input:          "🙂2b",
			expectedOutput: "🙂🙂b",
			expectError:    false,
		},
		{
			name:           "Repetition count of zero",
			input:          "a0b3",
			expectedOutput: "bbb",
			expectError:    false,
		},
		{
			name:           "Repetition count of one",
			input:          "a1b1c1",
			expectedOutput: "abc",
			expectError:    false,
		},
		{
			name:           "Character followed by escaped number",
			input:          "a\\2",
			expectedOutput: "a2",
			expectError:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actualOutput, err := Unpack(tc.input)
			if tc.expectError {
				if err == nil {
					t.Errorf("expected an error, but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if actualOutput != tc.expectedOutput {
				t.Errorf("expected %q, but got %q", tc.expectedOutput, actualOutput)
			}
		})
	}
}
