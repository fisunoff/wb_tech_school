package model

import (
	"os"
	"path/filepath"
	"testing"
)

func readTestData(t *testing.T, filename string) []byte {
	t.Helper()
	path := filepath.Join("testdata", filename)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Не удалось прочитать тестовый файл %s: %v", filename, err)
	}
	return data
}

func TestParseOrderFromJSONPositive(t *testing.T) {
	correctData := readTestData(t, "valid_order.json")
	_, err := ParseOrderFromJSON(correctData)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestParseOrderFromJSONInvalidSyntax(t *testing.T) {
	correctData := readTestData(t, "invalid_syntax.json")
	_, err := ParseOrderFromJSON(correctData)
	if err == nil {
		t.Error("Ожидалась ошибка, но ее нет")
	}
}

func TestParseOrderFromJSONInvalidType(t *testing.T) {
	correctData := readTestData(t, "invalid_type.json")
	_, err := ParseOrderFromJSON(correctData)
	if err == nil {
		t.Error("Ожидалась ошибка, но ее нет")
	}
}

func TestParseOrderFromJSONInvalidEmail(t *testing.T) {
	data := readTestData(t, "invalid_email.json")
	_, err := ParseOrderFromJSON(data)
	if err == nil {
		t.Error("Ожидалась ошибка, но ее нет")
	}
}
