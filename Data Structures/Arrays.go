package main

import (
	"fmt"
	"math/rand"
	"time"
)

var randomList []int

func initRandomArray(init *[]int) {
	rand.NewSource(time.Now().UnixNano())
	// Load 100 random numbers into array
	for i := 0; i < 100; i++ {
		*init = append(*init, rand.Intn(100))
	}
}

func main() {
	initRandomArray(&randomList)
	fmt.Print(randomList)
}
