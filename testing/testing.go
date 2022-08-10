package testing

import (
	"fmt"
	"sync"

	Queue "github.com/realTristan/GoQueue/queue"
)

// ThreadSafety() Function checks to make sure reading/writing the queue
//     is thread safe
func ThreadSafety() {
	// Track GoRoutines
	var count int = 0

	// Initialize queue
	queue := Queue.Create()

	// Create waitgroup for goroutines
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Create goroutines (100)
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

func main() {

}
