package search

// O(n)
func LinearSearch(arr []int, item int) (int, error) {
	for index := range arr {
		if arr[index] == item {
			return index, nil
		}
	}
	return 0, ErrNotFound
}
