package sort_test

import (
	"go-algo-ds/algo/sort"
	gosort "sort"
	"testing"
)

func TestSelection(t *testing.T) {
	var i = randomUnsortedInts(1000)

	sort.SelectionSort(i)

	if !gosort.IntsAreSorted(i) {
		t.Errorf("SelectionSort result is not sorted. %v", i)
	}
}
