
# GoQueue ![Stars](https://img.shields.io/github/stars/realTristan/GoQueue?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/realTristan/GoQueue?label=Watchers)

![Go Queue Banner](https://user-images.githubusercontent.com/75189508/183435878-e5669071-df93-478a-a364-245862dadddb.png)

Flexible Queue System for Go.

# Why GoQueue?
GoQueue is a light weight, easy to read open source module that uses solely native golang code. GoQueue's functions are based off of the python queue library so the learning curve is not as time consuming.

# Installation
`go get -u github.com/realTristan/GoQueue`


# Quick Usage
```go
package main

import (
	Queue "github.com/realTristan/GoQueue/queue"
)

func main() {
	// Create a new queue
	queue := Queue.Create[string]()
	
	// Put item into the queue
	queue.Put("Item")
	
	// Get the item from the queue
	item := queue.Get()
	
	// Print the item
	println(*item)
	
	// Output -> "Item"
}
```

# GoQueue Usage
```go
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
	queue := Queue.Create[int]()

	// Get the item from the queue (doesn't remove it from the queue)
	item := queue.Get()
	println(*item)

	// Grab the item from the queue (removes it from the queue)
	_item := queue.Grab()
	println(*_item)
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
	queue.Show()

	// Get item at specific index
	itemAtIndex := queue.GetAtIndex(1)

	// Returns whether queue is empty
	isEmpty := queue.IsEmpty()

	// Returns whether queue is not empty
	isNotEmpty := queue.IsNotEmpty()

	// Returns the length of the queue slice
	queueLength := queue.Size()
}
```
