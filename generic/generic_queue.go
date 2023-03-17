package generic_queue

import "sync"

////////////////////////////////////////////
// Generics Implementation by lil5		  //
// Visit him at: https://github.com/lil5  //
// Thank You!							  //
////////////////////////////////////////////

// Generics Require Go v1.18+

// Generic T
//
//	 The 'T' Type is the type of variables that will be going inside the queue slice
//	 The Generic T is can be declared as any so it is possible to have multiple types
//		   within the Queue Slice
type Item[T any] any

// The 'ItemQueue' Struct contains the []T slice
// This struct holds two keys,
//   - items -> the []T slice
//   - mutex -> the mutex lock which prevents overwrites and data corruption
//     ↳ We use RWMutex instead of Mutex as it's better for majority read slices
type ItemQueue[T any] struct {
	items []T
	mutex *sync.RWMutex
}

// Create() -> *ItemQueue
// The Create() function will return an empty ItemQueue
func Create[T any]() *ItemQueue[T] {
	return &ItemQueue[T]{
		mutex: &sync.RWMutex{},
		items: []T{},
	}
}

// q.Index(index integer) -> *Item
// The RemoveAtIndex() function is used to remove an item at the provided index of the ItemQueue
// The function will then return the removed item if the user requires it's use
func (q *ItemQueue[T]) RemoveAtIndex(i int) T {
	// Lock the queue then unlock once function closes
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// Create a copy of the item
	var copy T = q.items[i]

	// Remove the item at the index
	q.items = append(q.items[:i], q.items[i+1:]...)

	// Returen the copy of the item
	return copy
}

// q.Contains(Item) -> None
// The Contains() function will scheck whether the provided ItemQueue contains
//
//	the given Item (_item)
func (q *ItemQueue[T]) Contains(item any) bool {
	// Lock Reading
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Iterate over the queue
	for i := 0; i < len(q.items); i++ {
		if any(q.items[i]) == item {
			return true
		}
	}
	return false
}

// q.Remove(Item) -> None
// The Remove() function will secure the ItemQueue before iterating
//
//	through said ItemQueue and remove the given Item (_item)
func (q *ItemQueue[T]) Remove(item any) {
	// Lock the queue then unlock once function closes
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// Iterate ove the queue items
	for i := 0; i < len(q.items); i++ {
		if any(q.items[i]) == item {
			q.items = append(q.items[:i], q.items[i+1:]...)
			return
		}
	}
}

// q.Put(Item) -> None
// The Put() function is used to add a new item to the provided ItemQueue
func (q *ItemQueue[T]) Put(i T) {
	// Lock the queue then unlock once function closes
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// Append the new item to the item queue
	q.items = append(q.items, i)
}

// q.Get() -> Item
// The Get() function will append the first item of the ItemQueue to the back of the slice
//
//	then remove it from the front
//
// The function returns the first item of the ItemQueue
func (q *ItemQueue[T]) Get() T {
	// Lock the queue then unlock once function closes
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// Create a copy of the item
	var copy T = q.items[0]

	// Move the item at the front to the back
	q.items = append(q.items, q.items[0])
	q.items = q.items[1:]

	// Return the copy
	return copy
}

// q.Grab() -> Item
// The Grab() function will return the first item of the ItemQueue then
//
//	remove it from said ItemQueue
func (q *ItemQueue[T]) Grab() T {
	// Lock the queue then unlock once function closes
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// Create a copy of the item
	var copy T = q.items[0]

	// Remove it from the queue
	q.items = q.items[1:]

	// Return the copy
	return copy
}

// q.Clear() -> None
// The Clear() function will secure the queue then remove all of its items
func (q *ItemQueue[T]) Clear() {
	// Lock the queue then unlock once function closes
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// Set the queue items to empty
	q.items = []T{}
}

// q.Show() -> *[]Item
// The Show() function will return the ItemQueue's items
func (q *ItemQueue[T]) Show() []T {
	// Lock Reading
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Create a copy of the items
	var copy []T = q.items

	// Return the copy
	return copy
}

// q.GetAtIndex(index integer) -> *Item
// The GetAtIndex() function is used to return an item at the provided index of the ItemQueue
func (q *ItemQueue[T]) GetAtIndex(i int) T {
	// Lock Reading
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Create a copy of the item at the index
	var copy T = q.items[i]

	// Return the copy
	return copy
}

// q.IsEmpty() -> bool
// The IsEmpty() function will return whether the provided ItemQueue contains any Items
func (q *ItemQueue[T]) IsEmpty() bool {
	// Lock Reading
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Return whether queue is empty
	return len(q.items) == 0
}

// q.IsNotEmpty() -> bool
// The IsNotEmpty() function will return whether the provided ItemQueue contains any Items
func (q *ItemQueue[T]) IsNotEmpty() bool {

	// Lock Reading
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Return whether length is greater than 0
	return len(q.items) > 0
}

// q.Size() -> int
// The Size() function will return the length of the ItemQueue slice
func (q *ItemQueue[T]) Size() int {

	// Lock Reading
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// Return the queue length
	return len(q.items)
}
