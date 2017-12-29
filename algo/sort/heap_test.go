package sort_test

import (
	"go-algo-ds/algo/sort"
	gosort "sort"
	"testing"
)

func TestHeap(t *testing.T) {
	i := randomUnsortedInts(1000)

	i = sort.HeapSort(i)

	if !gosort.IntsAreSorted(i) {
		t.Errorf("HeapSort result is not sorted. %v", i)
	}
}