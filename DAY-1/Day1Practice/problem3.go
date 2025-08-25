package main

import (
	"fmt"
	"time"
)

func chef(name string, cookTime time.Duration, channel chan string) {
	fmt.Println(name, "is cooking...")
	time.Sleep(cookTime)
	channel <- name + " is ready!"
}

func main() {
	channel := make(chan string)

	// Start 3 chefs
	go chef("Pasta", 1*time.Second, channel)
	go chef("Pizza", 2*time.Second, channel)
	go chef("Burger", 3*time.Second, channel)

	// Collect results from all chefs
	for i := 0; i < 3; i++ {
		message := <-channel
		fmt.Println(message)
	}
}
