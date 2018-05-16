package sort

// InsertionSort implements an insertion sort on the supplied int slice.
//
// Time complexity: O(n^2)
// Space complexity: no additional
func InsertionSort(ints []int) {
	var size = len(ints)
	for i := 1; i < size; i++ {
		j := i
		for j > 0 {
			if ints[j-1] > ints[j] {
				ints[j-1], ints[j] = ints[j], ints[j-1]
			}
			j--
		}
	}
}
