package stack

import (
	"bytes"
	"strconv"
)

// Stack implements the stack data structure for ints.
type Stack struct {
	elements []int
}

// New returns a pointer to a new Stack struct.
func New() *Stack {
	return &Stack{}
}

// Clear removes all elements from the stack.
func (s *Stack) Clear() {
	s.elements = []int{}
}

// Len returns the number of elements in a stack.
func (s *Stack) Len() int {
	return len(s.elements)
}

// Peek returns the next element in the stack without removing it.
func (s *Stack) Peek() int {
	return s.elements[len(s.elements)-1]
}

// Pop returns the next element and removes it from the stack.
func (s *Stack) Pop() int {
	r := s.elements[len(s.elements)-1]
	s.elements = s.elements[0 : len(s.elements)-1]
	return r
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

// Contains checks if the element is in the stack.
func (s *Stack) Contains(x int) bool {
	for _, e := range s.elements {
		if e == x {
			return true
		}
	}

	return false
}

func (s *Stack) String() string {
	var buffer bytes.Buffer

	for _, e := range s.elements {
		buffer.WriteString(strconv.Itoa(e) + " ")
	}

	return buffer.String()
}
