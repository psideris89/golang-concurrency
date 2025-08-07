package main

/*
4. Parallel summation
Given an array of 100 integers:

Split it into 4 parts.

Start 4 goroutines to sum each part.

Collect all partial sums and print the total sum.

✅ This teaches fan-out / fan-in pattern.
*/
func main() {

}

/*
Level 3: Advanced — Worker Pools and Sync
✨ Goal: real-world patterns.

5. Build a mini-worker pool
Start N workers (say, 3).

Have a list of 10 "jobs" (strings like "Task 1", "Task 2", etc).

Each worker picks jobs from a channel and "processes" them (just prints).

Main function waits until all jobs are done.

✅ Worker pool pattern!

6. Timeout a slow goroutine
Start a goroutine that does some "work" (sleep 5 seconds).

But if it takes more than 2 seconds, cancel it.

✅ Use context.WithTimeout to solve it!

🏋️‍♂️ Bonus Challenges 🚀
7. Race condition demo
Write a program where:

Two goroutines both increment a shared counter variable 1000 times each.

First without locks ➔ you'll see wrong final count!

Then fix it using:

sync.Mutex or

sync/atomic package.

✅ Teaches why data race is dangerous.
*/
