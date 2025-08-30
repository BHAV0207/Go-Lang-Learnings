package main

import (
	"fmt"
	"sync"
)

func square(num int, channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ans := num * num

	channel <- ans

	// close(channel)

}

func main() {
	var wg sync.WaitGroup
	// wg.Add(10);

	ch := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go square(i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)

		//  here we have written the close function in here because there is only one channel having all 10 results
	}()

	for res := range ch {
		fmt.Println(res)
	}

	fmt.Println("all calc done")

}
