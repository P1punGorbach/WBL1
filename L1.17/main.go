package main
import (
	"fmt"
	"sort"
)

func main() {
	
	numbers := []int{42, 17, 89, 23, 56, 78, 11, 34, 67, 90}
	sort.Ints(numbers) // Сортировка массива
	
	target := 56 

	index := sort.Search(len(numbers), func(i int) bool { // Используем sort.Search для индекса искомого элемента 
		return numbers[i] >= target // true, если numbers[i] >= target false, если numbers[i] < target
	})


	if index < len(numbers) && numbers[index] == target { // Сравнение с искомым элементом
		fmt.Printf("Элемент %d найден на индексе %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
