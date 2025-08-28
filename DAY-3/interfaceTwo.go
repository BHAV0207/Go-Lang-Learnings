package main

import "fmt"

type Shape interface {
	area() float64
}

type Circle struct {
	Radius float64
}

type Square struct {
	Height float64
	Width  float64
}

func (c Circle) area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (s Square) area() float64 {
	return s.Height * s.Width
}

func main() {
	var s Shape

	s = Circle{5.6434}
	fmt.Println(s.area())
	s = Square{10.0, 10.0}
	fmt.Print(s.area())

}
