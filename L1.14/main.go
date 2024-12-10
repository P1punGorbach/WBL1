package main

import "fmt"

func Checker(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Println("Type int:\n", v)
	case string:
		fmt.Println("Type string:\n", v)
	case bool:
		fmt.Println("Type bool:\n", v)
	case chan int:
		fmt.Println("Type channel:\n", v)
	default:
		fmt.Println("Cant check the type\n")
	}
}
func main() {
	var temp interface{}
	temp = 42
	Checker(temp)
	temp = "Hello, world!"
	Checker(temp)
	temp = true
	Checker(temp)
	temp = make(chan int)
	Checker(temp)
}
