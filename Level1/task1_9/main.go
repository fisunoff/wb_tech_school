package main

import (
	"fmt"
	"sync"
)

func moveDataToRawData(wg *sync.WaitGroup, data []int, rawData chan<- int) {
	defer wg.Done()
	for _, value := range data {
		rawData <- value
	}
	close(rawData)
}

func processData(wg *sync.WaitGroup, rawData <-chan int, processedData chan<- int) {
	defer wg.Done()
	for value := range rawData {
		processedData <- value * 2
	}
	close(processedData)
}

func printData(wg *sync.WaitGroup, processedData <-chan int) {
	defer wg.Done()
	for value := range processedData {
		fmt.Println(value)
	}
}

func main() {
	data := []int{2, 4, 6, 8, 10}

	rawData := make(chan int)
	processedData := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	go moveDataToRawData(&wg, data, rawData)
	wg.Add(1)
	go processData(&wg, rawData, processedData)
	wg.Add(1)
	go printData(&wg, processedData)
	wg.Wait()
}
