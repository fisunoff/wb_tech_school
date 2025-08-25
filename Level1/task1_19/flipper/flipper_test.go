package flipper

import "testing"

func TestFlip(t *testing.T) {
	type testCase struct {
		name  string
		input string
		want  string
	}

	tests := []testCase{
		{
			name:  "ASCII",
			input: "Hello, world",
			want:  "dlrow ,olleH",
		},
		{
			name:  "Кириллица",
			input: "Привет, мир!",
			want:  "!рим ,тевирП",
		},
		{
			name:  "Строка с Unicode символами (иероглифы, эмодзи)",
			input: "Go-랭 🚀",
			want:  "🚀 랭-oG",
		},
		{
			name:  "Пустая строка",
			input: "",
			want:  "",
		},
		{
			name:  "Строка из одного символа",
			input: "a",
			want:  "a",
		},
		{
			name:  "Строка-палиндром",
			input: "level",
			want:  "level",
		},
		{
			name:  "Строка-палиндром с кириллицей",
			input: "топот",
			want:  "топот",
		},
		{
			name:  "Четное количество символов",
			input: "abcd",
			want:  "dcba",
		},
		{
			name:  "Нечетное количество символов",
			input: "abcde",
			want:  "edcba",
		},
		{
			name:  "Со специальными символами",
			input: "a\tb\n",
			want:  "\nb\ta",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Flip(tc.input)
			if got != tc.want {
				t.Errorf("Flip(%s) = %s, ожидалось %s.", tc.input, got, tc.want)
			}
		})
	}
}
