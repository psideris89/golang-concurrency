package main

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

}
