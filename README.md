
# GoQueue ![Stars](https://img.shields.io/github/stars/realTristan/goqueue?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/realTristan/goqueue?label=Watchers)

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
	Queue "github.com/realTristan/goqueue/queue"
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

# GoQueue Functions
```go

// Create() -> *ItemQueue
// The Create() function will return an empty ItemQueue
func Create() *ItemQueue {}

// q.Index(index integer) -> *Item
// The RemoveAtIndex() function is used to remove an item at the provided index of the ItemQueue
// The function will then return the removed item if the user requires it's use
func (q *ItemQueue) RemoveAtIndex(i int) *Item {}

// q.Contains(Item) -> None
// The Contains() function will scheck whether the provided ItemQueue contains
//	  the given Item (_item)
func (q *ItemQueue) Contains(item Item) bool {}

// q.Remove(Item) -> None
// The Remove() function will secure the ItemQueue before iterating
//	  through said ItemQueue and remove the given Item (_item)
func (q *ItemQueue) Remove(item Item) {}

// q.Put(Item) -> None
// The Put() function is used to add a new item to the provided ItemQueue
func (q *ItemQueue) Put(i Item) {}

// q.Get() -> Item
// The Get() function will append the first item of the ItemQueue to the back of the slice
//    then remove it from the front
// The function returns the first item of the ItemQueue
func (q *ItemQueue) Get() *Item {}

// q.Grab() -> Item
// The Grab() function will return the first item of the ItemQueue then
//    remove it from said ItemQueue
func (q *ItemQueue) Grab() *Item {}

// q.Clear() -> None
// The Clear() function will secure the queue then remove all of its items
func (q *ItemQueue) Clear() {}

// q.Show() -> *[]Item
// The Show() function will return the ItemQueue's items
func (q *ItemQueue) Show() *[]Item {}

// q.GetAtIndex(index integer) -> *Item
// The GetAtIndex() function is used to return an item at the provided index of the ItemQueue
func (q *ItemQueue) GetAtIndex(i int) *Item {}

// q.IsEmpty() -> bool
// The IsEmpty() function will return whether the provided ItemQueue contains any Items
func (q *ItemQueue) IsEmpty() bool {}

// q.IsNotEmpty() -> bool
// The IsNotEmpty() function will return whether the provided ItemQueue contains any Items
func (q *ItemQueue) IsNotEmpty() bool {}

// q.Size() -> int
// The Size() function will return the length of the ItemQueue slice
func (q *ItemQueue) Size() int {}

```

# License
MIT License

Copyright (c) 2022 Tristan Simpson

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
