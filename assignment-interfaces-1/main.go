package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Triangle struct {
	base   float64
	height float64
}
type Square struct {
	sideLength float64
}

func (t *Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s *Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func getArea(s Shape) float64 {
	return s.getArea()
}

func main() {
	triangle := Triangle{
		base:   4,
		height: 5,
	}
	square := Square{
		sideLength: 10,
	}

	fmt.Printf("Triangle's area = %v\n", getArea(&triangle))
	fmt.Printf("Square's area = %v\n", getArea(&square))
}
