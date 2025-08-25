package concurrent_counter

import (
	"sync"
	"testing"
)

func TestConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	counter := Counter{}
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
	if value != expected {
		t.Errorf("Ожидалось значение счетчика %d, но получили %d", expected, value)
	}
}
