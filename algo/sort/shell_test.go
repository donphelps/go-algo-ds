package sort_test

import (
	"go-algo-ds/algo/sort"
	gosort "sort"
	"testing"
)

func TestShellSort(t *testing.T) {
	var i = randomUnsortedInts(1000, 99999)

	sort.ShellSort(i)

	if !gosort.IntsAreSorted(i) {
		t.Errorf("ShellSort result is not sorted. %v", i)
	}
}

func TestShellSortCiura(t *testing.T) {
	var i = randomUnsortedInts(1000, 99999)

	sort.ShellSortCiura(i)

	if !gosort.IntsAreSorted(i) {
		t.Errorf("ShellSort result is not sorted. %v", i)
	}
}
