package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Job struct {
	index int
	value int
}

var workerCount = runtime.NumCPU() // Не больше количества ядер

func work(jobs <-chan Job, result chan<- Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		result <- Job{j.index, j.value * j.value}
	}
}

func slicePow2(data []int) []int {
	cnt := len(data)
	result := make([]int, cnt)
	input := make(chan Job, cnt)
	output := make(chan Job, cnt)

	var wg sync.WaitGroup

	for worker := 1; worker <= workerCount; worker++ {
		wg.Add(1)
		go work(input, output, &wg)
	}

	for index, element := range data {
		input <- Job{index, element}
	}
	close(input)

	go func() {
		wg.Wait()
		close(output)
	}()

	for jobResult := range output {
		result[jobResult.index] = jobResult.value
	}
	return result
}

func main() {
	res := slicePow2([]int{2, 4, 6, 8, 10})
	fmt.Println(res)
}
