package main

import (
	"fmt"
	"sync"
	"time"
)

// @ignore-format
/*
🧪 Exercise 1: WaitGroup + Goroutine Basics

Requirements:
	- Create 2 goroutines that performs work (e.g., prints something), and track it using a sync.WaitGroup.
*/
func main() {
	started := time.Now()
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Processing A")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Processing B")
	}()

	wg.Wait()

	elapsed := time.Since(started)
	fmt.Printf("elapsed: %s\n", elapsed)
}
