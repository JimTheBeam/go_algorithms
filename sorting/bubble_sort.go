package sorting

func Bubble(arr []int) []int {
	lenth := len(arr)
	swap := true

	for swap {
		swap = false

		for i := 1; i < lenth; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swap = true
			}
		}
		lenth -= 1
	}
	return arr
}
