package main

import (
	"fmt"
	"sync"

	Queue "github.com/realTristan/GoQueue/queue"
)

func main() {
	// Track GoRoutines
	var count int = 0

	// Initialize queue
	queue := Queue.Create()

	// Create waitgroup for goroutines
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Create goroutines
	for i := 0; i < 100; i++ {
		go func() {
			for {
				count++

				// Put item in the queue
				queue.Put(count)

				// Get the item from the queue
				item := queue.Grab()

				// Print the item
				fmt.Println(*item)
			}
		}()
	}
	// Wait for goroutines to finish (never)
	wg.Wait()
}
