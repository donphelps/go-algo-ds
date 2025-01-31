package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/donphelps/go-algo-ds/algo/sort/heap"
	"github.com/donphelps/go-algo-ds/algo/sort/merge"
	"github.com/donphelps/go-algo-ds/algo/sort/quick"
	"github.com/donphelps/go-algo-ds/algo/sort/shell"
	"github.com/donphelps/go-algo-ds/datastructures/linkedlist"
	"golang.org/x/text/message"
)

const sampleQty = 10000000
const sampleIntMax = 10000
const maxGoRoutines = 63

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	fmt.Println()
	defer fmt.Println()

	// linkedListDemo()
	benchmarkSorts()
}

func linkedListDemo() {
	ll := linkedlist.New[int]()
	ll.InsertFirst(0)
	ll.InsertLast(1)
	ll.InsertLast(2)

	fmt.Printf("ll = %v\n\n", ll.String())
}

func benchmarkSorts() {
	fmt.Printf("Sample contains %v integer elements ranging in values +/- %v\n", formatInt(sampleQty), formatInt(sampleIntMax))

	start := time.Now()
	b := newSample(sampleQty, sampleIntMax)
	elapsed := time.Since(start)
	fmt.Printf("Samples generated in %v\n", elapsed)
	// o := copySample(b)
	// s := copySample(b)
	// i := copySample(b)
	h := copySample(b)
	m := copySample(b)
	mm := copySample(b)
	sh := copySample(b)
	sho := copySample(b)
	q := copySample(b)
	elapsed = time.Since(start)
	fmt.Printf("Samples copied in %v\n\n", elapsed)

	// start = time.Now()
	// bubble.Sort(b)
	// elapsed = time.Since(start)
	// fmt.Printf("  bubble complete in %v\n", elapsed)

	// start = time.Now()
	// bubble.SortOptimized(o)
	// elapsed = time.Since(start)
	// fmt.Printf("  bubble opt complete in %v\n", elapsed)

	// start = time.Now()
	// selection.Sort(s)
	// elapsed = time.Since(start)
	// fmt.Printf("  selection complete in %v\n", elapsed)

	// start = time.Now()
	// insertion.Sort(i)
	// elapsed = time.Since(start)
	// fmt.Printf("  insertion complete in %v\n", elapsed)

	start = time.Now()
	_ = heap.Sort(h)
	elapsed = time.Since(start)
	fmt.Printf("  heap complete in %v\n", elapsed)

	start = time.Now()
	_ = merge.MergeSort(m)
	elapsed = time.Since(start)
	fmt.Printf("  merge complete in %v\n", elapsed)

	start = time.Now()
	merge := merge.New(maxGoRoutines)
	_ = merge.MergeSortMulti(mm)
	elapsed = time.Since(start)
	fmt.Printf("  goroutine merge complete in %v\n", elapsed)

	start = time.Now()
	shell.Sort(sh)
	elapsed = time.Since(start)
	fmt.Printf("  shell complete in %v\n", elapsed)

	start = time.Now()
	shell.SortCiura(sho)
	elapsed = time.Since(start)
	fmt.Printf("  shell optimized complete in %v\n", elapsed)

	start = time.Now()
	_ = quick.Sort(q)
	elapsed = time.Since(start)
	fmt.Printf("  quick complete in %v\n", elapsed)
}

func newSample(qty, max int) []int {
	sample := make([]int, qty)

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
