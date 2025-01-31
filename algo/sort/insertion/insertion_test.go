package insertion_test

import (
	gosort "sort"
	"testing"

	"github.com/donphelps/go-algo-ds/algo/sort/insertion"
	"github.com/donphelps/go-algo-ds/algo/sort/testdata"
)

func TestInsertion(t *testing.T) {
	var i = testdata.RandomUnsortedInts(1000, 99999)

	insertion.Sort(i)

	if !gosort.IntsAreSorted(i) {
		t.Errorf("InsertionSort result is not sorted. %v", i)
	}
}
