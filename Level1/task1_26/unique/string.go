package unique

import "unicode"

// CheckString регистронезависимо проверяет, есть ли в строке повторы рун
func CheckString(s string) bool {
	usedRunes := map[rune]bool{}
	for _, r := range s {
		lowerRune := unicode.ToLower(r)
		if usedRunes[lowerRune] {
			return false
		}
		usedRunes[lowerRune] = true
	}
	return true
}
