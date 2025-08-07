package main

import (
	"fmt"
	"sync"
	"time"
)

/*
3. Ping-Pong with channels
Write two goroutines:

Goroutine A sends "ping" into a channel every second.

Goroutine B receives "ping" and prints "pong" to the console.

✅ Run this for 5 seconds then exit.
*/

func main() {
	msg := make(chan string)
	done := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go ping(msg, done, wg)
	go pong(msg, done, wg)

	interrupt := time.After(time.Duration(5) * time.Second)

	select {
	case <-interrupt:
		fmt.Println("interrupted")
		close(done)
		close(msg)
		wg.Wait()
	}

	fmt.Println("done")
}

func ping(ch chan<- string, done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-done:
			fmt.Println("Ping: received shutdown signal")
			return
		case ch <- "ping":
			time.Sleep(1 * time.Second)
		}
	}
}

func pong(ch <-chan string, done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Println("Pong: received shutdown signal")
			return
		case msg, ok := <-ch:
			if !ok {
				// msg channel closed
				fmt.Println("Pong: msg channel closed")
				return
			}
			fmt.Println("Pong received:", msg)
		}
	}
}
