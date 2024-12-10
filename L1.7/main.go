package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(2)
	velikayamapa := make(map[int]int)
	go func() {
		mutex.Lock() // Блокирует mutex
		velikayamapa[0] = 1
		mutex.Unlock() // Разблокирует mutex
		wg.Done() // Уменьшаем счетчик го рутин
	}()

	go func() {
		mutex.Lock() // Блокирует mutex
		velikayamapa[0] = 3
		mutex.Unlock() // Разблокирует mutex
		wg.Done()
	}()
	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println("The value:", velikayamapa[0])
}
