package queue

import (
	"sync"
)

// The WaitGroup makes the queue thread-safe as it will
//    make all mutex's have to wait until the other ones have
//    finished locking/unlocking before running the function
var WaitGroup sync.WaitGroup = sync.WaitGroup{}

// Type Item interface{}
//	The 'Item' Type is the type of variables that will be going inside the queue slice
//  The Item is declared as interface so it is possible to have multiple types
// 	   within the Queue Slice
type Item interface{}

// type ItemQueue struct
//	The 'ItemQueue' Struct contains the []'Type Item interface{}' slice
//  This struct holds two keys,
//     - items -> the []'Type Item interface{}' slice
//     - lock -> the mutex lock which prevents overwrites and data corruption
//			  ↳ We use RWMutex instead of Mutex as it's better for majority read slices
type ItemQueue struct {
	items []Item
	mutex sync.RWMutex
}

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
	// Wait Until previous WaitGroup task is finished
	WaitGroup.Wait()

	// Increase WaitGroup tasks then defer it
	WaitGroup.Add(1)
	defer WaitGroup.Done()

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
