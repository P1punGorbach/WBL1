package main

import "fmt"

type Target interface {
	Request() string
}

type Adaptee struct{} // Объект, который нужно адаптировать под интерфейс

func (a *Adaptee) adapteeRequest() string { // Несовместимый с Target метод, который должен работать
	return "Adaptee work"
}

type Adapter struct { // Адаптер, который приводит Adaptee к интерфейсу Target
	adaptee *Adaptee
}

func (a *Adapter) Request() string { // Реализует интерфейс Target (переходник)
	return a.adaptee.adapteeRequest() // Вызывает метод dapteeRequest() и передает его результат
}

func main() {
	adaptee := &Adaptee{}        // Создание объекта Adaptee
	adapter := &Adapter{adaptee} // Создание адаптера, которому передается сслыка на объект Adaptee
	fmt.Println(adapter.Request())

}
