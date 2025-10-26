package anagram

import (
	"slices"
	"strings"
)

// Find - ищет анаграммы
func Find(words []string) map[string][]string {
	// Будем сортировать буквы в слове по алфавиту. Ключ во внутреннем словаре - отсортированные буквы.
	// Потом уже из него будем составлять словарь как просят по заданию, в котором ключ - первое встретившееся слово.
	rawWords := make(map[string][]string)
	for _, word := range words {
		lowerWord := strings.ToLower(word)
		symbols := []rune(lowerWord)
		slices.Sort(symbols)
		key := string(symbols)
		rawWords[key] = append(rawWords[key], word)
	}

	ans := make(map[string][]string)
	for _, anagramWords := range rawWords {
		if len(anagramWords) > 1 {
			key := anagramWords[0]
			slices.Sort(anagramWords)
			ans[key] = anagramWords
		}
	}
	return ans
}
