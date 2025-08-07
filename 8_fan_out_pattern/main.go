package main

import (
	"fmt"
	"sync"
	"time"
)

// @ignore-format
/*
🧪 Exercise 8: Fan-Out / Worker Pool

Requirements:
	1. Creates a job channel (e.g., chan int) and sends 10 jobs into it.
	2. Starts 3 worker goroutines, each:
		- Reads from the job channel.
		- "Processes" the job by printing and sleeping for 500ms.
	3. Use a sync.WaitGroup to ensure the program only exits after all jobs are processed.
*/
func main() {
	chn := make(chan int, 10)
	wg := sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		chn <- i
	}

	close(chn)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(&wg, chn, i)
	}

	wg.Wait()
}

func worker(wg *sync.WaitGroup, chn <-chan int, i int) {
	defer wg.Done()

	// Option 2
	//for msg := range chn {
	//	fmt.Printf("Worker %d processing job %d\n", i, msg)
	//	time.Sleep(500 * time.Millisecond)
	//}

	for {
		select {
		case msg, ok := <-chn:
			if !ok {
				return
			}
			fmt.Printf("Worker %d processing job %d\n", i, msg)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
