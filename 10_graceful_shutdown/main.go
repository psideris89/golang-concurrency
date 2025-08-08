package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//@ignore-format
/*
🧪 Exercise 10: Graceful Shutdown with Signal Handling

Requirements:
	1. Starts a worker pool with 3 workers reading from a channel.
	2. Produces jobs continuously in a goroutine (e.g., every 200ms).
	3. Listens for OS interrupt signal (SIGINT, i.e. Ctrl+C).
	4. When the signal is received:
		- Stop accepting new jobs.
		- Gracefully shut down all workers after completing any jobs in-flight.
		- Print: Shutting down gracefully...
*/
func main() {
	chn := make(chan int)
	stop := make(chan struct{})

	wg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			for msg := range chn {
				fmt.Printf("Worker %d processing message %d\n", workerId, msg)
				//time.Sleep(1 * time.Second)
			}
		}(i)
	}

	go func() {
		ctr := 1
		for {
			select {
			case <-stop:
				fmt.Println("Closing messages channel")
				close(chn)
				return
			case chn <- ctr:
				fmt.Printf("Adding message %d to channel\n", ctr)
				ctr++
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("Shutting down gracefully...")
	fmt.Println("Closing stop channel")
	close(stop)
	fmt.Println("Waiting for workers to finish processing")
	wg.Wait()
}
