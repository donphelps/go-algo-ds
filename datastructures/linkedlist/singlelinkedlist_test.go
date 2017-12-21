package linkedlist_test

import (
	"go-algo-ds/datastructures/linkedlist"
	"testing"
)

func TestInsertHead(t *testing.T) {
	var s linkedlist.SingleLinkedList

	s.InsertHead(1)
	if s.String() != "1 " {
		t.Errorf("InsertHead into empty list fail. Expected '1 ', got '%v'", s.String())
	}
}

func TestInsert(t *testing.T) {
	var s linkedlist.SingleLinkedList

	s.Insert(3)
	s.Insert(4)
	s.Insert(5)
	s.InsertHead(2)
	if s.String() != "2 3 4 5 " {
		t.Errorf("Insert fail. Expected '2 3 4 5 ', got '%s'", s.String())
	}
}

func TestCount(t *testing.T) {
	var s linkedlist.SingleLinkedList

	s.Insert(3)
	s.Insert(4)
	s.Insert(5)
	s.InsertHead(2)

	c := s.Count()
	if c != 4 {
		t.Errorf("Count fail. Expected 4, got %v", c)
	}
}

func TestClear(t *testing.T) {
	var s linkedlist.SingleLinkedList

	s.InsertHead(1)
	s.Clear()
	c := s.Count()
	if c != 0 {
		t.Errorf("Clear list fail. Expected '1 ', got '%v'", s.String())
	}
}

func TestSearch(t *testing.T) {
	var s linkedlist.SingleLinkedList

	s.Insert(3)
	s.Insert(4)
	s.Insert(5)
	s.InsertHead(2)

	three := s.Search(3)
	if three == nil {
		t.Errorf("Search fail. Expected non-nil node, got nil")
	}

	onehundred := s.Search(100)
	if onehundred != nil {
		t.Errorf("Search fail. Expected nil node, got non-nil")
	}
}

func TestDelete(t *testing.T) {
	var s linkedlist.SingleLinkedList

	s.Insert(1)
	s.Insert(2)
	s.Insert(3)
	s.Insert(4)
	s.Insert(5)
	s.Delete(1) // del head
	s.Delete(3) // del mid
	s.Delete(5) // del tail
	if s.String() != "2 4 " {
		t.Errorf("Delete fail. Expected '2 4 ', got '%s'", s.String())
	}

	s.Delete(2)
	s.Delete(4)
	if s.Count() != 0 {
		t.Errorf("Delete fail. Expected count of 0, got %v", s.Count())
	}
}
