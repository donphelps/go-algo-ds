package main

import (
	"fmt"
	"go-algo-ds/algo/sort"
	"math/rand"
	"time"

	"golang.org/x/text/message"
)

/*
[] Redo all comments
[] Rename parameters
[] What sorts to implement?

[] do all functions have their computational cost noted?

[] make sure all big-O's are accurate.
*/

var sampleQty = 1000000
var sampleSize = 9999

func main() {
	fmt.Println()
	defer fmt.Println()

	testSorts(sampleQty)
}

func testSorts(qty int) {
	fmt.Printf("Sample contains %v elements ranging in values +/- %v\n", formatInt(sampleQty), formatInt(sampleSize))

	start := time.Now()
	b := randomUnsortedInts(qty)
	/*o := append([]int(nil), b...)
	s := append([]int(nil), b...)
	i := append([]int(nil), b...)*/
	h := append([]int(nil), b...)
	m := append([]int(nil), b...)
	mm := append([]int(nil), b...)
	sh := append([]int(nil), b...)
	sho := append([]int(nil), b...)
	q := append([]int(nil), b...)
	elapsed := time.Since(start)
	fmt.Printf("Samples generated in %v\n\n", elapsed)

	/*fmt.Println("O(n^2):")
	start = time.Now()
	sort.BubbleSort(b)
	elapsed = time.Since(start)
	fmt.Printf("  bubble complete in %v\n", elapsed)
	// ------------
	start = time.Now()
	sort.BubbleSortOptimized(o)
	elapsed = time.Since(start)
	fmt.Printf("  bubble opt complete in %v\n", elapsed)
	// ------------
	start = time.Now()
	sort.SelectionSort(s)
	elapsed = time.Since(start)
	fmt.Printf("  selection complete in %v\n", elapsed)
	// ------------
	start = time.Now()
	sort.InsertionSort(i)
	elapsed = time.Since(start)
	fmt.Printf("  insertion complete in %v\n", elapsed)*/
	// ------------
	fmt.Println("O(n log(n)):")
	start = time.Now()
	h = sort.HeapSort(h)
	elapsed = time.Since(start)
	fmt.Printf("  heap complete in %v\n", elapsed)
	// ------------
	start = time.Now()
	m = sort.MergeSort(m)
	elapsed = time.Since(start)
	fmt.Printf("  merge complete in %v\n", elapsed)
	// ------------
	start = time.Now()
	mm = sort.MergeSort(mm)
	elapsed = time.Since(start)
	fmt.Printf("  goroutine merge complete in %v\n", elapsed)
	// ------------
	start = time.Now()
	sort.ShellSort(sh)
	elapsed = time.Since(start)
	fmt.Printf("  shell complete in %v\n", elapsed)
	// ------------
	start = time.Now()
	sort.ShellSortCiura(sho)
	elapsed = time.Since(start)
	fmt.Printf("  shell optimized complete in %v\n", elapsed)
	// ------------
	start = time.Now()
	q = sort.QuickSort(q)
	elapsed = time.Since(start)
	fmt.Printf("  quick complete in %v\n", elapsed)
}

func randomUnsortedInts(count int) []int {
	slice := make([]int, count, count)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		slice[i] = rand.Intn(sampleSize) - rand.Intn(sampleSize)
	}

	return slice
}

func formatInt(x int) string {
	ptr := message.NewPrinter(message.MatchLanguage("en"))
	return ptr.Sprintf("%d", x)
}
