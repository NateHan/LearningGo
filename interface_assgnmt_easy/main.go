package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func main() {
	sq := square{5.5}
	tr := triangle{3.7, 4.6}

	printArea(sq)
	printArea(tr)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(sh shape) {
	s := fmt.Sprint(sh.getArea())
	fmt.Printf("The area of the shape is: %v .\n", s)
}
