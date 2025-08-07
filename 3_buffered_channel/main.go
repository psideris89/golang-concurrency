package main

import (
	"fmt"
	"sync"
)

// @ignore-format
/*
🧪 Exercise 3: Fan-Out with a Buffered Channel

Requirements:
	1. Create a buffered channel (size 5) and send integers 1–5 into the channel
	2. Start 3 worker goroutines that consume the data in the channel
*/
func main() {
	chn := make(chan int, 5)
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			//time.Sleep(2 * time.Second)
			defer wg.Done()
			for val := range chn {
				fmt.Printf("Worker %d processed job %d\n", n, val)
			}
		}(i)
	}

	// Don't need to put this in go-routine because the channel is fully buffered, you could
	// safely send and close synchronously without using a goroutine.
	for i := 1; i <= 5; i++ {
		chn <- i
	}
	fmt.Println("Done producing")
	close(chn)

	wg.Wait()
}
