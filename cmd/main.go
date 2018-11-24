package main

import (
	"fmt"
	"go-algo-ds/algo/sort"
	"math/rand"
	"runtime"
	"time"

	"golang.org/x/text/message"
)

var sampleQty = 10000000
var sampleIntMax = 10000
var sample []int

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	fmt.Println()
	defer fmt.Println()

	/*originalSample := newSample(sampleQty, sampleIntMax)
	sample := []int{}
	for i := 10; i <= 1000; i += 10 {
		sample = copySample(originalSample)
		start := time.Now()
		sort.MergeSortMulti(sample, i)
		elapsed := time.Since(start)
		log.Printf("%v %v\n", i, elapsed)
	}*/

	benchmarkSorts(sampleQty, sampleIntMax)
}

func benchmarkSorts(qty, max int) {
	fmt.Printf("Sample contains %v integer elements ranging in values +/- %v\n", formatInt(qty), formatInt(max))

	start := time.Now()
	b := newSample(qty, max)
	//o := copySample(b)
	//s := copySample(b)
	//i := copySample(b)
	h := copySample(b)
	m := copySample(b)
	mm := copySample(b)
	sh := copySample(b)
	sho := copySample(b)
	q := copySample(b)
	elapsed := time.Since(start)
	fmt.Printf("Samples generated in %v\n\n", elapsed)

	// O(n^2) examples are too slow to bother with.

	/*fmt.Println("O(n^2):")
	start = time.Now()
	sort.BubbleSort(b)
	elapsed = time.Since(start)
	fmt.Printf("  bubble complete in %v\n", elapsed)

	start = time.Now()
	sort.BubbleSortOptimized(o)
	elapsed = time.Since(start)
	fmt.Printf("  bubble opt complete in %v\n", elapsed)

	start = time.Now()
	sort.SelectionSort(s)
	elapsed = time.Since(start)
	fmt.Printf("  selection complete in %v\n", elapsed)

	start = time.Now()
	sort.InsertionSort(i)
	elapsed = time.Since(start)
	fmt.Printf("  insertion complete in %v\n", elapsed)*/

	fmt.Println("O(n log(n)):")
	start = time.Now()
	h = sort.HeapSort(h)
	elapsed = time.Since(start)
	fmt.Printf("  heap complete in %v\n", elapsed)

	start = time.Now()
	m = sort.MergeSort(m)
	elapsed = time.Since(start)
	fmt.Printf("  merge complete in %v\n", elapsed)

	start = time.Now()
	mm = sort.MergeSortMulti(mm, 100)
	elapsed = time.Since(start)
	fmt.Printf("  goroutine merge complete in %v\n", elapsed)

	start = time.Now()
	sort.ShellSort(sh)
	elapsed = time.Since(start)
	fmt.Printf("  shell complete in %v\n", elapsed)

	start = time.Now()
	sort.ShellSortCiura(sho)
	elapsed = time.Since(start)
	fmt.Printf("  shell optimized complete in %v\n", elapsed)

	start = time.Now()
	q = sort.QuickSort(q)
	elapsed = time.Since(start)
	fmt.Printf("  quick complete in %v\n", elapsed)
}

func newSample(qty, max int) []int {
	sample := make([]int, qty, qty)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < qty; i++ {
		sample[i] = rand.Intn(max) - rand.Intn(max)
	}

	return sample
}

func copySample(sample []int) []int {
	return append([]int(nil), sample...)
}

func formatInt(x int) string {
	ptr := message.NewPrinter(message.MatchLanguage("en"))
	return ptr.Sprintf("%d", x)
}
