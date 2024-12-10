package main

import "fmt"

func main() {
	array1 := []int{-25, -27, 13, 19, 29, 52, -21, 13}
	array2 := []int{-29, -27, 52, -21, 13, 77, 14, 9, 13}
	mapa := make(map[int]int)
	func() {
		for i, _ := range array1 {
			// mapa[array1[i]]++
			key := array1[i]
			count := mapa[key] // Сравнение ключей
			count++            // Изменение счетчика количества появлений
			mapa[key] = count

		}

		for j, _ := range array2 {
			key := array2[j]
			count := mapa[key] // Проверяем количество вхождений элемента

			if count != 0 { // Если элемент найден в `array1`
				count--                // Уменьшаем счётчик
				fmt.Println(array2[j]) // Печатаем пересекающийся элемент
			}

			mapa[key] = count // Обновляем значение в карте
		}

	}()

}
