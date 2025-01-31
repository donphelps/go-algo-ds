package stack

import (
	"errors"
)

type Stack[T comparable] struct {
	head *stackNode[T]
}

type stackNode[T comparable] struct {
	Value T
	Next  *stackNode[T]
}

func New[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Clear() {
	s.head = nil
}

func (s Stack[T]) Len() int {
	if s.head == nil {
		return 0
	}
	if s.head.Next == nil {
		return 1
	}

	n := 0

	for node := s.head; node.Next != nil; node = node.Next {
		n++
	}
	n++

	return n
}

func (s *Stack[T]) Peek() T {
	return s.head.Value
}

func (s *Stack[T]) Pop() (T, error) {
	if s.head == nil {
		defaultValue := new(T)
		return *defaultValue, errors.New("stack is empty")
	}

	val := s.head.Value
	s.head = s.head.Next

	return val, nil
}

func (s *Stack[T]) Push(value T) {
	node := stackNode[T]{
		Value: value,
		Next:  s.head,
	}

	s.head = &node
}

func (s Stack[T]) Contains(value T) bool {
	for node := s.head; node.Next != nil; node = node.Next {
		if node.Value == value {
			return true
		}
	}

	return false
}
