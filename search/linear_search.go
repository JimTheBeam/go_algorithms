package search

import "errors"

// O(n)
func LinearSearch(arr []int, item int) (int, error) {
	var err error
	for index, elem := range arr {
		if elem == item {
			err = nil
			return index, err
		}
	}
	err = errors.New("Not found")
	return 0, err
}
