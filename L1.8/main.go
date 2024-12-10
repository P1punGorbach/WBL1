package main

import (
	"fmt"
)

func setBit(num int64, i int, value int) int64 {
	if value == 1 { // Инвертируем нужный нам бит, 0 меняем на 1 или 1 меняем на 0
		return num | (1 << i)
	} else {
		return num & ^(1 << i)
	}
}

func main() {
	var num int64
	var i, value int

	fmt.Print("Enter number: ")
	fmt.Scan(&num)

	fmt.Print("Enter bit number: ")
	fmt.Scan(&i)

	fmt.Print("Enter new bit: ")
	fmt.Scan(&value)

	result := setBit(num, i, value)
	fmt.Printf("Result: %d\n", result)
}
