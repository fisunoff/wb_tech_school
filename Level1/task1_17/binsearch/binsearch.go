package binsearch

import "cmp"

// Search выполняет бинарный поиск значения key в отсортированном срезе arr.
// Если элемент найден, возвращается его индекс.
// Если элемент не найден, возвращается -1.
func Search[T cmp.Ordered](arr []T, key T) int {
	// cmp.Ordered выставлен чтобы типы поддерживали операции сравнения
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if key == arr[mid] {
			return mid
		}
		if key > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
