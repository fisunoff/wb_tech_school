package main

import (
	"fmt"
	"sync"
	"task1_18/concurrent_counter"
)

func main() {
	var wg sync.WaitGroup
	counter := concurrent_counter.Counter{}
	numGoroutines := 1000
	numIncrements := 1000
	expected := uint64(numGoroutines * numIncrements)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numIncrements; j++ {
				counter.Inc()
			}
		}()
	}
	wg.Wait()
	value := counter.GetValue()
	fmt.Printf("Ожидалось значение счетчика %d, получили %d", expected, value)
}
