package sort_test

import (
	"go-algo-ds/algo/sort"
	gosort "sort"
	"testing"
)

func TestInsertion(t *testing.T) {
	var i = randomUnsortedInts(1000, 99999)

	sort.InsertionSort(i)

	if !gosort.IntsAreSorted(i) {
		t.Errorf("InsertionSort result is not sorted. %v", i)
	}
}
