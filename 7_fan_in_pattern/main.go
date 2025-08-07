package main

import (
	"context"
	"fmt"
	"time"
)

// @ignore-format
/*
🧪 Exercise 7: Fan-In Pattern

Requirements:
	1. Starts 2 producer goroutines, each sending strings into their own channel (e.g., "P1-msg1", "P2-msg1"…).
	2. Starts 1 fan-in goroutine that:
		- Reads from both producer channels.
		- Sends all messages into a single merged channel.
	3. The main() function:
		- Reads and prints 5 total messages from the merged channel, then exits.
*/
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch1 := make(chan string)
	ch2 := make(chan string)
	merged := make(chan string)

	go producer(ctx, "P1", ch1)
	go producer(ctx, "P2", ch2)

	go fanIn(ctx, ch1, ch2, merged)

	for i := 0; i < 5; i++ {
		msg := <-merged
		fmt.Println(msg)
	}

	cancel()
}

func producer(ctx context.Context, msg string, chn chan string) {
	defer close(chn)

	ctr := 1
	for {
		select {
		case <-ctx.Done():
			return
		default:
			chn <- fmt.Sprintf("%s-msg%d", msg, ctr)
			ctr++
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func fanIn(ctx context.Context, ch1, ch2 <-chan string, out chan<- string) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-ch1:
			if ok {
				out <- msg
			}
		case msg, ok := <-ch2:
			if ok {
				out <- msg
			}
		}
	}
}
