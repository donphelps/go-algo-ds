package sort

// BubbleSort initiates a bubble sort on the supplied int slice.
// Time complexity: O(n^2)
// Space complexity: no additional
func BubbleSort(ints []int) {
	var size = len(ints)

	for i := 0; i < size; i++ {
		idx := 0

		for idx+1 < size {
			if ints[idx] > ints[idx+1] {
				ints[idx], ints[idx+1] = ints[idx+1], ints[idx]
			}

			idx++
		}
	}
}

// BubbleSortOptimized initiates an optimized bubble sort on the supplied int slice.
func BubbleSortOptimized(ints []int) {
	// Don't compare numbers that are already in their final position.
	// Stop sorting if our list is already sorted.
	var size = len(ints)
	for i := 0; i < size; i++ {
		if !bubblepass(ints, i) {
			return
		}
	}
}

func bubblepass(ints []int, completePasses int) bool {
	var size = len(ints)
	var idx = 0
	var swapped = false

	for idx+1 < (size - completePasses) {
		if ints[idx] > ints[idx+1] {
			ints[idx], ints[idx+1] = ints[idx+1], ints[idx]
			swapped = true
		}

		idx++
	}

	return swapped
}
