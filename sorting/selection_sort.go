package sorting

// Selection sort  - complexity O(n^2)
func Selection(arr []int) []int {

	var low int // index of a lowest number

	for n := 0; n < len(arr); n++ {

		low = n
		for i := n + 1; i < len(arr); i++ {
			if arr[low] > arr[i] {
				low = i
			}
		}
		arr[n], arr[low] = arr[low], arr[n]
	}
	return arr
}
