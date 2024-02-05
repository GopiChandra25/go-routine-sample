package main

import (
	"fmt"
	"goroutine-sample/hello"
	"goroutine-sample/pkg/util"
	"strconv"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, done chan bool) {
	defer wg.Done()
	fmt.Println("Before sleep ==>" + strconv.Itoa(id))
	// Simulate a long-running task
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Println("After ==>" + strconv.Itoa(id))

	// Signal that the task is done
	done <- true
}

func main() {
	var wg sync.WaitGroup
	// Create a channel to signal completion
	done := make(chan bool)

	// Number of Goroutines
	numGoroutines := 5

	// Add Goroutines to the WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go worker(i+1, &wg, done)
	}

	// Start a Goroutine to wait for all workers to finish
	go func() {
		wg.Wait()
		close(done)
	}()

	/*for item := range done {
		fmt.Println(item)
	}*/

	// Wait for the tasks to complete or timeout after a certain duration
	timeout := 5 * time.Second
	timer := time.NewTimer(timeout)
ForSelectLoop:
	for {
		select {
		case value, ok := <-done:
			if !ok {
				fmt.Println("All tasks completed successfully.")
				//return
				break ForSelectLoop
			}
			fmt.Println(value)
		case <-timer.C:
			fmt.Println("Timeout: Some tasks took too long. Terminating Goroutines.")
			// Optionally, you can take further actions like closing resources or handling cleanup.
			// However, forcefully stopping Goroutines is not recommended.
			break ForSelectLoop
			//return
		}
	}

	fmt.Println("Main function completed")
}
