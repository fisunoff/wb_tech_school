package flipper

// Flip перевернуть Unicode строку. Сделано допущение, что каждый символ строки - один элемент из Unicode.
func Flip(s string) string {
	// Для корректной обработки unicode символов надо сначала все преобразовать в руны
	runes := []rune(s)
	for i, tail := 0, len(runes)-1; i < tail; i, tail = i+1, tail-1 {
		runes[i], runes[tail] = runes[tail], runes[i]
	}
	return string(runes)
}
