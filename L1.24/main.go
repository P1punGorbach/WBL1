package main

import (
	"fmt"
	"math"
)

type Point struct { // Собственная структура
	x float64
	y float64
}

func NewPoint(x, y float64) *Point { // Конструктор
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) GetX() float64 { // Геттеры для реализации принципов инкапсуляции
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func Distance(p1, p2 *Point) float64 { // Вычисление растояния между точками
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(3, 4)
	p2 := NewPoint(7, 8)
	dist := Distance(p1, p2)
	fmt.Printf("Расстояние: %f\n", dist)
}
