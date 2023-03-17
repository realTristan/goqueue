package goqueue

// Import sync package for mutexes
import "sync"

// Type Item interface{}
//
//	 The 'Item' Type is the type of variables that will be going inside the queue slice
//	 The Item is declared as interface so it is possible to have multiple types
//		   within the Queue Slice
type Item interface{}

// type ItemQueue struct
//
//	 The 'ItemQueue' Struct contains the []'Type Item interface{}' slice
//	 This struct holds two keys,
//	    - items -> the []'Type Item interface{}' slice
//	    - mutex -> the mutex lock which prevents overwrites and data corruption
//				  ↳ We use RWMutex instead of Mutex as it's better for majority read slices
type ItemQueue struct {
	items []Item
	mutex *sync.RWMutex
}

// The Create() function is used to create
// a new, empty ItemQueue
func Create() *ItemQueue {
	return &ItemQueue{
		mutex: &sync.RWMutex{},
		items: []Item{},
	}
}

// The RemoveAtIndex() function is used to remove an
// item at the provided index of the ItemQueue
//
// Returns the removed item
func (q *ItemQueue) RemoveAtIndex(i int) Item {
	// Mutex Locking/Unlocking
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// Create a copy of the item
	var copy Item = q.items[i]

	// Remove the item at the specific index
	q.items = append(q.items[:i], q.items[i+1:]...)

	// Return the copied item
	return copy
}

// The Contains() function will check whether
// the provided ItemQueue contains the given item
func (q *ItemQueue) Contains(item Item) bool {
	// Mutex Locking/Unlocking
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
// over the queue items, removing the given Item (_item)
func (q *ItemQueue) Remove(item Item) {
	// Mutex Locking/Unlocking
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// Iterate over the queue items
	for i := 0; i < len(q.items); i++ {
		if q.items[i] == item {
			q.items = append(q.items[:i], q.items[i+1:]...)
			return
		}
	}
}

// The Put() function is used to add a new
// item to the provided ItemQueue
func (q *ItemQueue) Put(item Item) {
	// Mutex Locking/Unlocking
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// Append the item to the queue
	q.items = append(q.items, item)
}

// The Get() function will append the first
// item of the ItemQueue to the front of the
// slice then remove it from the back
func (q *ItemQueue) Get() Item {
	// Mutex Locking/Unlocking
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// Create a copy of the first item of the queue
	var copy Item = q.items[0]

	// Move the first item of the queue to the front
	q.items = append(q.items, q.items[0])
	q.items = q.items[1:]

	// Return the item
	return copy
}

// The Grab() function will return the first item of the
// queue items slikce then remove it from said slice
func (q *ItemQueue) Grab() Item {
	// Mutex Locking/Unlocking
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// Get the first item of the queue
	var item Item = q.items[0]

	// And remove it
	q.items = q.items[1:]

	// Return the item
	return item
}

// The Clear() function will secure the
// queue then remove all of its items
func (q *ItemQueue) Clear() {
	// Mutex Locking/Unlocking
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// Remove all items from the queue
	q.items = []Item{}
}

// The Show() function will return the ItemQueue's items
func (q *ItemQueue) Show() []Item {
	// Mutex Locking/Unlocking
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Create a copy of the queue items
	var copy []Item = q.items

	// Return the copy
	return copy
}

// The GetAtIndex() function is used to return an item
// at the provided index of the ItemQueue
func (q *ItemQueue) GetAtIndex(i int) Item {
	// Mutex Locking/Unlocking
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Create a copy of the item at the index
	var copy Item = q.items[i]

	// Return the copy
	return copy
}

// The IsEmpty() function will return whether the
// provided ItemQueue contains any Items
func (q *ItemQueue) IsEmpty() bool {

	// Mutex Locking/Unlocking
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Return whether queue is empty
	return len(q.items) == 0
}

// The IsNotEmpty() function will return whether
// the provided ItemQueue contains any Items
func (q *ItemQueue) IsNotEmpty() bool {
	// Mutex Locking/Unlocking
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Return whether length is greater than 0
	return len(q.items) > 0
}

// The Size() function will return the
// length of the ItemQueue slice
func (q *ItemQueue) Size() int {
	// Mutex Locking/Unlocking
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Return the queue length
	return len(q.items)
}
