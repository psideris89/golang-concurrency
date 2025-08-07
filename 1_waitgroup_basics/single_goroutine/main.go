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
	- Create a goroutine that performs work (e.g., prints something), and track it using a sync.WaitGroup.
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

	wg.Wait()

	elapsed := time.Since(started)
	fmt.Printf("elapsed: %s\n", elapsed)
}
