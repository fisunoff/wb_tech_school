package quicksort

func partition(arr []int, low, high int) int {
	// Опорный элемент - первый
	pivot := arr[low]
	i := low + 1
	j := high

	for {
		for i <= j && arr[i] <= pivot {
			i++
		}

		for j >= i && arr[j] > pivot {
			j--
		}

		if i > j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}

	arr[low], arr[j] = arr[j], arr[low]
	return j
}

// ByRange реализация алгоритма быстрой сортировки с заданием откуда до куда сортировать
func ByRange(A []int, low int, high int) {
	if low < high {
		p := partition(A, low, high)
		ByRange(A, low, p-1)
		ByRange(A, p+1, high)
	}
}

// QuickSort реализация алгоритма быстрой сортировки
func QuickSort(A []int) []int {
	Copy := make([]int, len(A))
	copy(Copy, A)
	ByRange(Copy, 0, len(A)-1)
	return Copy
}
