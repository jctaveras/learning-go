package main

import (
	"fmt"
	"math"
	"strings"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	unitSquare := square{sideLength: 1.0}
	unitTriangle := triangle{base: 1.0, height: 1.0}

	fmt.Println(printArea(unitSquare))
	fmt.Println(printArea(unitTriangle))
}

func printArea(s shape) string {
	return strings.Join([]string{"Area:", fmt.Sprint(s.getArea())}, " ")
}

func (s square) getArea() float64 {
	return math.Pow(2, s.sideLength)
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}
