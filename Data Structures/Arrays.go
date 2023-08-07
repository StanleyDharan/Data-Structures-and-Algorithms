package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Data sets to play with
var randomList []int

// Initalizing functions
func initRandomArray(init *[]int) {
	rand.NewSource(time.Now().UnixNano())
	// Load 100 random numbers into array
	for i := 0; i < 100; i++ {
		*init = append(*init, rand.Intn(100))
	}
}

// Sorting functions
// Merge sort

// Main function
func main() {
	initRandomArray(&randomList)
	fmt.Print(randomList)
}
