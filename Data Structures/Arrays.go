package data_structures

import (
	"math/rand"
	"time"
)

// Initalizing functions
func InitRandomArray(init []int, size int) []int {
	rand.Seed(time.Now().UnixNano())
	// Load 100 random numbers into array
	for i := 0; i < size; i++ {
		init = append(init, rand.Intn(size))
	}
	return init
}

func InitOrderedArray(init []int) []int {
	for i := 0; i < 100; i++ {
		init = append(init, i)
	}
	return init
}
