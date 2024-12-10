package main

import "fmt"

func main() {
	result := make(map[string]struct{})
	array := []string{"cat", "cat", "dog", "cat", "tree"}

	for i, _ := range array {
		result[array[i]] = struct{}{} // Заполняем мапу
	}
	for key, _ := range result {
		fmt.Println(key) // Выводим индексы мапы
	}

}
