package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// Конструктор
func NewPointer(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func DistanceBetween(p1 *Point, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func (p *Point) DistanceTo(pTo *Point) float64 {
	return math.Sqrt(math.Pow(p.x-pTo.x, 2) + math.Pow(p.y-pTo.y, 2))
}

func main() {
	p1 := NewPointer(2, 4)
	p2 := NewPointer(4, 4)

	fmt.Println("Distance between p1 and p2:", DistanceBetween(p1, p2))
	fmt.Println("Distance from p1 to p2:", p1.DistanceTo(p2))
}
