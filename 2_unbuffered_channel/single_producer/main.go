package main

import "fmt"

// @ignore-format
/*
🧪 Exercise 2: Unbuffered Channel Communication

Requirements:
	- Use an unbuffered chan int to send 5 integers from multiple goroutines to a single receiver. Print each received value.
	- Use a single producer.
*/
func main() {
	chn := make(chan int)

	go func() {
		for i := range [5]int{} {
			fmt.Printf("Adding %d\n", i)
			chn <- i
		}
		close(chn)
	}()

	fmt.Println("Starting processing")
	for val := range chn {
		fmt.Println(val)
	}
}
