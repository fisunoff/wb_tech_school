package main

import (
	"bufio"
	"io"
)

// readLines читает строки из io.Reader в срез.
func readLines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines записывает строки из среза в io.Writer.
func writeLines(w io.Writer, lines []string) error {
	writer := bufio.NewWriter(w)
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return writer.Flush()
}

// getUnique оставляет только уникальные элементы в отсортированном срезе.
func getUnique(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}
	// Используем тот же базовый массив для экономии памяти
	j := 1
	for i := 1; i < len(lines); i++ {
		if lines[i] != lines[i-1] {
			lines[j] = lines[i]
			j++
		}
	}
	return lines[:j]
}
