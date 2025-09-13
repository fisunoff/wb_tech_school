package main

import "fmt"

type User struct {
	Name string
}

// removeWithMemoryClear удаляет элемент и затирает ссылку
func removeWithMemoryClear[T any](s []*T, i int) []*T {
	if i < 0 || i >= len(s) {
		return s
	}
	copy(s[i:], s[i+1:])

	var zero T // Получаем нулевое значение для типа T
	s[len(s)-1] = &zero
	return s[:len(s)-1]
}

func main() {
	users := []*User{
		{Name: "Антон"},
		{Name: "Даша"},
		{Name: "Александр"},
		{Name: "Иван"},
	}

	fmt.Printf("Исходный слайс: len=%d, cap=%d\n", len(users), cap(users))
	for _, u := range users {
		fmt.Printf(" - %s\n", u.Name)
	}

	users = removeWithMemoryClear(users, 2)

	fmt.Printf("\nПосле удаления: len=%d, cap=%d\n", len(users), cap(users))
	for _, u := range users {
		fmt.Printf(" - %s\n", u.Name)
	}
}
