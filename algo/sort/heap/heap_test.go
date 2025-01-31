package heap_test

import (
	gosort "sort"
	"testing"

	"github.com/donphelps/go-algo-ds/algo/sort/heap"
	"github.com/donphelps/go-algo-ds/algo/sort/testdata"
)

func TestHeap(t *testing.T) {
	var i = testdata.RandomUnsortedInts(1000, 99999)

	i = heap.Sort(i)

	if !gosort.IntsAreSorted(i) {
		t.Errorf("HeapSort result is not sorted. %v", i)
	}
}
