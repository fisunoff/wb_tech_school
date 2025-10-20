package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Config хранит все параметры запуска программы.
type Config struct {
	Column       int  // -k: номер колонки для сортировки
	Numeric      bool // -n: сортировать по числовому значению
	Reverse      bool // -r: сортировать в обратном порядке
	Unique       bool // -u: выводить только уникальные строки
	Month        bool // -M: сортировать по названию месяца
	IgnoreBlanks bool // -b: игнорировать хвостовые пробелы
	CheckSorted  bool // -c: проверить, отсортированы ли данные
	HumanNumeric bool // -h: сортировать с учётом числовых суффиксов (K, M, G)
}

// parseFlags парсит флаги командной строки и возвращает структуру Config.
func parseFlags() *Config {
	config := &Config{}

	flag.IntVar(&config.Column, "k", 0, "сортировать по столбцу N")
	flag.BoolVar(&config.Numeric, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&config.Reverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&config.Unique, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&config.Month, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&config.IgnoreBlanks, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&config.CheckSorted, "c", false, "проверить, отсортированы ли данные")
	flag.BoolVar(&config.HumanNumeric, "h", false, "сортировать по числовому значению с учётом суффиксов (K, M, G)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Использование: %s [ФЛАГИ] [ФАЙЛ]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Сортирует строки из ФАЙЛА или стандартного ввода.")
		fmt.Fprintln(os.Stderr, "Флаги:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if config.Column < 0 {
		log.Fatalf("неверный номер столбца: %d", config.Column)
	}

	return config
}
