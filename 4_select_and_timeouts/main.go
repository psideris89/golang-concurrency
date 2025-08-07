package main

import (
	"fmt"
	"time"
)

// @ignore-format
/*
🧪 Exercise 4: Select and Timeouts

Requirements:
	1. Create a channel messages of type chan string.
	2. Start a goroutine that:
		- Sleeps for 2 seconds, then sends "hello" into the messages channel.
	3. In main(), use a select statement to:
		- Receive from the messages channel.
		- Timeout if a message doesn’t arrive within 1 second, and print "timeout".
*/
func main() {
	chn := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chn <- "hello"
	}()

	for {
		select {
		case msg := <-chn:
			fmt.Println("received: " + msg)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout")
			return
		}
	}
}
