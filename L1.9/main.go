package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	move := make(chan int) // Канал для исходных данных
	mult := make(chan int) // Канал для результирующих данных
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go func() { // Запускаем горутину, запись в первый канал
		for i, _ := range array {
			move <- array[i] // Отправка данных в канал 1
		}
		wg.Done()
	}()
	go func() { // Запускаем горутину, обработку данных из первого канала
		for vr := range move { // Чтение из канала 1 
			vr = vr * 2
			mult <- vr // Запись в канал 2
		}
		wg.Done()
	}()

	for i := 0; i < len(array); i++ { // Вывод на экран
		result := <-mult
		fmt.Println(result)
	}

	close(move) // Закрываем каналы
	close(mult)
	wg.Wait()
}
