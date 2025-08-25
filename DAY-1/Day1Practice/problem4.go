package main

import (
	"fmt"
	"time"
)

func print1to5(channel chan int) {
	for i := 1; i <= 5; i++ {
		channel <- i
		time.Sleep(300 * time.Millisecond)
	}
	close(channel)
}

func printAtoE(channel chan string) {
	arr := []string{"A", "B", "C", "D", "E"}
	for _, val := range arr {
		channel <- val
		time.Sleep(1000 * time.Millisecond)
	}
	close(channel)
}
func main() {

	channel1 := make(chan int)
	channel2 := make(chan string)

	go print1to5(channel1)
	go printAtoE(channel2)

	for channel1 != nil || channel2 != nil {
		select {
		case num, ok := <-channel1:
			if !ok {
				channel1 = nil
				continue
			}

			fmt.Println("Number ", num)
		case letter, ok := <-channel2:
			if !ok {
				channel2 = nil
				continue
			}
			fmt.Println("Letter ", letter)
		}
	}

}
