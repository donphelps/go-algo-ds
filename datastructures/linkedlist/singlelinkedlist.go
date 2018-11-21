package linkedlist

import (
	"bytes"
	"fmt"
)

// SingleNode is a node for use in a SingleLinkedList struct.
type SingleNode struct {
	Value int
	Next  *SingleNode
}

// SingleLinkedList implements the Single Linked List data structure for ints.
type SingleLinkedList struct {
	Head *SingleNode
	Tail *SingleNode
}

// NewSingleNode returns a SingleNode with the supplied int value.
func NewSingleNode(value int) SingleNode {
	return SingleNode{
		value,
		nil,
	}
}

// Insert an integer as the last item in the list.
func (s *SingleLinkedList) Insert(i int) {
	n := NewSingleNode(i)

	if s.Head == nil { // inserting into empty list.
		s.Head = &n
		s.Tail = &n
		return
	}

	s.Tail.Next = &n
	s.Tail = &n
}

// InsertHead inserts an integer as the first item in the list.
func (s *SingleLinkedList) InsertHead(i int) {
	n := NewSingleNode(i)

	if s.Head == nil { // inserting into head position in an empty list.
		s.Head = &n
		s.Tail = &n
		return
	}

	n.Next = s.Head
	s.Head = &n
}

// Search for the integer and return a pointer to the node.
func (s *SingleLinkedList) Search(i int) *SingleNode {
	current := s.Head
	for current != nil && current.Value != i {
		current = current.Next
	}

	return current
}

// Delete the first instance of an integer from the list.
func (s *SingleLinkedList) Delete(i int) {
	current := s.Head
	var last *SingleNode

	if current != nil && current.Value == i {
		s.Head = current.Next
		return
	}

	for current != nil && current.Value != i {
		last = current
		current = current.Next
	}

	if current != nil && last != nil {
		last.Next = current.Next
	}
}

// Clear the list of all entries.
func (s *SingleLinkedList) Clear() {
	s.Head = nil
	s.Tail = nil
}

// Count returns the number of elements in the list.
func (s *SingleLinkedList) Count() int {
	r := 0
	current := s.Head

	for current != nil {
		r++
		current = current.Next
	}

	return r
}

func (s *SingleLinkedList) String() string {
	var buffer bytes.Buffer
	current := s.Head

	for current != nil {
		buffer.WriteString(fmt.Sprintf("%v ", current.Value))
		current = current.Next
	}

	return buffer.String()
}
