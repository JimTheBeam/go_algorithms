package sorting

// Merge sort
func Merge(arr []int) []int {
	lenth := len(arr)

	if lenth <= 1 {
		return arr
	}

	mid := lenth / 2
	left := Merge(arr[:mid])
	right := Merge(arr[mid:])

	return mergeLists(left, right)
}

// mergeLists - merge two lists
func mergeLists(left, right []int) []int {
	res := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(res, right...)
		}
		if len(right) == 0 {
			return append(res, left...)
		}

		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	return res
}
