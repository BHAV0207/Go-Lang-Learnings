package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	num := 1
	for {
		ch <- num
		num++
		time.Sleep(200 * time.Millisecond) // send every 200ms
	}
}

func main() {
	ch := make(chan int)

	// Start producer goroutine
	go producer(ch)

	// Run for 1 second only
	timeout := time.After(1 * time.Second)

	for {
		select {
		case val := <-ch:
			fmt.Println("Received:", val)
		case <-timeout:
			fmt.Println("⏰ Timeout! Exiting...")
			return
		}
	}
}
