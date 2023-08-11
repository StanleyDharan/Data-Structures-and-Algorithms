package algorithms

/*
	Merge Sort Algorithm
	----------------------------
	Steps:
	1. Recursively split list in half until you reach base case of: index[start] == index[end]
	2. Preform merge
		2a.
	3. Merge array back together
*/

func merge(list []int, leftStart, middle, rightEnd int) {
	// create temp arrays isolating left and right sub-arrays
	leftCap := middle - leftStart + 1 // +1 to account for inclusive middle element
	rightCap := rightEnd - middle
	left := make([]int, leftCap)
	right := make([]int, rightCap)

	for i := 0; i < leftCap; i++ {
		left[i] = list[leftStart+i]
	}

	for i := 0; i < rightCap; i++ {
		right[i] = list[middle+i+1]
		// +1 because left array is inclusive of middle element so we can skip it
	}

	// Sorting left and right arrays and adding them to List in order
	leftIndex := 0
	rightIndex := 0
	sortIndex := leftStart

	for leftIndex < leftCap && rightIndex < rightCap {
		if left[leftIndex] < right[rightIndex] {
			list[sortIndex] = left[leftIndex]
			leftIndex++
		} else {
			list[sortIndex] = right[rightIndex]
			rightIndex++
		}
		sortIndex++
	}

	// Add remaining elements because they are already sorted
	for leftIndex < leftCap {
		list[sortIndex] = left[leftIndex]
		leftIndex++
		sortIndex++
	}

	for rightIndex < rightCap {
		list[sortIndex] = right[rightIndex]
		rightIndex++
		sortIndex++
	}
}

func MergeSort(list []int, start int, end int) {
	if list[start] == list[end] {
		return
	}
	mid := (start + end) / 2

	MergeSort(list, start, mid)
	MergeSort(list, mid+1, end)
	merge(list, start, mid, end)
}
