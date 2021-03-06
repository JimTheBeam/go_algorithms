package sorting

// Quick - sort algorithm. Average complexity O(n log n)
func Quick(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivot := (left + right) / 2

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]

	Quick(arr[:left])
	Quick(arr[left+1:])

	return arr
}
