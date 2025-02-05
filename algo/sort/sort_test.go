package sort_test

import (
	"math/rand"
	"time"
)

// [] delete me in favor of testdata
func randomUnsortedInts(count, max int) []int {
	slice := make([]int, count)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		slice[i] = rand.Intn(max) - rand.Intn(max)
	}

	return slice
}
