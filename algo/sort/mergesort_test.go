package sort_test

import (
	"go-algo-ds/algo/sort"
	"math/rand"
	"testing"
	"time"
)

func randomUnsortedInts(count int) []int {
	slice := make([]int, count, count)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		slice[i] = rand.Intn(99999) - rand.Intn(99999)
	}

	return slice
}

func TestMergeSort(t *testing.T) {
	goodSort := true
	i := randomUnsortedInts(100)

	sorted := sort.MergeSort(i)

	if len(sorted) > 2 {
		for d := 0; d <= len(sorted)-2; d++ {
			if sorted[d] >= sorted[d+1] {
				goodSort = false
			}
		}
	}

	if !goodSort {
		t.Errorf("MergeSort() result was %v", sorted)
	}
}
