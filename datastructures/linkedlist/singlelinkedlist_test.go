package linkedlist_test

import (
	"go-algo-ds/datastructures/linkedlist"
	"testing"
)

func sampleSingleLinkedList() linkedlist.SingleLinkedList {
	var s linkedlist.SingleLinkedList

	for i := 1; i <= 5; i++ {
		s.Insert(i)
	}

	return s
}

func TestSingleInsertHead(t *testing.T) {
	var s linkedlist.SingleLinkedList

	s.InsertHead(1)
	if s.String() != "1 " {
		t.Errorf("InsertHead into empty list fail. Expected '1 ', got '%v'", s.String())
	}
}

func TestSingleInsert(t *testing.T) {
	var s linkedlist.SingleLinkedList

	s.Insert(3)
	s.Insert(4)
	s.Insert(5)
	s.InsertHead(2)
	if s.String() != "2 3 4 5 " {
		t.Errorf("Insert fail. Expected '2 3 4 5 ', got '%s'", s.String())
	}
}

func TestSingleCount(t *testing.T) {
	s := sampleSingleLinkedList()

	c := s.Count()
	if c != 5 {
		t.Errorf("Count fail. Expected 5, got %v", c)
	}
}

func TestSingleClear(t *testing.T) {
	var s linkedlist.SingleLinkedList

	s.InsertHead(1)
	s.Clear()
	c := s.Count()
	if c != 0 {
		t.Errorf("Clear list fail. Expected '1 ', got '%v'", s.String())
	}
}

func TestSingleSearch(t *testing.T) {
	s := sampleSingleLinkedList()

	three := s.Search(3)
	if three == nil {
		t.Errorf("Search fail. Expected non-nil node, got nil")
	}

	onehundred := s.Search(100)
	if onehundred != nil {
		t.Errorf("Search fail. Expected nil node, got non-nil")
	}
}

func TestSingleDelete(t *testing.T) {
	s := sampleSingleLinkedList()

	s.Delete(1) // del head
	s.Delete(3) // del mid
	s.Delete(5) // del tail
	if s.Count() != 2 {
		t.Errorf("Delete fail. Expected count of 2, got '%s'", s.String())
	}

	s.Delete(2)
	s.Delete(4)
	if s.Count() != 0 {
		t.Errorf("Delete fail. Expected count of 0, got %v", s.Count())
	}
}
