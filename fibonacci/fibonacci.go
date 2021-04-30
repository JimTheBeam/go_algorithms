package fibonacci

func Fibbonacci(n int) int {
	if n == 0 {
		return 0
	}
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1

	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}
