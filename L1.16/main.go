package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{42, 17, 89, 23, 56, 78, 11, 34, 67, 90}
	fmt.Println("Unsorted array: ", numbers)

	sort.Ints(numbers) // Сортировка массива
	fmt.Println("Sorted array: ", numbers)
}
