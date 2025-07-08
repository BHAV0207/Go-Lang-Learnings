package main

import "fmt"

func main() {
	var slice []string = []string{"go" , "is" , "fun"};
	fmt.Println(slice);


	slice = append(slice , "language");
	fmt.Println(slice);


	for _ , value := range slice{
		fmt.Println(value);
	}
}

//  the difference between an array and a slice is that an array has a fixed size , while a slice is a dynamically sized  flexible view into the elements of an array.