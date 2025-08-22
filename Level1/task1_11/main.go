package main

import "fmt"

// Intersect находит пересечение двух множеств, лежащих в слайсах
func Intersect(a []int, b []int) []int {
	var short []int
	var long []int
	answer := make([]int, 0)
	if len(a) < len(b) {
		short, long = a, b
	} else {
		short, long = b, a
	}
	hashedShort := make(map[int]bool)
	for _, element := range short {
		hashedShort[element] = true
	}
	for _, element := range long {
		if hashedShort[element] {
			answer = append(answer, element)
		}
	}
	return answer
}

func main() {
	ans := Intersect([]int{1, 2, 3}, []int{2, 3, 4})
	fmt.Println(ans)
}
