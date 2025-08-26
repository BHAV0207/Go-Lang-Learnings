package main

import "fmt"

type Car struct {
	Brand string
	Year  int
}

func (c Car) Details() {
	fmt.Println(c.Brand)
	fmt.Println(c.Year)
}

func main() {
	car1 := Car{"Tesla", 2023}
	car2 := Car{"ford", 2011}

	car1.Details()
	car2.Details()
}
