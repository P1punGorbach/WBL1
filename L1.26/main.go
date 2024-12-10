package main

import (
	"fmt"
	"strings"
)

func allUnique(str string) bool {

	str = strings.ToLower(str) // Приводим к нижнему регистру

	charMap := make(map[rune]bool)

	for _, char := range str {

		if charMap[char] { // Проверка на уникальность
			return false //False если символ уже есть в мапе
		}
		charMap[char] = true //True если
	}

	return true //True если все символы уникальны
}

func main() {

	array := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, value := range array { // Проверяем каждый элемент массива
		fmt.Printf("Строка: %s, %v\n", value, allUnique(value))
	}
}
