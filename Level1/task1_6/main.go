package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

// randomNot3 - генерировать случайные числа до тех пор, пока не выведется число 3
func randomNot3(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		a := rand.Intn(50)
		if a == 3 {
			break
		}
		println(a)
	}
}

func doneChannel(done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		default:
			println("Работа горутины с выходом через канал уведомления")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func exitByContext(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер: контекст отменен. Завершаюсь. Причина: %v\n", ctx.Err())
			return
		default:
			println("Работа горутины с выходом через контекст")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func withGoExit(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		a := rand.Intn(50)
		if a == 3 {
			runtime.Goexit()
		}
		println(a)
	}
}

// withOsExit - вообще так делать не нужно, но привел для полноты картины
func withOsExit(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		a := rand.Intn(50)
		if a == 3 {
			println("Полное завершение программы")
			os.Exit(0)
		}
		println(a)
	}
}

func main() {
	// Выход по условию
	var wg sync.WaitGroup
	wg.Add(1)
	go randomNot3(&wg)
	wg.Wait()

	// Выход через канал уведомления
	wg.Add(1)
	done := make(chan struct{})
	go doneChannel(done, &wg)
	time.Sleep(time.Second)
	close(done)
	wg.Wait()

	// Выход через контекст
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go exitByContext(&wg, ctx)
	time.Sleep(1 * time.Second)
	cancel()
	wg.Wait()

	// Выход через контекст по времени
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	wg.Add(1)
	println("Тоже выход через контекст, то таймер установлен в контексте")
	go exitByContext(&wg, ctx)
	wg.Wait()

	// Выход через GoExit
	wg.Add(1)
	go withGoExit(&wg)
	wg.Wait()
	println("Произведен выход через через Goexit")

	// Полное завершение программы
	wg.Add(1)
	go withOsExit(&wg)
	wg.Wait()
	println("Этот код не будет напечатан, так как программа завершится из горутины")
}
