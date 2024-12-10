package main

import (
	"fmt"
	"strings"
)

func reverseString(input string) string {

	words := strings.Fields(input) // Создаем срез из слов

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 { // Меняем соответствующие слова
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ") // Возвращем строку, созданную из среза
}

func main() {

	line := "snow dog sun"
	fmt.Println("Начальная строка:", line)
	reversed := reverseString(line)
	fmt.Println("Перевёрнутая строка:", reversed)
}
