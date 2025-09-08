package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	// Значения больше 2^20
	a.SetString("999999999999999999999", 10) // Очень большое число
	b.SetString("123456789012345678901", 10) // Еще одно большое число

	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n", b.String())

	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Printf("Сложение:\n%s + %s = \n%s\n\n", a, b, sum)

	diff := new(big.Int)
	diff.Sub(a, b)
	fmt.Printf("Вычитание:\n%s - %s = \n%s\n\n", a, b, diff)

	mult := new(big.Int)
	mult.Mul(a, b)
	fmt.Printf("Умножение:\n%s * %s = \n%s\n\n", a, b, mult)

	if b.Cmp(big.NewInt(0)) != 0 {
		div := new(big.Int)
		mod := new(big.Int)
		div.DivMod(a, b, mod)
		fmt.Printf("Деление:\n%s / %s = \n%s (остаток: %s)\n\n", a, b, div, mod)

		aFloat := new(big.Float).SetInt(a)
		bFloat := new(big.Float).SetInt(b)
		divFloat := new(big.Float)
		divFloat.Quo(aFloat, bFloat)
		fmt.Printf("Деление (float):\n%s / %s = \n%s\n", a, b, divFloat.Text('f', 10))
	} else {
		fmt.Println("Ошибка: деление на ноль!")
	}
}
