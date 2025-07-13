//  my code is below

// package main 
// import "fmt"


// type ErrorCustom struct{
// 		intA int
// 		intB int
// }

// func (e ErrorCustom) Error(a , b int){
// 	if b == 0 {
// 		fmt.Println("cannot devide by" , b)
// 		return;
// 	}

// 	fmt.Println(a/b);
// }

// func main(){
// 	var error ErrorCustom = ErrorCustom{
// 		intA: 10,
// 		intB: 10,
// 	}

// 	error.Error(error.intA , error.intB)

// }






//  gpt code 
package main

import (
	"fmt"
)

// Step 1: Define custom error type
type DivideError struct {
	A, B int
}

// Step 2: Implement error interface
func (e DivideError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", e.A, e.B)
}

// Step 3: Function that returns custom error
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, DivideError{A: a, B: b}
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
