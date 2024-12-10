package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	sum := new(big.Int)
	subt := new(big.Int)
	mult := new(big.Int)
	divis := new(big.Int)

	a.SetString("1048577", 10)
	b.SetString("2097153", 10)

	sum.Add(a, b) // Сложение
	fmt.Println("Summa:", sum)

	subt.Sub(a, b) // Вычитание
	fmt.Println("Subtraction:", sum)

	mult.Mul(a, b) // Умножение
	fmt.Println("Multiplication:", sum)

	divis.Div(a, b) // Деление
	fmt.Println("Division:", sum)
}
