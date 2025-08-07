package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// @ignore-format
/*
🧪 Exercise 9 Recap: Timeout for Job Processing

Requirements:
	1. You have 3 worker goroutines.
	2. You send 10 jobs into a channel.s
	3. Each job must be processed within 300ms.
	4. Simulate some jobs taking longer than 300ms, and mark them as timed out if so.
	5. Use context.WithTimeout inside the worker to enforce this.
*/
func main() {
	chn := make(chan int, 10)
	for i := 0; i < 10; i++ {
		chn <- i
	}
	close(chn)

	wg := &sync.WaitGroup{}
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(chn, i, wg)
	}
	wg.Wait()
}

func worker(chn <-chan int, workerId int, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range chn {
		ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
		reschn := make(chan bool, 1)
		go func(m int) {
			processJob(m, reschn)
		}(msg)

		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d aborting job %d due to timeout\n", workerId, msg)
		case <-reschn:
			fmt.Printf("Worker %d processed message %d\n", workerId, msg)
		}
		cancel()
	}
}

func processJob(msg int, reschn chan<- bool) {
	if msg%3 == 0 {
		time.Sleep(1 * time.Second)
	} else {
		time.Sleep(100 * time.Millisecond)
	}
	reschn <- true
}
