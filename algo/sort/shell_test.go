package sort_test

import (
	"go-algo-ds/algo/sort"
	gosort "sort"
	"testing"
)

func TestShellSort(t *testing.T) {
	var i = randomUnsortedInts(1000)

	sort.ShellSort(i)

	if !gosort.IntsAreSorted(i) {
		t.Errorf("ShellSort result is not sorted. %v", i)
	}
}
