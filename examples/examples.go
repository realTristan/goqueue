package examples

import (
	"fmt"

	Queue "github.com/realTristan/GoQueue/queue"
)

// Add Items to the queue
func AddItems() {
	queue := Queue.Create[any]()

	// Add items from a slice to the queue
	items := [3]interface{}{1.1, 1, "String"}
	for _, item := range items {
		queue.Put(item)
	}
}

// Remove items from the queue
func RemoveItems() {
	queue := Queue.Create[string]()

	// Remove at index
	removedItem := queue.RemoveAtIndex(0)
	println(removedItem)

	// Search and remove
	queue.Remove("Item")
}

// Get items from the queue
func GetItems() {
	queue := Queue.Create[any]()

	// Get the item from the queue (doesn't remove it from the queue)
	item := queue.Get()
	println(item)

	// Grab the item from the queue (removes it from the queue)
	_item := queue.Grab()
	println(_item)
}

// Other Queue Functions
func OtherFunctions() {
	queue := Queue.Create[string]()
	if queue.Contains("Item") {
		println("Contains Item")
	}

	// Clear the queue
	queue.Clear()

	// Show the queue contents
	queueContents := queue.Show()
	fmt.Printf("Queue contents: %v\n", queueContents)

	// Get item at specific index
	itemAtIndex := queue.GetAtIndex(1)
	fmt.Printf("Item at index 1: %d\n", itemAtIndex)

	// Returns whether queue is empty
	isEmpty := queue.IsEmpty()
	fmt.Printf("Is queue empty: %v\n", isEmpty)

	// Returns whether queue is not empty
	isNotEmpty := queue.IsNotEmpty()
	fmt.Printf("Is queue not empty: %v\n", isNotEmpty)

	// Returns the length of the queue slice
	queueLength := queue.Size()
	fmt.Printf("Length of queue: %d\n", queueLength)
}
