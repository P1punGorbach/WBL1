package main

import "fmt"

type Human struct { // 
	name string
	age  int
}

func (h Human) Introduce() {
	fmt.Printf("Hello, im %s, and im %d yo\n", h.name, h.age)
}

type Action struct {
	Human // Реализация встраивания
	actiontype string
}

func (a Action) MakeAction() {
	fmt.Printf("Action type: %s\n", a.actiontype)
}
func main() {
	act := Action{ // Создание структурной переменной
		Human:      Human{name: "Peter", age: 17},
		actiontype: "jump",
	}
	act.Introduce() // Вызов методов для Action и Human
	act.MakeAction()
}
