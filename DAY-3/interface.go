package main

import "fmt"

type Animal interface {
	speak() string
}

type Dog struct {
	name string
}

func (d Dog) speak() string {
	return d.name + " is speaking"
}

type Cat struct {
	name string
}

func (c Cat) speak() string {
	return " teh cat and meaw"
}

func main() {
	var a Animal

	a = Dog{"dogo"}
	fmt.Println("Dog bark: ", a.speak())

	a = Cat{"tom"}
	fmt.Println("cat meaw: ", a.speak())

}
