package main

import "fmt"

											
func createHugeString(size int) string {	// Функция для создания строки (эмуляция createHugeString)
	runes := make([]rune, size)
	for i := range runes { // Для корректной работы с символами - руны
		runes[i] = '-' // Заполняем строку символами 'a'
	}
	return string(runes)
}


func someFunc() string { // Работа со строкой
	v := createHugeString(1 << 10) // Создаём большую строку
	if len(v) < 100 {
		return v // Возвращаем строку, если она короче 100 символов
	}
	return string([]rune(v[:100])) // Явно копируем первые 100 символов
}

func main() {
	result := someFunc()
	fmt.Println("Результат:", result)
}

