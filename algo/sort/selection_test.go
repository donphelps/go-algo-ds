package sort_test

import (
	"fmt"
	"go-algo-ds/algo/sort"
	"testing"
)

func TestSelection(t *testing.T) {
	var i = randomUnsortedInts(1000)
	var d int

	sort.SelectionSort(i)

	fmt.Println(i)
	if len(i) > 2 {
		for d = 0; d <= len(i)-2; d++ {
			if i[d] > i[d+1] {
				t.Errorf("SelectionSort result had adjacent values: (%v, %v)", i[d], i[d+1])
				break
			}
		}
	}
}
