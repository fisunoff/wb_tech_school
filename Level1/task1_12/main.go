package main

import "fmt"

// MakeUniqueSlice оставляет в слайсе только уникальные строки
func MakeUniqueSlice(elems []string) []string {
	// Получаем уникальное значения в map
	hashed := make(map[string]struct{})
	for _, elem := range elems {
		hashed[elem] = struct{}{}
	}

	answer := make([]string, 0, len(hashed))
	for elem := range hashed {
		answer = append(answer, elem)
	}
	return answer
}

func main() {
	testData := []string{"cat", "cat", "dog", "cat", "tree"}
	answer := MakeUniqueSlice(testData)
	fmt.Println(answer)
}
