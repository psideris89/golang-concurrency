package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// @ignore-format
/*
🧪 Exercise 5: Context Cancellation with Goroutines

Requirements:
	- Starts a worker goroutine that performs some repetitive task (e.g., printing or sleeping).
	- The main function should create a context.Context with cancellation.
	- After 3 seconds, cancel the context and gracefully stop the worker.
*/
func main() {
	ctx, cancel := context.WithCancel(context.TODO())

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("Cancelling context")
		cancel()
	}()

	go work(ctx, &wg)

	wg.Wait()
}

func work(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Terminating worker")
			return
		default:
			fmt.Println("Worker is processing data...")
			time.Sleep(1 * time.Second)
		}
	}
}
