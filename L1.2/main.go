package main

import (
	"fmt"
	"sync"
)

func Square(i int, v int, wg *sync.WaitGroup) {
	v = v * v
	fmt.Printf("Goroutine number: %d | Squared number: %d\n", i, v)
	wg.Done() // Уменьшаем счётчик по окончании функции
}

func main() {
	var wg sync.WaitGroup  // объявляем WaitGroup для ожидания окончания всех Го рутин
	wg.Add(5) // Группа из 5 элементов
	var numbers = []int{2, 4, 6, 8, 10}
	i := 1
	for _, value := range numbers {
		go Square(i, value, &wg) // Вызов го рутины с передачей адреса переменной wg
		i++
	}
	wg.Wait() // Ожидаем завершения всех го рутин
}
