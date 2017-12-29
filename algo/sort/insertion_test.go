package sort_test

import (
	"go-algo-ds/algo/sort"
	"testing"
)

func TestInsertion(t *testing.T) {
	var i = randomUnsortedInts(1000)
	var d int

	sort.InsertionSort(i)

	if len(i) > 2 {
		for d = 0; d <= len(i)-2; d++ {
			if i[d] > i[d+1] {
				t.Errorf("InsertionSort result had adjacent values: (%v, %v)", i[d], i[d+1])
				break
			}
		}
	}
}
