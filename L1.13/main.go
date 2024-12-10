package main

import (
	"fmt"
	"time"
)

func main() {
	a := 5
	b := 10
	fmt.Printf("Значения A и B до обмена: %d, %d\n", a, b)
	a, b = b, a
	time.Sleep(2 * time.Second)
	fmt.Printf("Значения A и B после обмена: %d, %d\n", a, b)
}
