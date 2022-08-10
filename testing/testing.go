package main

import (
	"fmt"
	"sync"

	Queue "github.com/realTristan/GoQueue/queue"
)

// ThreadSafety() Function checks to make sure reading/writing the queue
//     is thread safe
func ThreadSafety(queue *Queue.ItemQueue) {
	// Track GoRoutines
	var (
		count int            = 0
		wg    sync.WaitGroup = sync.WaitGroup{}
	)
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

// Testing: Show -> show queue contents
func ShowQueueContents(queue *Queue.ItemQueue, test string) {
	queueContents := queue.Show()
	fmt.Printf("Testing: %s: Queue Contents: %v\n", test, queueContents)
}

// Testing: Put -> putting an item into the queue
func TestPut(queue *Queue.ItemQueue) {

	// Show queue contents before put
	ShowQueueContents(queue, "Put")

	// Put item into queue
	queue.Put("Item")

	// Show queue contents after put
	ShowQueueContents(queue, "Put")
}

// Testing: GetAtIndex -> getting an item at a specific index
func TestGetAtIndex(queue *Queue.ItemQueue) {

	// Get item from the queue at index: 1
	itemAtIndex := queue.GetAtIndex(1)
	fmt.Printf("Item at index 1: %d\n", itemAtIndex)
}

// Testing: Get -> getting an item from the queue
func TestGet(queue *Queue.ItemQueue) {

	// Get item from the queue
	getItem := queue.Get()
	fmt.Printf("Get Item in queue: %v\n", getItem)
}

// Testing: Grab -> grabbing an item from the queue
func TestGrab(queue *Queue.ItemQueue) {

	// Show queue contents before grab
	ShowQueueContents(queue, "Grab")

	// Grab the item
	grabItem := queue.Grab()
	fmt.Printf("Grab Item in queue: %v\n", grabItem)

	// Show queue contents after grab
	ShowQueueContents(queue, "Grab")
}

// Testing: IsEmpty -> checking whether queue is empty
func TestIsEmpty(queue *Queue.ItemQueue) {

	// Check if queue is empty
	isEmpty := queue.IsEmpty()
	fmt.Printf("Is queue empty: %v\n", isEmpty)

	// Show whether queue is empty or not
	ShowQueueContents(queue, "IsEmpty")
}

// Testing: IsNotEmpty -> checking whether queue is not empty
func TestIsNotEmpty(queue *Queue.ItemQueue) {

	// Check if queue is not empty
	isNotEmpty := queue.IsNotEmpty()
	fmt.Printf("Is queue not empty: %v\n", isNotEmpty)

	// Show whether queue is empty or not
	ShowQueueContents(queue, "IsNotEmpty")
}

// Testing: Size -> checking the length of the queue slice
func TestSize(queue *Queue.ItemQueue) {

	// Get queue size
	queueLength := queue.Size()
	fmt.Printf("Length of queue: %d\n", queueLength)

	// Show queue contents to check whether size is correct
	ShowQueueContents(queue, "Size")
}

// Testing: Contains -> check if queue contains item
func TestContains(queue *Queue.ItemQueue) {

	// Get whether queue contains "Item"
	contains := queue.Contains("Item")
	fmt.Printf("Contains \"item\": %v\n", contains)

	// Show queue contents to show whether it contains item or not
	ShowQueueContents(queue, "Contains")
}

// Testing: Clear -> clear item queue
func TestClear(queue *Queue.ItemQueue) {
	// Put new item into the queue
	queue.Put("Item")

	// Show the queue
	ShowQueueContents(queue, "Clear")

	// Clear the queue
	queue.Clear()

	// Show the cleared queue
	ShowQueueContents(queue, "Clear")
}

// Testing: Remove -> remove item in the queue
func TestRemove(queue *Queue.ItemQueue) {
	// Add new item to the queue
	queue.Put("Item")

	// Show item queue
	ShowQueueContents(queue, "Remove")

	// Remove item from the queue
	queue.Remove("Item")

	// Show item queue
	ShowQueueContents(queue, "Remove")

}

// Testing: RemoveAtIndex -> remove item at specific index
func TestRemoveAtIndex(queue *Queue.ItemQueue) {
	// Add new item to the queue
	queue.Put("Item")

	// Show item queue
	ShowQueueContents(queue, "RemoveAtIndex")

	// Remove item from the queue
	queue.RemoveAtIndex(0)

	// Show item queue
	ShowQueueContents(queue, "RemoveAtIndex")
}

// Main function
func main() {
	// Testing: Create -> creating a new queue
	var queue *Queue.ItemQueue = Queue.Create()
	ShowQueueContents(queue, "Create")

	/////////////////////////////
	// Testing Functions
	//
	// ThreadSafety(queue)
	// TestPut(queue)
	// TestGetAtIndex(queue)
	// TestGet(queue)
	// TestGrab(queue)
	// TestIsEmpty(queue)
	// TestIsNotEmpty(queue)
	// TestSize(queue)
	// TestClear(queue)
	// TestRemove(queue)
	// TestRemoveAtIndex(queue)
	//
	/////////////////////////////
}
