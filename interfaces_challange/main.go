package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLenth float64
}

func main() {
	tri := triangle{
		height: 10.0,
		base:   15.0,
	}

	squ := square{
		sideLenth: 20,
	}

	printArea(tri)
	printArea(squ)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLenth * s.sideLenth
}
