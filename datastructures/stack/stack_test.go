package stack_test

import (
	"go-algo-ds/datastructures/stack"
	"testing"
)

func newTestStack(length int) *stack.Stack {
	if length <= 0 {
		return nil
	}

	s := stack.New()

	for i := 1; i <= length; i++ {
		s.Push(i)
	}

	return s
}

func TestPeek(t *testing.T) {
	s := newTestStack(3)

	var i int
	i = s.Peek()
	if i != 3 {
		t.Errorf("Stack peek fail. Expected %v, got %v.", 3, i)
	}
}

func TestLen(t *testing.T) {
	s := newTestStack(3)

	var l int
	l = s.Len()
	if l != 3 {
		t.Errorf("Stack len fail. Expected %v, got %v.", 3, l)
	}
}

func TestPop(t *testing.T) {
	s := newTestStack(3)

	var i int
	i = s.Pop()
	if i != 3 {
		t.Errorf("Stack pop fail. Expected %v, got %v.", 3, i)
	}
}

func TestString(t *testing.T) {
	s := newTestStack(3)

	str := s.String()
	if str != "1 2 3 " {
		t.Errorf("Stack string fail. Expected '%v', got '%v'.", "1 2 3 ", str)
	}
}
