package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func work(jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for number := range jobs {
		println(number)
	}
}

func main() {
	workerCountPtr := flag.Int("workers", runtime.NumCPU(), "Количество горутин-воркеров для запуска")
	flag.Parse()
	workerCount := *workerCountPtr
	input := make(chan int)

	var wg sync.WaitGroup

	fmt.Printf("Запуск с %d воркерами...\n", workerCount)

	for worker := 1; worker <= workerCount; worker++ {
		wg.Add(1)
		go work(input, &wg)
	}

	for {
		input <- rand.Int()
	}
}
