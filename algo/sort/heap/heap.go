package heap

// HeapSort initiates a heap sort on the supplied int slice and returns a sorted int slice.
//
// Time complexity: O(n log(n))
// Space complexity: O(1)
func Sort(slice []int) []int {
	h := buildMaxHeap(slice)

	for i := len(h.slice) - 1; i >= 1; i-- {
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.heapSize--
		h.maxHeapify(0)
	}

	return h.slice
}

type maxHeap struct {
	slice    []int
	heapSize int
}

func buildMaxHeap(slice []int) maxHeap {
	h := maxHeap{slice: slice, heapSize: len(slice)}
	for i := len(slice) / 2; i >= 0; i-- {
		h.maxHeapify(i)
	}
	return h
}

func (h maxHeap) maxHeapify(i int) {
	l, r := 2*i+1, 2*i+2
	max := i

	if l < h.size() && h.slice[l] > h.slice[max] {
		max = l
	}
	if r < h.size() && h.slice[r] > h.slice[max] {
		max = r
	}

	if max != i {
		h.slice[i], h.slice[max] = h.slice[max], h.slice[i]
		h.maxHeapify(max)
	}
}

func (h maxHeap) size() int {
	return h.heapSize
}
