package bubble_test

import (
	gosort "sort"
	"testing"

	"github.com/donphelps/go-algo-ds/algo/sort/bubble"
	"github.com/donphelps/go-algo-ds/algo/sort/testdata"
)

func TestBubble(t *testing.T) {
	var i = testdata.RandomUnsortedInts(1000, 99999)

	bubble.Sort(i)

	if !gosort.IntsAreSorted(i) {
		t.Errorf("BubbleSort result is not sorted. %v", i)
	}
}

func TestBubbleSortOptimized(t *testing.T) {
	var i = testdata.RandomUnsortedInts(1000, 99999)

	bubble.SortOptimized(i)

	if !gosort.IntsAreSorted(i) {
		t.Errorf("BubbleSortOptimized result is not sorted. %v", i)
	}
}
