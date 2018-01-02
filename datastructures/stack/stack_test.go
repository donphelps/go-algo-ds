package stack_test

import (
	"go-algo-ds/datastructures/stack"
	"testing"
)

func newTestStack(length int) *stack.Stack {
	s := stack.New()

	if length < 0 {
		return nil
	}

	for i := 1; i <= length; i++ {
		s.Push(i)
	}

	return s
}

func TestClear(t *testing.T) {
	s := newTestStack(3)

	s.Clear()
	if s.Len() != 0 {
		t.Errorf("Stack clear fail. Expected len %v, got %v.", 0, s.Len())
	}
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

func TestStackContains(t *testing.T) {
	s := newTestStack(0)

	s.Push(42)
	s.Push(64)
	s.Push(11)
	s.Push(22)

	c := s.Contains(11)
	if !c {
		t.Error("Stack Contains(11) fail. Expected true, got false.", true, false)
	}

	c = s.Contains(99)
	if c {
		t.Error("Stack Contains(99) fail. Expected false, got true.")
	}
}
