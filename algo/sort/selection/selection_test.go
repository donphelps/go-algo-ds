package selection_test

import (
	gosort "sort"
	"testing"

	"github.com/donphelps/go-algo-ds/algo/sort/selection"
	"github.com/donphelps/go-algo-ds/algo/sort/testdata"
)

func TestSelection(t *testing.T) {
	var i = testdata.RandomUnsortedInts(1000, 99999)

	selection.Sort(i)

	if !gosort.IntsAreSorted(i) {
		t.Errorf("SelectionSort result is not sorted. %v", i)
	}
}
