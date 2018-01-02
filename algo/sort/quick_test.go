package sort_test

import (
	"go-algo-ds/algo/sort"
	gosort "sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var i = randomUnsortedInts(1000)

	sorted := sort.QuickSort(i)

	if !gosort.IntsAreSorted(sorted) {
		t.Errorf("QuickSort result is not sorted. %v", i)
	}
}
