package linkedlist_test

import (
	"testing"

	"github.com/donphelps/go-algo-ds/datastructures/linkedlist"
)

//[]
// func sampleList() *linkedlist.LinkedList[int] {
// 	ll := linkedlist.New[int]()

// 	ll.InsertFirst(0)
// 	for i := 1; i <= 5; i++ {
// 		ll.InsertLast(i)
// 	}

// 	return ll

// }

func TestInsertHead(t *testing.T) {
	ll := linkedlist.New[int]()
	ll.InsertFirst(0)
	ll.InsertFirst(1)

	if ll.Head.Value != 1 && ll.Head.Next.Value != 0 && ll.Tail.Value != 0 {
		t.Errorf("InsertHead fail.")
	}
}

func TestInsertTail(t *testing.T) {
	ll := linkedlist.New[int]()
	ll.InsertLast(0)

	if ll.Head == nil || ll.Tail == nil || ll.Head.Value != 0 || ll.Tail.Value != 0 {
		t.Errorf("AppendTail on empty linkedlist fail.")
	}

	ll.InsertLast(1)
	//[test]
}
