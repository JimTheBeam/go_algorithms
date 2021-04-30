package search

import (
	"errors"
)

// binary search - complexity O(log n)
func BinarySearch(array []int, item int) (int, error) {
	leftIndex := 0               // left index of array
	rightIndex := len(array) - 1 // right index of array

	for leftIndex <= rightIndex {
		mid := (leftIndex + rightIndex) / 2
		guess := array[mid]
		if guess == item {
			return mid, nil
		}
		if guess > item {
			rightIndex = mid - 1
		} else {
			leftIndex = mid + 1
		}
	}
	return 0, errors.New("Not found")
}
