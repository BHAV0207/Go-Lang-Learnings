package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	s := Rectangle{Width: 4, Height: 6}
	fmt.Println("Area of Rectangle:", s.Area())
}

