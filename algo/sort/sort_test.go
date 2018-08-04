package sort_test

import (
	"math/rand"
	"time"
)

func randomUnsortedInts(count, max int) []int {
	slice := make([]int, count, count)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		slice[i] = rand.Intn(max) - rand.Intn(max)
	}

	return slice
}
