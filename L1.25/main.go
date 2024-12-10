package main

import (
	"fmt"
	"time"
)

func MySleep(duration time.Duration) { // Функция
	timer := time.NewTimer(duration)
	<-timer.C // Блокирует выполнение, пока таймер не истечет, timer.c - канал связанный с таймером
}

func main() {
	fmt.Println("Начало")
	MySleep(2 * time.Second) // Ожидание 2 секунды
	fmt.Println("Конец")
}
