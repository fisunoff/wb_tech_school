package reverse_words

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	type testCase struct {
		name  string
		input string
		want  string
	}

	tests := []testCase{
		{
			name:  "Стандартный случай",
			input: "snow dog sun",
			want:  "sun dog snow",
		},
		{
			name:  "Более длинное предложение",
			input: "the quick brown fox jumps over the lazy dog",
			want:  "dog lazy the over jumps fox brown quick the",
		},
		{
			name:  "Пустая строка",
			input: "",
			want:  "",
		},
		{
			name:  "Одно слово",
			input: "hello",
			want:  "hello",
		},
		{
			name:  "Строка из одного пробела",
			input: " ",
			want:  " ",
		},
		{
			name:  "Кириллица",
			input: "раз два три",
			want:  "три два раз",
		},
		{
			name:  "Слова со знаками препинания",
			input: "first-word, second. third!",
			want:  "third! second. first-word,",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Reverse(tc.input)
			if got != tc.want {
				t.Errorf("reverseWords(%q) = %q; ожидалось %q", tc.input, got, tc.want)
			}
		})
	}
}
