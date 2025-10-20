package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	config := parseFlags()

	var reader io.Reader
	var inputFileName string

	if flag.NArg() > 0 {
		inputFileName = flag.Arg(0)
		file, err := os.Open(inputFileName)
		if err != nil {
			log.Fatalf("Ошибка открытия файла %s: %v", inputFileName, err)
		}
		defer func() {
			err = file.Close()
			if err != nil {
				println("Error while closing file")
			}
		}()
		reader = file
	} else {
		reader = os.Stdin
		inputFileName = "stdin"
	}

	err := run(reader, os.Stdout, inputFileName, config)
	if err != nil {
		if config.CheckSorted {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		log.Fatalf("Ошибка выполнения: %v", err)
	}
}
