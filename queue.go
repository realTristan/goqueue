package goqueue

import "sync"

// Type Item interface{}
//
//		The 'Item' Type is the type of variables that will be going inside the queue slice
//	 The Item is declared as interface so it is possible to have multiple types
//		   within the Queue Slice
type Item interface{}

// type ItemQueue struct
//
//		The 'ItemQueue' Struct contains the []'Type Item interface{}' slice
//	 This struct holds two keys,
//	    - items -> the []'Type Item interface{}' slice
//	    - mutex -> the mutex lock which prevents overwrites and data corruption
//				  ↳ We use RWMutex instead of Mutex as it's better for majority read slices
type ItemQueue struct {
	items []Item
	mutex *sync.RWMutex
}

// Create() -> *ItemQueue
// The Create() function will return an empty ItemQueue
func Create() *ItemQueue {
	return &ItemQueue{mutex: &sync.RWMutex{}, items: []Item{}}
}

// q.secure(func()) -> None
// The Secure() function is used to lock the ItemQueue before executing the provided function
//
//	then unlock the ItemQueue after the function has been executed
func (q *ItemQueue) secure(function func()) {
	// Lock the queue then unlock once function closes
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// Run the provided function
	function()
}

// q.RemoveAtIndex(index integer) -> *Item
// The RemoveAtIndex() function is used to remove an item at the provided index of the ItemQueue
// The function will then return the removed item if the user requires it's use
func (q *ItemQueue) RemoveAtIndex(i int) *Item {
	var item Item
	q.secure(func() {
		item = q.items[i]
		q.items = append(q.items[:i], q.items[i+1:]...)
	})
	return &item
}

// q.Contains(Item) -> None
// The Contains() function will scheck whether the provided ItemQueue contains
//
//	the given Item (_item)
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
//
//	through said ItemQueue and remove the given Item (_item)
func (q *ItemQueue) Remove(item Item) {
	q.secure(func() {
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
	q.secure(func() {
		q.items = append(q.items, i)
	})
}

// q.Get() -> Item
// The Get() function will append the first item of the ItemQueue to the back of the slice
//
//	then remove it from the front
//
// The function returns the first item of the ItemQueue
func (q *ItemQueue) Get() *Item {
	var item Item
	q.secure(func() {
		item = q.items[0]
		q.items = append(q.items, q.items[0])
		q.items = q.items[1:]
	})
	return &item
}

// q.Grab() -> Item
// The Grab() function will return the first item of the ItemQueue then
//
//	remove it from said ItemQueue
func (q *ItemQueue) Grab() *Item {
	var item Item
	q.secure(func() {
		item = q.items[0]
		q.items = q.items[1:]
	})
	return &item
}

// q.Clear() -> None
// The Clear() function will secure the queue then remove all of its items
func (q *ItemQueue) Clear() {
	q.secure(func() {
		q.items = []Item{}
	})
}

// q.Show() -> []Item
// The Show() function will return the ItemQueue's items
func (q *ItemQueue) Show() []Item {

	// Lock Reading
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Return the queue items
	return q.items
}

// q.GetAtIndex(index integer) -> Item
// The GetAtIndex() function is used to return an item at the provided index of the ItemQueue
func (q *ItemQueue) GetAtIndex(i int) Item {

	// Lock Reading
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Return the item at the specific index
	return q.items[i]
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
