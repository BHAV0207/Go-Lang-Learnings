package main

import (
	"fmt"
	"sync"
)

func Producer(channel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		channel <- i
	}

	close(channel)
}

func Reciver(channel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range channel {
		fmt.Println(val)
	}
}

func main() {
	channel := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go Producer(channel, &wg)

	go Reciver(channel, &wg)

	wg.Wait()

	fmt.Println("done printing ")
}
