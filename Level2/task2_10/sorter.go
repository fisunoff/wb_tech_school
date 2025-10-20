package main

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// run выполняет основную логику: чтение, сортировку/проверку, вывод.
func run(reader io.Reader, writer io.Writer, inputName string, config *Config) error {
	lines, err := readLines(reader)
	if err != nil {
		return fmt.Errorf("ошибка чтения строк: %w", err)
	}

	sorter := &lineSorter{
		lines:  lines,
		config: config,
	}

	if config.CheckSorted {
		for i := 1; i < len(lines); i++ {
			if sorter.Less(i, i-1) {
				return fmt.Errorf("sort: %s:%d: disorder: %s", inputName, i+1, lines[i])
			}
		}
		return nil
	}

	sort.Sort(sorter)

	if config.Unique {
		lines = getUnique(lines)
	}

	return writeLines(writer, lines)
}

// lineSorter - реализация sort.Interface для сортировки строк с учетом флагов.
type lineSorter struct {
	lines  []string
	config *Config
}

func (s *lineSorter) Len() int      { return len(s.lines) }
func (s *lineSorter) Swap(i, j int) { s.lines[i], s.lines[j] = s.lines[j], s.lines[i] }

func (s *lineSorter) Less(i, j int) bool {
	lineA, lineB := s.lines[i], s.lines[j]
	valA := getValue(lineA, s.config)
	valB := getValue(lineB, s.config)

	var less bool
	switch {
	case s.config.Numeric:
		numA, _ := strconv.ParseFloat(strings.TrimSpace(valA), 64)
		numB, _ := strconv.ParseFloat(strings.TrimSpace(valB), 64)
		less = numA < numB
	case s.config.HumanNumeric:
		numA, _ := parseHumanNumeric(valA)
		numB, _ := parseHumanNumeric(valB)
		less = numA < numB
	case s.config.Month:
		monthA := parseMonth(valA)
		monthB := parseMonth(valB)
		less = monthA < monthB
	default:
		less = valA < valB
	}

	if s.config.Reverse {
		return !less
	}
	return less
}

// getValue извлекает из строки часть для сравнения.
func getValue(line string, config *Config) string {
	var val string
	if config.Column > 0 {
		parts := strings.Split(line, "\t")
		if config.Column-1 < len(parts) {
			val = parts[config.Column-1]
		} else {
			val = ""
		}
	} else {
		val = line
	}

	if config.IgnoreBlanks {
		val = strings.TrimRight(val, " \t")
	}

	return val
}

// Хелперы для парсинга

var monthMap = map[string]int{
	"jan": 1, "feb": 2, "mar": 3, "apr": 4, "may": 5, "jun": 6,
	"jul": 7, "aug": 8, "sep": 9, "oct": 10, "nov": 11, "dec": 12,
}

func parseMonth(s string) int {
	s = strings.TrimSpace(s)
	if len(s) < 3 {
		return 0
	}
	monthPrefix := strings.ToLower(s[:3])
	if val, ok := monthMap[monthPrefix]; ok {
		return val
	}
	return 0
}

func parseHumanNumeric(s string) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, nil
	}

	lastChar := s[len(s)-1]
	multiplier := int64(1)
	numPart := s

	switch unicode.ToUpper(rune(lastChar)) {
	case 'K':
		multiplier = 1024
		numPart = s[:len(s)-1]
	case 'M':
		multiplier = 1024 * 1024
		numPart = s[:len(s)-1]
	case 'G':
		multiplier = 1024 * 1024 * 1024
		numPart = s[:len(s)-1]
	case 'T':
		multiplier = 1024 * 1024 * 1024 * 1024
		numPart = s[:len(s)-1]
	}

	val, err := strconv.ParseFloat(numPart, 64)
	if err != nil {
		return 0, err
	}

	return int64(val * float64(multiplier)), nil
}
