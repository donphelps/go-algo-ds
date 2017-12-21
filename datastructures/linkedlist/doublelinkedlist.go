package linkedlist

import (
	"bytes"
	"fmt"
)

// DoubleNode is a node for use in a DoubleLinkedList struct.
type DoubleNode struct {
	Value    int
	Previous *DoubleNode
	Next     *DoubleNode
}

// DoubleLinkedList implements the Double Linked List datastructure for ints.
type DoubleLinkedList struct {
	Head *DoubleNode
	Tail *DoubleNode
}

// NewDoubleNode returns a DoubleNode with the supplied int value. O(1).
func NewDoubleNode(v int) *DoubleNode {
	return &DoubleNode{
		v,
		nil,
		nil,
	}
}

// Insert an integer as the last item in the list. O(1).
func (d *DoubleLinkedList) Insert(i int) {
	n := NewDoubleNode(i)

	if d.Head == nil { // inserting into empty list.
		d.Head = n
		d.Tail = n
		return
	}

	n.Previous = d.Tail
	d.Tail.Next = n
	d.Tail = n
}

// InsertHead inserts an integer as the first item in the list. O(1).
func (d *DoubleLinkedList) InsertHead(i int) {
	n := NewDoubleNode(i)

	if d.Head == nil { // inserting into head position in an empty list.
		d.Head = n
		d.Tail = n
		return
	}

	fmt.Println("particular code reached.")
	n.Next = d.Head
	d.Head = n
}

// Search searches for the integer and return a pointer to the node. O(n).
func (d *DoubleLinkedList) Search(i int) *DoubleNode {
	current := d.Head
	for current != nil && current.Value != i {
		current = current.Next
	}

	return current
}

// SearchReverse searches for the integer starting from the end of the list and return a pointer to the node. O(n).
func (d *DoubleLinkedList) SearchReverse(i int) *DoubleNode {
	if d.Head == nil || d.Tail == nil {
		return nil
	}

	current := d.Tail
	for current != nil && current.Value != i {
		current = current.Previous
	}

	return current
}

// Delete the first instance of an integer from the list. O(n).
func (d *DoubleLinkedList) Delete(i int) {
	current := d.Head
	var last *DoubleNode

	if current != nil && current.Value == i {
		d.Head = current.Next
	}

	for current != nil && current.Value != i {
		last = current
		current = current.Next
	}

	if current != nil && last != nil {
		last.Next = current.Next
	}
}

// Clear the list of all entries. O(1).
func (d *DoubleLinkedList) Clear() {
	d.Head = nil
	d.Tail = nil
}

// Count returns the number of elements in the list. O(n).
func (d *DoubleLinkedList) Count() int {
	r := 0
	current := d.Head

	for current != nil {
		r++
		current = current.Next
	}

	return r
}

func (d *DoubleLinkedList) String() string {
	var buffer bytes.Buffer
	current := d.Head

	for current != nil {
		buffer.WriteString(fmt.Sprintf("%v ", current.Value))
		current = current.Next
	}

	return buffer.String()
}
