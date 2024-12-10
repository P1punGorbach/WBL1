package main

import (
	"fmt"
)

func reverseString(input string) string {

	runes := []rune(input) // Преобразуем строку в массив рун

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { // Меняем местами соответствующие элементы
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes) // Возвращаем массив рун обратно в строку
}

func main() {
	var input string
	fmt.Println("Введите строку для переворота:")
	fmt.Scanln(&input)

	reversed := reverseString(input) // Переворот строки

	fmt.Println("Перевёрнутая строка:", reversed)
}
