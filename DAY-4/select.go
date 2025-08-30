package main

import (
	"fmt"
	"sync"
	"time"
)

func Ping(channel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		channel <- "ping"
		time.Sleep(200 * time.Millisecond)
	}

	close(channel)
}

func pong(channel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for true {
		channel <- "Pong"
		time.Sleep(350 * time.Millisecond)
	}
	close(channel)
}

func main() {

	var wg sync.WaitGroup
	chanPing := make(chan string)
	chanPong := make(chan string)

	wg.Add(2)

	go Ping(chanPing, &wg)
	go pong(chanPong, &wg)

	go func() {
		for {
			select {
			case msg, ok := <-chanPing:
				if ok {
					fmt.Println("Received:", msg)
				} else {
					chanPing = nil // disable this case when channel closes
				}
			case msg, ok := <-chanPong:
				if ok {
					fmt.Println("Received:", msg)
				} else {
					chanPong = nil
				}
			}

			if chanPing == nil && chanPong == nil {
				break
			}
		}
	}()

	wg.Wait()
	time.Sleep(time.Second) // wait for receiver to finish
	fmt.Println("🏓 Ping-pong game over")

}
