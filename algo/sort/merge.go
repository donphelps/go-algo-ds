package sort

import "sync"

// MergeSort initiates a merge sort on the supplied int slice and returns a sorted int slice.
//
// Time complexity: O(n log(n))
// Space complexity: O(n)
func MergeSort(ints []int) []int {
	if len(ints) < 2 {
		return ints
	}

	mid := (len(ints)) / 2

	return merge(MergeSort(ints[:mid]), MergeSort(ints[mid:]))
}

func merge(left, right []int) []int {
	size, leftIndex, rightIndex := len(left)+len(right), 0, 0
	merged := make([]int, size, size)

	for i := 0; i < size; i++ {
		if leftIndex > len(left)-1 && rightIndex <= len(right)-1 { // left is done, right is not.
			merged[i] = right[rightIndex]
			rightIndex++
		} else if rightIndex > len(right)-1 && leftIndex <= len(left)-1 { // right is done, left is not.
			merged[i] = left[leftIndex]
			leftIndex++
		} else if left[leftIndex] < right[rightIndex] { // next left value less than next right value.
			merged[i] = left[leftIndex]
			leftIndex++
		} else {
			merged[i] = right[rightIndex]
			rightIndex++
		}
	}

	return merged
}

var sem chan struct{}

// MergeSortMulti uses goroutines to initiates a merge sort on the supplied slice and returns a sorted int slice.
// From https://medium.com/@_orcaman/when-too-much-concurrency-slows-you-down-golang-9c144ca305a
func MergeSortMulti(ints []int, numGoRoutines int) []int {
	sem = make(chan struct{}, numGoRoutines)
	return mergeSortMulti(ints)
}

func mergeSortMulti(ints []int) []int {
	if len(ints) <= 1 {
		return ints
	}

	n := len(ints) / 2

	wg := sync.WaitGroup{}
	wg.Add(2)

	var l []int
	var r []int

	select {
	case sem <- struct{}{}:
		go func() {
			l = mergeSortMulti(ints[:n])
			<-sem
			wg.Done()
		}()
	default:
		l = MergeSort(ints[:n])
		wg.Done()
	}

	select {
	case sem <- struct{}{}:
		go func() {
			r = mergeSortMulti(ints[n:])
			<-sem
			wg.Done()
		}()
	default:
		r = MergeSort(ints[n:])
		wg.Done()
	}

	wg.Wait()
	return merge(l, r)
}
