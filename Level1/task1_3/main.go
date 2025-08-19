package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

func work(jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for number := range jobs {
		println(number)
	}
}

func main() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM) // SIGTERM не просили, но добавил чтобы всегда корректно завершалось

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

	go func() {
		for {
			input <- rand.Int()
		}
	}()

	<-sigChan
	close(input)
	fmt.Print("Получен сигнал завершения. Закрываемся.")
}
