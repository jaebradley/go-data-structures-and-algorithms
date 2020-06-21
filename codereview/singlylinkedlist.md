# Implementing A Simple Singly Linked List

I'm trying to learn Go and wanted to implement various data structures and algorithms from scratch.

I wanted to try and implement a singly linked list without using slices.

In the sake of keeping things simple, I am not taking into account concurrent access to this list and am exposing a very limited API.

My focus is on adhering to Go best principles and style (as well as logical correctness, of course).

One of the things I struggled with was how to design tests for the `SinglyLinkedList` `struct`, particularly around naming test cases.

I saw [a recommendation from the `testing` Go documentation](https://golang.org/pkg/testing/#hdr-Examples) as well as [this discussion on `StackOverflow`](https://stackoverflow.com/questions/155436/unit-test-naming-best-practices).

Another thing I had questions about was how to test generic values that could be inputted into the list.

I used `interface{}` to represent the values that could be inserted into the list to avoid restricting the list values to a given type - for this type of API, what is the best practice around testing?

Is it to create an array of possible input types and to iterate over that array for each test case?

## Implementation

```golang
package main

type node struct {
	value interface{}
	next  *node
}

// SinglyLinkedList represents a basic singly linked list data structure
type SinglyLinkedList struct {
	size int
	head *node
	tail *node
}

// IsEmpty returns true if list has no elements and false if list has elements
func (list SinglyLinkedList) IsEmpty() (isEmpty bool) {
	return list.size == 0
}

// Push adds an element with the specified value to the end of the list - it will return true
func (list *SinglyLinkedList) Push(value interface{}) (pushed bool) {
	next := &node{value: value}

	if list.IsEmpty() {
		list.head = next
	} else {
		list.tail.next = next
	}

	list.tail = next
	list.size++

	return true
}

// Peek returns the value at the head of the list, but does not remove the element at the head of the list
func (list SinglyLinkedList) Peek() (value interface{}, ok bool) {
	if list.IsEmpty() {
		return nil, false
	}

	return list.head.value, true
}

// Pop returns the value at the head of the list, and removes the element at the head of the list
func (list *SinglyLinkedList) Pop() (value interface{}, ok bool) {
	if list.IsEmpty() {
		return nil, false
	}

	value = list.head.value
	list.head = list.head.next

	if list.size == 1 {
		list.tail = nil
	}

	list.size--

	return value, true
}
```

## Tests

```golang
package main

import "testing"

func TestSinglyLinkedList_IsEmpty_IsTrueWhenInitialized(t *testing.T) {
	list := SinglyLinkedList{}
	if list.IsEmpty() != true {
		t.Errorf("Expected list to be empty")
	}
}

func TestSinglyLinkedList_IsEmpty_IsFalseWhenElementIsPushed(t *testing.T) {
	list := SinglyLinkedList{}
	list.Push(1)
	if list.IsEmpty() == true {
		t.Errorf("Expected list to not be empty")
	}
}

func TestSinglyLinkedList_IsEmpty_IsTrueWhenListHasElementPushedThenPopped(t *testing.T) {
	list := SinglyLinkedList{}
	list.Push(1)
	list.Pop()
	if list.IsEmpty() != true {
		t.Errorf("Expected list to be empty")
	}
}
func TestSinglyLinkedList_IsEmpty_IsTrueWhenListHasElementPushedTwiceThenPoppedTwice(t *testing.T) {
	list := SinglyLinkedList{}
	list.Push(1)
	list.Push(2)
	list.Pop()
	list.Pop()
	if list.IsEmpty() != true {
		t.Errorf("Expected list to be empty")
	}
}

func TestSinglyLinkedList_IsEmpty_IsFalseWhenListHasElementPushedTwiceThenPoppedOnce(t *testing.T) {
	list := SinglyLinkedList{}
	list.Push(1)
	list.Push(2)
	list.Pop()
	if list.IsEmpty() != false {
		t.Errorf("Expected list not to be empty")
	}
}

func TestSinglyLinkedList_Push_ReturnsTrue(t *testing.T) {
	list := SinglyLinkedList{}

	if list.Push(100) != true {
		t.Errorf("Expected Push to return true")
	}
}

func TestSinglyLinkedList_Peek_ReturnFalseOkValueWhenListIsEmpty(t *testing.T) {
	list := SinglyLinkedList{}

	_, ok := list.Peek()

	if ok != false {
		t.Errorf("Expected peeking empty list to not be ok")
	}
}

func TestSinglyLinkedList_Peek_ReturnsNilValueWhenListIsEmpty(t *testing.T) {
	list := SinglyLinkedList{}

	value, _ := list.Peek()

	if value != nil {
		t.Errorf("Expected value from peeking empty list to be nil and not %v", value)
	}
}

func TestSinglyLinkedList_Peek_ReturnTrueOkValueWhenListIsNotEmpty(t *testing.T) {
	list := SinglyLinkedList{}
	list.Push(1)
	_, ok := list.Peek()

	if ok != true {
		t.Errorf("Expected peeking single element list to be ok")
	}
}

func TestSinglyLinkedList_Peek_ReturnsFirstPushedValueWhenListHasBeenPushedToOnce(t *testing.T) {
	list := SinglyLinkedList{}
	list.Push(1)
	value, _ := list.Peek()

	if value != 1 {
		t.Errorf("Expected value to be 1 instead of %d", value)
	}
}

func TestSinglyLinkedList_Peek_ReturnsFirstPushedValueWhenListHasBeenPushedToTwice(t *testing.T) {
	list := SinglyLinkedList{}
	list.Push(1)
	list.Push(2)

	value, _ := list.Peek()

	if value != 1 {
		t.Errorf("Expected value to be 1 instead of %d", value)
	}
}

func TestSinglyLinkedList_Pop_ReturnFalseOkValueWhenListIsEmpty(t *testing.T) {
	list := SinglyLinkedList{}

	_, ok := list.Pop()

	if ok != false {
		t.Errorf("Expected peeking empty list to not be ok")
	}
}

func TestSinglyLinkedList_Pop_ReturnsNilValueWhenListIsEmpty(t *testing.T) {
	list := SinglyLinkedList{}

	value, _ := list.Pop()

	if value != nil {
		t.Errorf("Expected value from peeking empty list to be nil and not %v", value)
	}
}

func TestSinglyLinkedList_Pop_ReturnTrueOkValueWhenListIsNotEmpty(t *testing.T) {
	list := SinglyLinkedList{}
	list.Push(1)
	_, ok := list.Pop()

	if ok != true {
		t.Errorf("Expected peeking single element list to be ok")
	}
}

func TestSinglyLinkedList_Pop_ReturnsFirstPushedValueWhenListHasBeenPushedToOnce(t *testing.T) {
	list := SinglyLinkedList{}
	list.Push(1)
	value, _ := list.Pop()

	if value != 1 {
		t.Errorf("Expected value to be 1 instead of %d", value)
	}
}

func TestSinglyLinkedList_Pop_ReturnsFirstPushedValueWhenListHasBeenPushedToTwice(t *testing.T) {
	list := SinglyLinkedList{}
	list.Push(1)
	list.Push(2)

	value, _ := list.Pop()

	if value != 1 {
		t.Errorf("Expected value to be 1 instead of %d", value)
	}
}
```
