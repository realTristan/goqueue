
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
	queue := Queue.Create()
	
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
	println(*item)

	// Grab the item from the queue (removes it from the queue)
	_item := queue.Grab()
	println(*_item)
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
	itemAtIndex := queue.GetAtIndex(1)

	// Returns whether queue is empty
	isEmpty := queue.IsEmpty()

	// Returns whether queue is not empty
	isNotEmpty := queue.IsNotEmpty()

	// Returns the length of the queue slice
	queueLength := queue.Size()
}
```

# GoQueue Functions
```go

// Create() -> *ItemQueue
// The Create() function will return an empty ItemQueue
func Create() *ItemQueue {
	return &ItemQueue{items: []Item{}}
}

// q.Secure(func()) -> None
// The Secure() function is used to lock the ItemQueue before executing the provided function
// 	   then unlock the ItemQueue after the function has been executed
func (q *ItemQueue) Secure(function func()) {
	// Lock the queue then unlock once function closes
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// Run the provided function
	function()
}

// q.Index(index integer) -> *Item
// The RemoveAtIndex() function is used to remove an item at the provided index of the ItemQueue
// The function will then return the removed item if the user requires it's use
func (q *ItemQueue) RemoveAtIndex(i int) *Item {
	var item Item
	q.Secure(func() {
		item = q.items[i]
		q.items = append(q.items[:i], q.items[i+1:]...)
	})
	return &item
}

// q.Contains(Item) -> None
// The Contains() function will scheck whether the provided ItemQueue contains
//	  the given Item (_item)
func (q *ItemQueue) Contains(item Item) bool {

	// Lock Reading
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Iterate over the queue
	for i := 0; i < len(q.items); i++ {
		if q.items[i] == item {
			return true
		}
	}
	return false
}

// q.Remove(Item) -> None
// The Remove() function will secure the ItemQueue before iterating
//	  through said ItemQueue and remove the given Item (_item)
func (q *ItemQueue) Remove(item Item) {
	q.Secure(func() {
		for i := 0; i < len(q.items); i++ {
			if q.items[i] == item {
				q.items = append(q.items[:i], q.items[i+1:]...)
				return
			}
		}
	})
}

// q.Put(Item) -> None
// The Put() function is used to add a new item to the provided ItemQueue
func (q *ItemQueue) Put(i Item) {
	q.Secure(func() {
		q.items = append(q.items, i)
	})
}

// q.Get() -> Item
// The Get() function will append the first item of the ItemQueue to the back of the slice
//    then remove it from the front
// The function returns the first item of the ItemQueue
func (q *ItemQueue) Get() *Item {
	var item Item
	q.Secure(func() {
		item = q.items[0]
		q.items = append(q.items, q.items[0])
		q.items = q.items[1:]
	})
	return &item
}

// q.Grab() -> Item
// The Grab() function will return the first item of the ItemQueue then
//    remove it from said ItemQueue
func (q *ItemQueue) Grab() *Item {
	var item Item
	q.Secure(func() {
		item = q.items[0]
		q.items = q.items[1:]
	})
	return &item
}

// q.Clear() -> None
// The Clear() function will secure the queue then remove all of its items
func (q *ItemQueue) Clear() {
	q.Secure(func() {
		q.items = []Item{}
	})
}

// q.Show() -> *[]Item
// The Show() function will return the ItemQueue's items
func (q *ItemQueue) Show() *[]Item {

	// Lock Reading
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Return the queue items
	return &q.items
}

// q.GetAtIndex(index integer) -> *Item
// The GetAtIndex() function is used to return an item at the provided index of the ItemQueue
func (q *ItemQueue) GetAtIndex(i int) *Item {

	// Lock Reading
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Return the item at the specific index
	return &q.items[i]
}

// q.IsEmpty() -> bool
// The IsEmpty() function will return whether the provided ItemQueue contains any Items
func (q *ItemQueue) IsEmpty() bool {

	// Lock Reading
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Return whether queue is empty
	return len(q.items) == 0
}

// q.IsNotEmpty() -> bool
// The IsNotEmpty() function will return whether the provided ItemQueue contains any Items
func (q *ItemQueue) IsNotEmpty() bool {

	// Lock Reading
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Return whether length is greater than 0
	return len(q.items) > 0
}

// q.Size() -> int
// The Size() function will return the length of the ItemQueue slice
func (q *ItemQueue) Size() int {

	// Lock Reading
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Return the queue length
	return len(q.items)
}

```
