package main

import (
	"fmt"
	"sync"
)

func numbers(i int, channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range channel {
		fmt.Printf("Worker %d received: %d\n", i, val)
	}
}

func main() {
	channel := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go numbers(i, channel, &wg)
	}

	for i := 1; i <= 9; i++ {
		channel <- i
	}

	close(channel)

	wg.Wait()
	fmt.Println("✅ All work done")

}
