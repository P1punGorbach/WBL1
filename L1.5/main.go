package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func Read(c chan float64, quit chan int, wg *sync.WaitGroup) {
	for {
		select {
		case rd := <-c:
			fmt.Printf("Считало: %f \n", rd)

		case <-quit:
			wg.Done()
			return
		}
	}

}
func main() {
	n := 5
	c := make(chan float64)
	quit := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for {
			select {
			case <-quit:
				wg.Done()
				return
			default:

				random := rand.Float64() // Оправка случайных чисел в канал
				c <- random
				fmt.Printf("Отправлено, ")
				time.Sleep(1 * time.Second)
			}

		}

	}()
	go Read(c, quit, &wg)
	time.Sleep(time.Duration(n) * time.Second)
	close(quit)
	wg.Wait()
	close(c)

}
