package main

import (
	"fmt"
	"sync"
)

// @ignore-format
/*
🧪 Exercise 2: Unbuffered Channel Communication

Requirements:
	- Use an unbuffered chan int to send 5 integers from multiple goroutines to a single receiver. Print each received value.
	- Use multiple producers, that produce in parallel.
*/
func main() {
	chn := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Adding %d\n", n)
			chn <- n
		}(i)
	}

	// To not block the main goroutine in order to wait for all producers to finish and then close we use a goroutine.
	// That way the processing can start while waiting for producers to finish.
	go func() {
		wg.Wait()
		close(chn)
	}()

	fmt.Println("Started processing")
	for r := range chn {
		fmt.Println(r)
	}
}
