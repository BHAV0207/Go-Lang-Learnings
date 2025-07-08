package main

import "fmt"

func main() {
	const pi = 3.14
	var str string = "hello"
	// if we write var str string ; then it will be assigned a default value of "" which is called zero value in go lang;

	num := 42
	// num := 42 is a shorthand for declaring and initializing a variable in Go.
	// it can only be used inside a function, and it infers the type of the variable from the value assigned to it.

	fmt.Println("Pi: ", pi)
	fmt.Println("String: ", str)
	fmt.Print("Number: ", num, "\n")
}

// in this code there are few things that you need to take care of :
// if we write const pi = 3.14 then it will be a constant and it cannot be changed
