package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(1 * time.Second) // simulate work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	for i := 1; i <= 5; i++ {
		go worker(i)
	}

	// ❌ main ends before goroutines finish
	fmt.Println("All work launched")
}

// this is a problem that we can face while using go routines that the goroutines does not finish in time before the main function ends
//  to solve this issue we use waitGroups
