package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func Add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("CPU's", runtime.NumCPU())           //returns number of logical CPU cores. (Your system has 16.)
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0)) //tells how many OS threads Go will use at once (default = all cores). So we’re just printing some environment info.

	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	/*
	   🔹 Analogy 🕹️
	   Imagine:
	   You rent a power generator (ctx) with an auto-shutoff at 150ms.
	   You also get a red stop button (cancel) that lets you shut it down sooner.

	   Now:
	   If you never press the button, the generator shuts off after 150ms anyway.
	   But Go style says: always press the stop button when you’re done, just in case.
	   That’s what defer cancel() ensures.
	*/

	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("operation finished (too late)")
	case <-ctx.Done():
		fmt.Println("context timed out")
	}

	/*n this main.go, context.WithTimeout is demonstrating how to set a deadline for an operation — and since our “operation” is intentionally too slow (200ms), the context cancels it after 150ms.*/
}
