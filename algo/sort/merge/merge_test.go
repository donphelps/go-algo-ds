package merge_test

import (
	gosort "sort"
	"testing"

	"github.com/donphelps/go-algo-ds/algo/sort/merge"
	"github.com/donphelps/go-algo-ds/algo/sort/testdata"
)

func TestMergeSort(t *testing.T) {
	i := testdata.RandomUnsortedInts(1000, 99999)

	sorted := merge.MergeSort(i)

	if !gosort.IntsAreSorted(sorted) {
		t.Errorf("MergeSort result is not sorted. %v", i)
	}
}

func TestMergeSortMulti(t *testing.T) {
	var i = testdata.RandomUnsortedInts(1000, 99999)

	m := merge.New(63)
	sorted := m.MergeSortMulti(i)

	if !gosort.IntsAreSorted(sorted) {
		t.Errorf("MergeSort result is not sorted. %v", i)
	}
}
