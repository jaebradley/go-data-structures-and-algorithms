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
