package quick_test

import (
	gosort "sort"
	"testing"

	"github.com/donphelps/go-algo-ds/algo/sort/quick"
	"github.com/donphelps/go-algo-ds/algo/sort/testdata"
)

func TestQuickSort(t *testing.T) {
	var i = testdata.RandomUnsortedInts(1000, 99999)

	sorted := quick.Sort(i)

	if !gosort.IntsAreSorted(sorted) {
		t.Errorf("QuickSort result is not sorted. %v", i)
	}
}
