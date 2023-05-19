//Yeah, tried to make the shape a receiver but looks like interfaces cannot be receivers
//Here's the final code anyways

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
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("the area is:", s.getArea())
}

func main() {
	s := square{sideLength: 2}
	t := triangle{height: 10, base: 5}
	printArea(s)
	printArea(t)
}
