package selection

// Sort initiates a selection sort on the supplied int slice.
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func Sort(ints []int) {
	size := len(ints)
	for i := 0; i < size; i++ {
		min := i
		for j := i + 1; j < size; j++ {
			if ints[j] < ints[min] {
				min = j
			}
		}
		ints[i], ints[min] = ints[min], ints[i]
	}
}
