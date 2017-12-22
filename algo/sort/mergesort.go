package sort

/*
Computational: O(n log(n))
Memory: O(n)
*/

// MergeSort initiates a Mergesort on the supplied int slice and returns a sorted int slice.
func MergeSort(ints []int) []int {
	if len(ints) < 2 {
		return ints
	}

	mid := (len(ints)) / 2

	return merge(MergeSort(ints[:mid]), MergeSort(ints[mid:]))
}

func merge(left, right []int) []int {
	size, leftIndex, rightIndex := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for i := 0; i < size; i++ {
		if leftIndex > len(left)-1 && rightIndex <= len(right)-1 { // left is done, right is not.
			slice[i] = right[rightIndex]
			rightIndex++
		} else if rightIndex > len(right)-1 && leftIndex <= len(left)-1 { // right is done, left is not.
			slice[i] = left[leftIndex]
			leftIndex++
		} else if left[leftIndex] < right[rightIndex] { // next left value less than next right value.
			slice[i] = left[leftIndex]
			leftIndex++
		} else {
			slice[i] = right[rightIndex]
			rightIndex++
		}
	}

	return slice
}
