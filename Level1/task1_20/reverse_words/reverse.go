package reverse_words

func reverseRange(runes []rune, start, end int) {
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}

// Reverse переворачивает порядок слов в строке
func Reverse(s string) string {
	runes := []rune(s)
	// сначала просто перевернем строку
	reverseRange(runes, 0, len(runes)-1)
	// теперь собираем слова (они сейчас написаны задом наперед)
	startWord := 0 // начало очередного слова
	for i := 0; i <= len(runes); i++ {
		// Конец слова найден, если мы дошли до конца строки или встретили пробел
		if i == len(runes) || runes[i] == ' ' {
			reverseRange(runes, startWord, i-1)
			startWord = i + 1
		}
	}
	return string(runes)
}
