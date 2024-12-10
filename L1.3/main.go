package main

import (
	"fmt"
	"sync"
)

func Square(ch chan int, v int, wg *sync.WaitGroup) {
	v = v * v
	ch <- v // Отправляем в канал 
	wg.Done() // Уменьшаем счетчик
}

func main() {
	intCh := make(chan int, 5) // Создаем канал емкостью 5
	var wg sync.WaitGroup // Создаём WaitGroup для ожидания завершения всех го-рутин вычисления квадратов
	wg.Add(5) // Группа из 5 элементов
	var numbers = []int{2, 4, 6, 8, 10}

	for _, value := range numbers {
		go Square(intCh, value, &wg) // Вызов го рутины
	}
	wg.Wait()
	close(intCh) // Закрываем канал после обработки всех чисел
	sum := 0
	for i := 0; i < cap(intCh); i++ {
		if val, opened := <-intCh; opened { // Складываем, пока канал открыт
			sum += val
		} else {
			fmt.Println("Channel closed!")
		}

	}
	fmt.Printf("Sum: %d", sum)
}
