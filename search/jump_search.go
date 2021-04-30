package search

import (
	"math"
)

// jump search - complexity O(sqrt n)
func JumpSearch(arr []int, item int) (int, error) {
	if len(arr) == 0 {
		return 0, ErrNotFound
	}

	step := int(math.Sqrt(float64(len(arr))))

	var closest int // closest index to item

	for i := 0; i <= (len(arr) - 1); i += step {
		if i+step > len(arr)-1 {
			closest = len(arr) - 1
			break
		}

		if arr[i] == item {
			return i, nil
		}
		if item < arr[i] {
			closest = i
		}
	}

	for i := closest; i >= (closest - step); i-- {
		if arr[i] == item {
			return i, nil
		}
	}
	return 0, ErrNotFound
}
