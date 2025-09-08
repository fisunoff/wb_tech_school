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

func TestSerializeOrderPositive(t *testing.T) {
	correctData := readTestData(t, "valid_order.json")
	_, err := SerializeOrder(correctData)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestSerializeOrderInvalidSyntax(t *testing.T) {
	correctData := readTestData(t, "invalid_syntax.json")
	_, err := SerializeOrder(correctData)
	if err == nil {
		t.Error("Ожидалась ошибка, но ее нет")
	}
}

func TestSerializeOrderInvalidType(t *testing.T) {
	correctData := readTestData(t, "invalid_type.json")
	_, err := SerializeOrder(correctData)
	if err == nil {
		t.Error("Ожидалась ошибка, но ее нет")
	}
}

func TestSerializeOrderInvalidEmail(t *testing.T) {
	data := readTestData(t, "invalid_email.json")
	_, err := SerializeOrder(data)
	if err == nil {
		t.Error("Ожидалась ошибка, но ее нет")
	}
}
