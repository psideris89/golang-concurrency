package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// @ignore-format
/*
🧪 Exercise 6: Context Cancellation Across Multiple Workers

Requirements:
	1. Launches 3 worker goroutines that each:
		- Periodically print a message.
		- Respect cancellation via context.Context.
	2. The main function:
		- Cancels the context after 3 seconds.
		- Waits for all workers to exit gracefully.
	3. Your code should show that:
		- All workers stop after cancellation.
		- There are no goroutine leaks or hanging behavior.
*/
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("Cancelling context")
		cancel()
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go work2(i, int64(i+1), ctx, &wg)
	}

	wg.Wait()
}

func work2(workerId int, processingWindowInSec int64, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Terminating worker %d\n", workerId)
			return
		default:
			fmt.Printf("Worker %d processing data\n", workerId)
			time.Sleep(time.Duration(processingWindowInSec) * time.Second)
		}
	}
}
