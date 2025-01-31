package testdata

import (
	"math/rand"
	"time"
)

func RandomUnsortedInts(count, max int) []int {
	slice := make([]int, count)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		slice[i] = rand.Intn(max) - rand.Intn(max)
	}

	return slice
}
