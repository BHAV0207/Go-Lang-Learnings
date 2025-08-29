package main

import (
	"fmt"
	"sync"
	"time"
)

func printAtoE(channel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	arr := []string{"A", "B", "C", "D", "E"}
	for _, val := range arr {
		channel <- val
		time.Sleep(150 * time.Millisecond)
	}

	close(channel)
}

func print1to5(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		channel <- i
		time.Sleep(350 * time.Millisecond)
	}
	close(channel)
}

func main() {
	var wg sync.WaitGroup

	channel1 := make(chan string)
	channel2 := make(chan int)

	wg.Add(2)

	go print1to5(channel2, &wg)
	go printAtoE(channel1, &wg)

	go func() {
		for num := range channel2 {
			fmt.Println("Number:", num)
		}
	}()
	go func() {
		for letter := range channel1 {
			fmt.Println("Letter:", letter)
		}
	}()

	wg.Wait() // wait until both goroutines finish
	fmt.Println("All work completed ✅")

}
