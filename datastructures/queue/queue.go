package queue

import (
	"bytes"
	"strconv"
)

// Queue implements the queue data structure for ints.
type Queue struct {
	elements []int
}

// New returns a pointer to a new Queue struct.
func New() *Queue {
	return &Queue{}
}

// Clear removes all elements from a queue.
func (i *Queue) Clear() {
	i.elements = []int{}
}

// Len returns the number of elements in a queue.
func (i *Queue) Len() int {
	return len(i.elements)
}

// Peek returns the next element in a queue without removing it.
func (i *Queue) Peek() int {
	return i.elements[0]
}

// Enqueue adds an element to the end of a queue.
func (i *Queue) Enqueue(x int) {
	i.elements = append(i.elements, x)
}

// Dequeue returns the next element and removes it from the queue.
func (i *Queue) Dequeue() int {
	r := i.Peek()
	i.elements = append(i.elements[:0], i.elements[1:]...)

	return r
}

// Contains checks if the element is in the queue.
func (i *Queue) Contains(x int) bool {
	for _, e := range i.elements {
		if e == x {
			return true
		}
	}

	return false
}

func (i *Queue) String() string {
	b := bytes.Buffer{}

	for _, e := range i.elements {
		b.WriteString(strconv.Itoa(e) + " ")
	}

	return b.String()
}
