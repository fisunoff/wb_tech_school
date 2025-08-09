package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// Env - получить данные из окружения
func Env(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}

// MustAtoi - привести к числу
func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("bad int value %q: %v", s, err)
	}
	return n
}

// SplitAndTrim - получить список строк из строки по разделителю ,
func SplitAndTrim(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
