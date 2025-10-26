package anagram

import (
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	type testCase struct {
		name  string
		input []string
		want  map[string][]string
	}

	tests := []testCase{
		{
			name:  "Основной случай из задания",
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"},
			want: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:  "Пустой входной срез",
			input: []string{},
			want:  map[string][]string{},
		},
		{
			name:  "Nil входной срез",
			input: nil,
			want:  map[string][]string{},
		},
		{
			name:  "Нет анаграмм",
			input: []string{"один", "два", "три", "четыре"},
			want:  map[string][]string{},
		},
		{
			name:  "Проверка нечувствительности к регистру",
			input: []string{"Пятак", "пятка", "ТЯПКА"},
			want: map[string][]string{
				// Ключом должно быть первое встреченное слово - "Пятак"
				"Пятак": {"Пятак", "ТЯПКА", "пятка"},
			},
		},
		{
			name:  "Слова с дублирующимися буквами",
			input: []string{"колокол", "клоокол"},
			want: map[string][]string{
				"колокол": {"клоокол", "колокол"},
			},
		},
		{
			name:  "Наличие дубликатов слов в исходном срезе",
			input: []string{"кот", "ток", "кот"},
			want: map[string][]string{
				"кот": {"кот", "кот", "ток"},
			},
		},
		{
			name:  "Слова разной длины, но с одинаковым набором букв (не анаграммы)",
			input: []string{"пост", "стоп", "постт"},
			want: map[string][]string{
				"пост": {"пост", "стоп"},
			},
		},
		{
			name:  "Граничный случай с пустыми строками",
			input: []string{"", "a", ""},
			want: map[string][]string{
				// Пустые строки являются анаграммами друг друга
				"": {"", ""},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Find(tc.input)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\nFind(%v) \nполучили: %#v, \nожидали:  %#v", tc.input, got, tc.want)
			}
		})
	}
}
