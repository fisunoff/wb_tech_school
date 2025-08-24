package binsearch

import (
	"cmp"
	"testing"
)

// runSearchTests — generic-хелпер для запуска тестов.
// Он принимает срез тестовых случаев для любого сравниваемого типа T.
func runSearchTests[T cmp.Ordered](t *testing.T, testCases []struct {
	name     string
	arr      []T
	key      T
	expected int
}) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Search(tc.arr, tc.key)
			if result != tc.expected {
				t.Errorf("Search(%v, %v) = %d; want %d", tc.arr, tc.key, result, tc.expected)
			}
		})
	}
}

// TestSearchInts - тест для int.
func TestSearchInts(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		key      int
		expected int
	}{
		{"элемент найден в середине", []int{1, 3, 5, 7, 9}, 5, 2},
		{"элемент найден в конце", []int{1, 3, 5, 7, 9}, 9, 4},
		{"элемент не найден", []int{1, 3, 5, 7, 9}, 6, -1},
		{"пустой срез", []int{}, 5, -1},
	}
	runSearchTests(t, testCases)
}

// TestSearchFloat64s - тест для float64.
func TestSearchFloat64s(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []float64
		key      float64
		expected int
	}{
		{"float найден", []float64{1.1, 2.2, 3.3, 4.4}, 3.3, 2},
		{"float не найден", []float64{1.1, 2.2, 3.3, 4.4}, 3.0, -1},
		{"отрицательные числа", []float64{-5.5, -3.1, 0, 1.2}, -5.5, 0},
	}
	runSearchTests(t, testCases)
}

// TestSearchStrings - тест для string.
func TestSearchStrings(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []string
		key      string
		expected int
	}{
		{"строка найдена", []string{"a", "b", "c", "d"}, "c", 2},
		{"строка в начале", []string{"apple", "banana", "cherry"}, "apple", 0},
		{"строка не найдена", []string{"apple", "banana", "cherry"}, "apricot", -1},
		{"пустая строка", []string{"", "a", "b"}, "", 0},
	}
	runSearchTests(t, testCases)
}
