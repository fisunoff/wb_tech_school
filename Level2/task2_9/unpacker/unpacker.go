package unpacker

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// Unpack - распаковывает строку с учетом escape - последовательностей.
func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	var result strings.Builder
	runes := []rune(s)
	i := 0

	for i < len(runes) {
		charToRepeat := runes[i]

		if charToRepeat == '\\' {
			if i+1 >= len(runes) {
				return "", errors.New("invalid string: escape character at the end")
			}
			// Символ для повторения — это тот, что идет после слэша
			charToRepeat = runes[i+1]
			i += 2
		} else {
			// Это обычный символ. Проверяем, не является ли он цифрой.
			// Самостоятельная цифра (без экранирования) не может быть символом для повторения.
			if unicode.IsDigit(charToRepeat) {
				return "", errors.New("invalid string: standalone digit")
			}
			i++
		}

		// Определяем количество повторений
		var repeatCount = 1

		if i < len(runes) && unicode.IsDigit(runes[i]) {
			numStr := ""
			for ; i < len(runes) && unicode.IsDigit(runes[i]); i++ {
				numStr += string(runes[i])
			}

			var err error
			repeatCount, err = strconv.Atoi(numStr)
			if err != nil {
				return "", err
			}
		}
		result.WriteString(strings.Repeat(string(charToRepeat), repeatCount))
	}
	return result.String(), nil
}
