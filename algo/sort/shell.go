package sort

var ciuraGaps = []int{1750, 701, 301, 132, 57, 23, 10, 4, 1}

// ShellSort initiates a shell sort on the supplied int slice.
//
// Time complexity: Approximately O(n log^2(n)), depending on gapping method.
// Space complexity: O(1)
func ShellSort(ints []int) {
	h := 1

	for h < len(ints) {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < len(ints); i++ {
			for j := i; j >= h && ints[j] < ints[j-h]; j = j - h {
				ints[j], ints[j-h] = ints[j-h], ints[j]
			}
		}
		h = h / 3
	}
}

// ShellSortCiura initiates a shell sort on the supplied int slice using the Ciura gap sequence.
// From pseudocode found at https://en.wikipedia.org/wiki/Shellsort
func ShellSortCiura(ints []int) {
	var i, j, temp int

	for _, gap := range ciuraGaps {
		for i = gap; i < len(ints); i++ {
			temp = ints[i]
			for j = i; j >= gap && ints[j-gap] > temp; j -= gap {
				ints[j] = ints[j-gap]
			}
			ints[j] = temp
		}
	}
}
