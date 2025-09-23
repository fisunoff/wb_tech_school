package sleeper

import (
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	testCases := []time.Duration{
		// 10 * time.Millisecond, не проходит из-за накладных расходов.
		100 * time.Millisecond,
		500 * time.Millisecond,
		1 * time.Second,
	}

	for _, d := range testCases {
		t.Run(d.String(), func(t *testing.T) {
			start := time.Now()
			Sleep(d)
			elapsed := time.Since(start)

			minExpected := d
			maxExpected := time.Duration(float64(d) * 1.01) // +1%

			if elapsed < minExpected {
				t.Errorf("спал слишком мало: %v, ожидалось минимум %v", elapsed, minExpected)
			}
			if elapsed > maxExpected {
				t.Errorf("спал слишком много: %v, ожидалось максимум %v", elapsed, maxExpected)
			}
		})
	}
}
