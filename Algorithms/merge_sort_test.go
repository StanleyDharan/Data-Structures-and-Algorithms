package algorithms

import (
	"data_structures"
	"fmt"
	"testing"
)

func TestMergeSort_5_element_slice(t *testing.T) {
	var intSlice []int
	randomIntSlice := data_structures.InitRandomArray(intSlice, 5)

	fmt.Println("Merge Sort 5 Element Array: ")
	fmt.Println("Pre-sort 5 Element Array: ")
	fmt.Println(randomIntSlice)
	fmt.Println("Post Sort 5 Element Array: ")
	MergeSort(randomIntSlice, 0, len(randomIntSlice)-1)
	fmt.Println(randomIntSlice)
}

func TestMergeSort_50_element_slice(t *testing.T) {
	var intSlice []int
	randomIntSlice := data_structures.InitRandomArray(intSlice, 50)

	fmt.Println("Merge Sort 50 Element Array: ")
	fmt.Println("Pre-sort 50 Element Array: ")
	fmt.Println(randomIntSlice)
	fmt.Println("Post Sort 50 Element Array: ")
	MergeSort(randomIntSlice, 0, len(randomIntSlice)-1)
	fmt.Println(randomIntSlice)
}

func TestMergeSort_500_element_slice(t *testing.T) {
	var intSlice []int
	randomIntSlice := data_structures.InitRandomArray(intSlice, 500)

	fmt.Println("Merge Sort 500 Element Array: ")
	fmt.Println("Pre-sort 500 Element Array: ")
	fmt.Println(randomIntSlice)
	fmt.Println("Post Sort 500 Element Array: ")
	MergeSort(randomIntSlice, 0, len(randomIntSlice)-1)
	fmt.Println(randomIntSlice)
}
