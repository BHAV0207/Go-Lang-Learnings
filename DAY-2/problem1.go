package main

import "fmt"

func main() {
	// var arr [3]string = [3]string{"apple", "banana", "cherry"}

	slice := []string{"apple", "banana", "cherry"}

	slice = append(slice, "mango")

	// fmt.Println(arr)
	// fmt.Println(slice)

	nums := []int{1, 2, 3}
	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	nums = append(nums, 4, 5)

	fmt.Println(len(nums))
	fmt.Println(cap(nums))

}
