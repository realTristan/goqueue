
![Untitled-1](https://user-images.githubusercontent.com/75189508/183445864-652f08ae-ee72-4368-b14e-4cd39362d62e.png) <h1>GoQueue</h1> ![Stars](https://img.shields.io/github/stars/realTristan/GoQueue?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/realTristan/GoQueue?label=Watchers) 

![Go Queue Banner](https://user-images.githubusercontent.com/75189508/183435878-e5669071-df93-478a-a364-245862dadddb.png)

Flexible Queue System Implementation for Go.

# Why GoQueue?
GoQueue is a light weight, easy to read open source module that uses solely native golang code. GoQueue's functions are based off of the python queue library so the learning curve is not time consuming.

# Installation
`go get https://github.com/realTristan/GoQueue`


# Usage
```go

///////////////////////
// GoQueue Usage //
///////////////////////

package examples

import (
	Queue "queue/src"
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
	itemAtIndex := queue.GetAtIndex(1)

	// Returns whether queue is empty
	isEmpty := queue.IsEmpty()

	// Returns whether queue is not empty
	isNotEmpty := queue.IsNotEmpty()

	// Returns the length of the queue slice
	queueLength := queue.Size()
}

///////////////////////
// GoQueue Functions //
///////////////////////

// Create() -> *ItemQueue
// The Create() function will return an empty ItemQueue
func Create() *ItemQueue {
	var q ItemQueue = ItemQueue{}
	q.items = []Item{}
	return &q
}

// q.Secure(func()) -> None
// The Secure() function is used to lock the ItemQueue before executing the provided function
// 	   then unlock the ItemQueue after the function has been executed
func (q *ItemQueue) Secure(function func()) {
	q.lock.Lock()
	function()
	q.lock.Unlock()
}

// q.Index(index integer) -> *Item
// The RemoveAtIndex() function is used to remove an item at the provided index of the ItemQueue
// The function will then return the removed item if the user requires it's use
func (q *ItemQueue) RemoveAtIndex(i int) *Item {
	var item Item = q.items[i]
	q.Secure(func() {
		q.items = append(q.items[:i], q.items[i+1:]...)
	})
	return &item
}

// q.Contains(Item) -> None
// The Contains() function will scheck whether the provided ItemQueue contains
//	  the given Item (_item)
func (q *ItemQueue) Contains(_item Item) bool {
	for i := 0; i < len(q.items); i++ {
		if q.items[i] == _item {
			return true
		}
	}
	return false
}

// q.Remove(Item) -> None
// The Remove() function will secure the ItemQueue before iterating
//	  through said ItemQueue and remove the given Item (_item)
func (q *ItemQueue) Remove(_item Item) {
	q.Secure(func() {
		for i := 0; i < len(q.items); i++ {
			if q.items[i] == _item {
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
	q.Secure(func() {
		q.items = append(q.items, q.items[0])
		q.items = q.items[1:len(q.items)]
	})
	return &q.items[0]
}

// q.Grab() -> Item
// The Grab() function will return the first item of the ItemQueue then
//    remove it from said ItemQueue
func (q *ItemQueue) Grab() *Item {
	var item Item = q.items[0]
	q.Secure(func() {
		q.items = q.items[1:len(q.items)]
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
	return &q.items
}

// q.GetAtIndex(index integer) -> *Item
// The GetAtIndex() function is used to return an item at the provided index of the ItemQueue
func (q *ItemQueue) GetAtIndex(i int) *Item {
	return &q.items[i]
}

// q.IsEmpty() -> bool
// The IsEmpty() function will return whether the provided ItemQueue contains any Items
func (q *ItemQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// q.IsNotEmpty() -> bool
// The IsNotEmpty() function will return whether the provided ItemQueue contains any Items
func (q *ItemQueue) IsNotEmpty() bool {
	return len(q.items) > 0
}

// q.Size() -> int
// The Size() function will return the length of the ItemQueue slice
func (q *ItemQueue) Size() int {
	return len(q.items)
}
```