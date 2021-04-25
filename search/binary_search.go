package search

import (
	"errors"
	"fmt"
)

func BinarySearch(array []int, item int) (int, error) {
	var err error
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := array[mid]
		if guess == item {
			err = nil
			return mid, err
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	err = errors.New("Not found")
	return 0, err
}

func PrintBinarySearch(index int, err error) {
	if err != nil {
		fmt.Println("Not found")
	} else {
		fmt.Println(index)
	}
}
