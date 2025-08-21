package main

import (
	"math/rand"
	"sync"
)

var mu = sync.Mutex{}

func randomWriteToMap(wg *sync.WaitGroup, m map[int]int) {
	defer wg.Done()
	for range 100_000 {
		// Чтобы держать лок поменьше, сначала делаю то, для чего лок не нужен, а только потом его ставлю
		newKey := rand.Intn(1000)
		newVal := rand.Intn(1000)
		mu.Lock()
		m[newKey] = newVal
		mu.Unlock()
	}
}

func main() {
	m := map[int]int{}
	wg := sync.WaitGroup{}
	for range 50 {
		wg.Add(1)
		go randomWriteToMap(&wg, m)
	}
	wg.Wait()
}
