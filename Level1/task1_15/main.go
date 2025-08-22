package main

import "strings"

var justString string

func someFunc() {
	// В переменной v длинная строка
	v := createHugeString(1 << 10)
	// Так как justString это слайс от v, то сборщик мусора не будет удалять v, пока есть justString.
	// По сути, у нас много памяти уходит в никуда на хранение уже не очень то нужной v целиком.
	justString = v[:100]
}

// someFuncGood Исправленный вариант
func someFuncGood() {
	v := createHugeString(1 << 10)
	// Теперь justString это самостоятельная строка, которая не завязана на v.
	// При выходе из функции someFuncGood сборщик мусора освободит память, занятую под переменную v.
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
	someFuncGood()
}
