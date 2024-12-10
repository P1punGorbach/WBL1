package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
)

func Read(n int, ch chan float64, wg *sync.WaitGroup, done chan float64) {

	for i := 0; i < n; i++ {
		go func() {
			for {
				select {
				case data := <-ch: // Если есть данные в канале ch, выводим их
					fmt.Println(data)
				case <-done: // Если пришёл сигнал завершения из done, выходим из горутины
					return
				}

			}

		}()
		if _, opened := <-done; opened { // Проверяем, открыт ли канал done. Если закрыт, выходим
		} else {
			return
		}
	}
}
func Write(ch chan float64, done chan float64) {

	for {

		r := rand.Float64()
		select {
		case ch <- r: //
		case <-done: // Если пришел сигнал завершения, выходим из горутины
			return
		}

	}

}

func main() {
	var wg sync.WaitGroup

	n := 3
	ch := make(chan float64)   // Канал для записи данных
	done := make(chan float64) // Сигнальный канал
	go Write(ch, done)
	go Read(n, ch, &wg, done)
	wg.Wait()
	c := make(chan os.Signal, 0)   // Канал для перехвата сигнала CTRL + C
	signal.Notify(c, os.Interrupt) // Подписываемся на сигнал

	<-c // Блокировка пока не получим сигнал
	fmt.Println("dead")
	close(done) // Закрываем канал done, чтобы завершить горутины

}
