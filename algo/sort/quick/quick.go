package quick

import "math/rand"

// Sort initiates a quicksort on the supplied int slice and returns a sorted int slice.
// From https://gist.github.com/vderyagin/4161347
//
// Time complexity: O(n^2)
// Space complexity: O(n)
func Sort(slice []int) []int {
	length := len(slice)

	if length <= 1 {
		sliceCopy := make([]int, length)
		copy(sliceCopy, slice)
		return sliceCopy
	}

	m := slice[rand.Intn(length)]

	less := make([]int, 0, length)
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range slice {
		switch {
		case item < m:
			less = append(less, item)
		case item == m:
			middle = append(middle, item)
		case item > m:
			more = append(more, item)
		}
	}

	less, more = Sort(less), Sort(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}
