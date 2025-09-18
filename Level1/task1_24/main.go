package main

import (
	"fmt"
	"task1_24/point"
)

func main() {
	p1 := point.NewPoint(0, 0)
	p2 := point.NewPoint(3, 4)

	distance := p1.Distance(p2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
