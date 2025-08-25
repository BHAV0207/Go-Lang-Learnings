package day1practice

import (
	"context"
	"fmt"
	"time"
)

func SlowOperation() string {
	time.Sleep(2 * time.Second)
	return "done"
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resultCh := make(chan string)

	go func() {
		resultCh <- SlowOperation()
	}()

	select {
	case res := <-resultCh:
		fmt.Print("Result", res)

	case <-ctx.Done():
		fmt.Print("operation timed out")
	}

}
