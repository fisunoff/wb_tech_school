package quicksort

import (
	"reflect"
	"testing"
)

type testCase struct {
	name     string
	input    []int
	expected []int
}

func TestQuickSort(t *testing.T) {
	testCases := []testCase{
		{
			name:     "Стандартный случай",
			input:    []int{10, 7, 8, 9, 1, 5},
			expected: []int{1, 5, 7, 8, 9, 10},
		},
		{
			name:     "Пустой слайс",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Слайс из одного элемента",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Уже отсортировано",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Отсортировано в обратном порядке",
			input:    []int{6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "С дублирующимися элементами",
			input:    []int{5, 2, 8, 2, 5, 1, 8},
			expected: []int{1, 2, 2, 5, 5, 8, 8},
		},
		{
			name:     "С отрицательными числами",
			input:    []int{-10, 5, -20, 0, 15},
			expected: []int{-20, -10, 0, 5, 15},
		},
		{
			name:     "Со всеми одинаковыми элементами",
			input:    []int{5, 5, 5, 5, 5},
			expected: []int{5, 5, 5, 5, 5},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sortedData := QuickSort(testCase.input)

			if !reflect.DeepEqual(sortedData, testCase.expected) {
				t.Errorf("QuickSort() = %v, ожидалось %v", sortedData, testCase.expected)
			}
		})
	}
}
