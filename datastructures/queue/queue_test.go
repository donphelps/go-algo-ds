package queue_test

import (
	"go-algo-ds/datastructures/queue"
	"testing"
)

func TestEnqueue(t *testing.T) {
	var q queue.Queue

	q.Enqueue(1)
	q.Enqueue(2)
	l := q.Len()
	if l != 2 {
		t.Errorf("Queue Len fail. Expected %v, got %v.", 2, l)
	}
}

func TestQueueString(t *testing.T) {
	var q queue.Queue

	q.Enqueue(1)
	q.Enqueue(2)

	s := q.String()
	if s != "1 2 " {
		t.Errorf("Queue String fail. Expected '%v', got '%v'.", "1 2 ", s)
	}
}

func TestQueueClear(t *testing.T) {
	var q queue.Queue

	q.Enqueue(1)
	q.Enqueue(2)

	q.Clear()
	l := q.Len()
	if l != 0 {
		t.Errorf("Queue Clear fail. Expected %v, got %v.", 0, l)
	}
}

func TestQueuePeek(t *testing.T) {
	var q queue.Queue

	q.Enqueue(42)
	q.Enqueue(64)
	p := q.Peek()
	if p != 42 {
		t.Errorf("Queue Peek fail. Expected %v, got %v.", 42, p)
	}
}

func TestQueueDequeue(t *testing.T) {
	var q queue.Queue

	q.Enqueue(42)
	q.Enqueue(64)
	d := q.Dequeue()
	if d != 42 {
		t.Errorf("Queue Dequeue fail. Expected %v, got %v.", 42, d)
	}

}

func TestQueueContains(t *testing.T) {
	var q queue.Queue

	q.Enqueue(42)
	q.Enqueue(64)
	q.Enqueue(11)
	q.Enqueue(22)

	c := q.Contains(11)
	if !c {
		t.Error("Queue Contains(11) fail. Expected true, got false.", true, false)
	}

	c = q.Contains(99)
	if c {
		t.Error("Queue Contains(99) fail. Expected false, got true.")
	}
}
