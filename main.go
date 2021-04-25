package main

import (
	"go_algorithms/search"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6} //cap = 6
	arr1 := []int{1}
	unsortArr := []int{1, 3, 7, 3, 2, 8, 0, 8, 4} // unsorted array

	search.PrintBinarySearch(search.BinarySearch(arr, 3))       // 2
	search.PrintBinarySearch(search.BinarySearch(arr1, 1))      // 0
	search.PrintBinarySearch(search.BinarySearch(arr, 7))       // not found
	search.PrintBinarySearch(search.BinarySearch(unsortArr, 3)) // not found

}
