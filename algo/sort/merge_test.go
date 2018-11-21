package sort_test

import (
	"go-algo-ds/algo/sort"
	gosort "sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var i = randomUnsortedInts(1000, 99999)

	sorted := sort.MergeSort(i)

	if !gosort.IntsAreSorted(sorted) {
		t.Errorf("MergeSort result is not sorted. %v", i)
	}
}

func TestMergeSortMulti(t *testing.T) {
	var i = randomUnsortedInts(1000, 99999)

	sorted := sort.MergeSortMulti(i, 64)

	if !gosort.IntsAreSorted(sorted) {
		t.Errorf("MergeSort result is not sorted. %v", i)
	}
}
