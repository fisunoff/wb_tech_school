package main

import "fmt"

func SplitByTens(temperatures []float64) map[int][]float64 {
	res := make(map[int][]float64)
	for _, temperature := range temperatures {
		// В задании ничего не было сказано про то, как интерпретировать числа в диапазоне (-10; 10)
		// этот диапазон, я решил взять как 0, к нему относится всё в пределах (-10; 10).
		// При этом продолжает подходить общая формула.
		scope := int(temperature/10) * 10 // чтобы оставить число с округлением в сторону 0 до десятков
		res[scope] = append(res[scope], temperature)
	}
	return res
}

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -5, 5, 0}
	temperatures := SplitByTens(data)
	fmt.Println(temperatures)
}
