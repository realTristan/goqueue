package queue

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutFuncAsync(t *testing.T) {
	// Initialize queue
	queue := Create[int]()

	// Create waitgroup for goroutines
	wg := sync.WaitGroup{}

	// Create goroutines (100)
	for i := 0; i < 100; i++ {
		ii := i
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			// give ii an memory address for each iteration
			// Put item in the queue
			queue.Put(index)

			assert.GreaterOrEqual(t, len(queue.items), 1)

			// Find the item from the queue
			assert.True(t, queue.Contains(index), "Contains(index) should be found in queue")
		}(ii)
	}
	// Wait for goroutines to finish (never)
	wg.Wait()

	item := queue.GetAtIndex(50)
	assert.True(t, queue.Contains(*item), "Contains(item) should contain item at index 50")
	assert.False(t, queue.Contains(float32(*item)), "Contains(item) should be type strict")

	assert.Len(t, queue.items, 100, ".items should still include all 100 items")
	assert.Equal(t, 100, queue.Size(), "Size() should return correct length")

	allItems := queue.Show()
	assert.Contains(t, *allItems, *item, "Show() should contain item at index 50")

	assert.True(t, queue.IsNotEmpty(), "queue should be not empty")

	queue.Clear()
	assert.True(t, queue.IsEmpty(), "queue should be empty")
}
