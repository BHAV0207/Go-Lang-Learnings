package main

import "fmt"

func nums(channel chan int) {
	for i := 1; i <= 5; i++ {
		channel <- i
	}
	close(channel)
}

func main() {
	channel := make(chan int)
	go nums(channel)

	arr := make([]int, 0, 5)

	for val := range channel {
		arr = append(arr, val)
	}

	fmt.Println(arr)

}
