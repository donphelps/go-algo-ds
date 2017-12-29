package sort_test

import (
	"go-algo-ds/algo/sort"
	gosort "sort"
	"testing"
)

func TestBubble(t *testing.T) {
	var i = randomUnsortedInts(1000)

	sort.BubbleSort(i)

	if !gosort.IntsAreSorted(i) {
		t.Errorf("BubbleSort result is not sorted. %v", i)
	}
}

func TestBubbleSortOptimized(t *testing.T) {
	var i = randomUnsortedInts(1000)

	sort.BubbleSortOptimized(i)

	if !gosort.IntsAreSorted(i) {
		t.Errorf("BubbleSortOptimized result is not sorted. %v", i)
	}
}
