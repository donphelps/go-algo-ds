package linkedlist

import (
	"bytes"
	"fmt"
)

type Nodeable interface {
	comparable
}

type node[T Nodeable] struct {
	Value T
	Next  *node[T]
}

// [doc]
type LinkedList[T Nodeable] struct {
	Head *node[T]
	Tail *node[T]
}

// [doc]
func New[T Nodeable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// [doc]
func (l *LinkedList[T]) Clear() {
	l.Head = nil
	l.Tail = nil
}

// [doc]
func (l *LinkedList[T]) Len() int {
	len := 0

	for node := l.Head; node.Next != nil; node = node.Next {
		len++
	}

	return len
}

// [doc]
func (l *LinkedList[T]) InsertFirst(value T) {
	node := &node[T]{
		Value: value,
		Next:  nil,
	}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	node.Next = l.Head
	l.Head = node
}

// [doc]
func (l *LinkedList[T]) InsertLast(value T) {
	node := &node[T]{
		Value: value,
		Next:  nil,
	}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	if l.Tail == nil {
		l.Tail = node
		return
	}

	l.Tail.Next = node
	l.Tail = node
}

// [doc]
func (l *LinkedList[T]) AddAfter(value T) {
	node := &node[T]{
		Value: value,
		Next:  nil,
	}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	if l.Tail == nil {
		l.Tail = node
		return
	}

	l.Tail.Next = node
	l.Tail = node
}

// [doc]
func (l *LinkedList[T]) AddBefore(value T) {
	node := &node[T]{
		Value: value,
		Next:  nil,
	}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	if l.Tail == nil {
		l.Tail = node
		return
	}

	l.Tail.Next = node
	l.Tail = node
}

// [doc]
func (l *LinkedList[T]) Remove(value T) {
	//[i]
}

// [doc]
func (l *LinkedList[T]) RemoveFirst() {
	//[i]
}

// [doc]
func (l *LinkedList[T]) RemoveLast(value T) {
	//[i]
}

// [doc]
func (l *LinkedList[T]) Contains(value T) bool {
	//[i]
	return false
}

// [doc]
func (l *LinkedList[T]) String() string {
	sb := bytes.Buffer{}

	sb.WriteString("[")

	node := node[T]{}
	for node := l.Head; node.Next != nil; node = node.Next {
		sb.WriteString(fmt.Sprintf("%v ", node.Value))
	}
	sb.WriteString(fmt.Sprintf("%v]", node.Value))

	return sb.String()
}
