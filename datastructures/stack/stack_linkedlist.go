package stack

import (
	"bytes"

	"github.com/donphelps/go-algo-ds/datastructures/linkedlist"
)

type StackAsLinkedList struct {
	list linkedlist.SingleLinkedList
}

type IntNode struct {
	Value int
	Next  *IntNode
}

// New returns a pointer to a new Stack struct.
func NewStackAsLinkedList() *StackAsLinkedList {
	return &StackAsLinkedList{}
}

// Clear removes all elements from the stack.
func (s *StackAsLinkedList) Clear() {
	s.list.Head = nil
	s.list.Tail = nil
}

// Len returns the number of elements in a stack.
func (s *StackAsLinkedList) Len() int {
	return s.list.Count()
}

// Peek returns the next element in the stack without removing it.
func (s *StackAsLinkedList) Peek() int {
	return s.list.Tail.Value
}

// Pop returns the next element and removes it from the stack.
func (s *StackAsLinkedList) Pop() int {
	//[i]
	return 0
}

// Push adds an element to the top of the stack.
func (s *StackAsLinkedList) Push(value int) {
	//[i]
}

// Contains checks if the element is in the stack.
func (s *StackAsLinkedList) Contains(x int) bool {
	//[i]

	return false
}

func (s *StackAsLinkedList) String() string {
	var buffer bytes.Buffer

	return buffer.String()
}
