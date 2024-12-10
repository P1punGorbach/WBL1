package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Инициализация среза
	fmt.Println("Slice before delete: ", slice)

	var n int
	fmt.Println("Enter index to delete: ")
	_, err := fmt.Scan(&n) // Инициализация индекса i из ввода пользователя

	if err != nil {
		fmt.Println(" Wrong number")
		return
	}
	if n < 0 || n >= len(slice) {
		fmt.Println("Out of bounds!!")
		return
	}
	slice = append(slice[:n], slice[n+1:]...) // Удаление элемента из среза

	fmt.Println("Changed slice: ", slice)
}
