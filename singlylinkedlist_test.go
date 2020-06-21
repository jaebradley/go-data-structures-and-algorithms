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
