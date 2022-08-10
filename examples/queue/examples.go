package queue_examples

import (
	Queue "github.com/realTristan/GoQueue/queue"
)

// Add Items to the queue
func AddItems() {
	queue := Queue.Create()

	// Add items from a slice to the queue
	items := [3]interface{}{1.1, 1, "String"}
	for _, item := range items {
		queue.Put(item)
	}
}

// Remove items from the queue
func RemoveItems() {
	queue := Queue.Create()

	// Remove at index
	removedItem := queue.RemoveAtIndex(0)
	println(removedItem)

	// Search and remove
	queue.Remove("Item")
}

// Get items from the queue
func GetItems() {
	queue := Queue.Create()

	// Get the item from the queue (doesn't remove it from the queue)
	item := queue.Get()
	println(item)

	// Grab the item from the queue (removes it from the queue)
	_item := queue.Grab()
	println(_item)
}

// Other Queue Functions
func OtherFunctions() {
	queue := Queue.Create()
	if queue.Contains("Item") {
		println("Contains Item")
	}

	// Clear the queue
	queue.Clear()

	// Show the queue contents
	queue.Show()

	// Get item at specific index
	// itemAtIndex := queue.GetAtIndex(1)

	// Returns whether queue is empty
	// isEmpty := queue.IsEmpty()

	// Returns whether queue is not empty
	// isNotEmpty := queue.IsNotEmpty()

	// Returns the length of the queue slice
	// queueLength := queue.Size()
}
