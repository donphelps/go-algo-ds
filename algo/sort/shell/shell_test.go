package shell_test

import (
	gosort "sort"
	"testing"

	"github.com/donphelps/go-algo-ds/algo/sort/shell"
	"github.com/donphelps/go-algo-ds/algo/sort/testdata"
)

func TestShellSort(t *testing.T) {
	var i = testdata.RandomUnsortedInts(1000, 99999)

	shell.Sort(i)

	if !gosort.IntsAreSorted(i) {
		t.Errorf("ShellSort result is not sorted. %v", i)
	}
}

func TestShellSortCiura(t *testing.T) {
	var i = testdata.RandomUnsortedInts(1000, 99999)

	shell.SortCiura(i)

	if !gosort.IntsAreSorted(i) {
		t.Errorf("ShellSort result is not sorted. %v", i)
	}
}
