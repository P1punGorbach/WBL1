package main

import (
	"fmt"
	"math"
)

func main() {
	var group float64
	wordMap := make(map[float64][]float64) // Мапа, которая хранит в себе ключи и подмножества
	array := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for i, _ := range array {
		group = array[i] - math.Mod(array[i], 10) // Определение наименования группы
		wordMap[group] = append(wordMap[group], array[i])
	}
	for key, value := range wordMap {
		fmt.Println(key, value)
	}

}
