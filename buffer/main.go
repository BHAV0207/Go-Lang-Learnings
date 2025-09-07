package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	fmt.Println("sending 1")
	ch <- 1
	fmt.Println("sent 1")

	fmt.Println("Sending 2")
	ch <- 2
	fmt.Println("Sent 2")

	fmt.Println("Receiving", <-ch)
	// without the above line there will be an error as the buffered channel has only two slots available and if both are filled then for the third to enter we need to atleast empty one of them which the above line does

	fmt.Println("Sending 3")
	ch <- 3
	fmt.Println("Sent 3")

}
