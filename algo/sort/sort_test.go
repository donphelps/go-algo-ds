package sort_test

import (
	"math/rand"
	"time"
)

func randomUnsortedInts(count int) []int {
	slice := make([]int, count, count)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		slice[i] = rand.Intn(99999) - rand.Intn(99999)
	}

	return slice
}
