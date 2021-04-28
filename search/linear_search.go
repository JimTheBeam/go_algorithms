package search

import "errors"

// O(n)
func LinearSearch(arr []int, item int) (int, error) {
	for index := range arr {
		if arr[index] == item {
			return index, nil
		}
	}
	return 0, errors.New("Not found")
}
