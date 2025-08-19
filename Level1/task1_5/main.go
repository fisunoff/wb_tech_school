package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func work(jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for number := range jobs {
		println(number)
	}
}

func producer(ctx context.Context, jobs chan<- int) {
	defer close(jobs) // когда закроется канал, цикл в work тоже прекратится

	for {
		select {
		case <-ctx.Done():
			return
		default:
			jobs <- rand.Int()
		}
	}
}

func main() {
	input := make(chan int)

	workerCountPtr := flag.Int("workers", runtime.NumCPU(), "Количество горутин-воркеров для запуска")
	secondsWorkPtr := flag.Int("time", 10, "Время работы в секундах")
	flag.Parse()
	workerCount := *workerCountPtr
	secondsWork := time.Duration(*secondsWorkPtr) * time.Second

	var wg sync.WaitGroup

	fmt.Printf("Запуск с %d воркерами...\n", workerCount)

	ctx, cancel := context.WithTimeout(context.Background(), secondsWork)
	defer cancel()

	go producer(ctx, input)

	for worker := 1; worker <= workerCount; worker++ {
		wg.Add(1)
		go work(input, &wg)
	}
	wg.Wait()
}
