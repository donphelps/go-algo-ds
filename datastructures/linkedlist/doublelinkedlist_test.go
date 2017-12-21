package linkedlist_test

import (
	"go-algo-ds/datastructures/linkedlist"
	"testing"
)

func sampleDoubleLinkedList() linkedlist.DoubleLinkedList {
	var d linkedlist.DoubleLinkedList

	for i := 1; i <= 5; i++ {
		d.Insert(i)
	}

	return d
}

func TestDoubleInsertHead(t *testing.T) {
	var d linkedlist.DoubleLinkedList

	d.InsertHead(2)
	if d.String() != "2 " {
		t.Errorf("InsertHead into empty list fail. Expected '2 ', got '%v'.", d.String())
	}

	d.InsertHead(1)
	if d.String() != "1 2 " {
		t.Errorf("InsertHead into empty list fail. Expected '1 2 ', got '%v'.", d.String())
	}
}

func TestDoubleInsert(t *testing.T) {
	d := sampleDoubleLinkedList()

	d.Insert(6)
	if d.String() != "1 2 3 4 5 6 " {
		t.Errorf("Insert fail. Expected '1 2 3 4 5 ', got '%s'.", d.String())
	}
}

func TestDoubleCount(t *testing.T) {
	d := sampleDoubleLinkedList()

	c := d.Count()
	if c != 5 {
		t.Errorf("Count fail. Expected 5, got %v.", c)
	}
}

func TestDoubleClear(t *testing.T) {
	d := sampleDoubleLinkedList()

	d.Clear()
	c := d.Count()
	if c != 0 {
		t.Errorf("Clear fail. Expected count of 0, got %v.", c)
	}
}

func TestDoubleSearch(t *testing.T) {
	d := sampleDoubleLinkedList()

	three := d.Search(3)
	if three == nil {
		t.Errorf("Search fail. Expected non-nil node, got nil.")
	}

	onehundred := d.Search(100)
	if onehundred != nil {
		t.Errorf("Search fail. Expected nil node, got non-nil.")
	}
}

func TestDoubleSearchReverse(t *testing.T) {
	var d linkedlist.DoubleLinkedList

	if d.SearchReverse(1) != nil {
		t.Errorf("SearchReverse an empty list fail. Expected false, got %v.", (d.SearchReverse != nil))
	}

	d = sampleDoubleLinkedList()

	three := d.SearchReverse(3)
	if three == nil {
		t.Errorf("SearchReverse fail. Expected non-nil node, got nil.")
	}

	onehundred := d.SearchReverse(100)
	if onehundred != nil {
		t.Errorf("Search fail. Expected nil node, got non-nil.")
	}
}

func TestDoubleDelete(t *testing.T) { //[] break out into individual tests
	d := sampleDoubleLinkedList()

	d.Delete(1) // del head
	d.Delete(3) // del mid
	d.Delete(5) // del tail
	if d.String() != "2 4 " {
		t.Errorf("Delete fail. Expected '2 4 ', got '%s'", d.String())
	}
}
