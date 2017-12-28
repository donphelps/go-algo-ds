package sort_test

import (
	"go-algo-ds/algo/sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var i = randomUnsortedInts(1000)
	var d int

	sorted := sort.MergeSort(i)

	if len(sorted) > 2 {
		for d = 0; d <= len(sorted)-2; d++ {
			if sorted[d] > sorted[d+1] {
				t.Errorf("MergeSort result had adjacent values: (%v, %v)", sorted[d], sorted[d+1])
				return
			}
		}
	}
}
