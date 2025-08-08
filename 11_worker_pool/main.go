package main

import (
	"fmt"
	"sync"
)

//@ignore-format
/*
🧪 Exercise 12: Worker Pool with Job Results

Requirements:
	- Create a worker pool with 3 workers that process 10 integer jobs and send their squared results back to the main
	goroutine, which prints them. Ensure proper use of sync.WaitGroup and close channels correctly to avoid leaks.
*/

type res struct {
	workerId   int
	initial    int
	calculated int
}

func main() {
	input := make(chan int, 10)
	out := make(chan res, 10)
	wg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			for val := range input {
				out <- res{
					workerId:   workerId,
					initial:    val,
					calculated: val * val,
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	go func() {
		// If we have only 10 jobs and the input channel has size 10 then we don't really need a goroutine
		// but this can be an issue in case we have jobs > input channel capacity
		for i := 1; i <= 10; i++ {
			input <- i
		}
		close(input)
	}()

	for val := range out {
		fmt.Printf("Initial value %d - calculated %d - workerId %d\n", val.initial, val.calculated, val.workerId)
	}
}
