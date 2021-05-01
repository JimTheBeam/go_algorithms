package main

import (
	"fmt"
	"go_algorithms/sorting"
)

func main() {
	arr := []int{741, 123, 369, 987, 753, 5, 159, 852, 2, 35, 94, 72, 85, 12, 34, 97, 65}

	sorted := sorting.Selection(arr)
	for _, elem := range sorted {
		fmt.Printf("%d, ", elem)
	}
	fmt.Printf("\n")
	// fmt.Println(sorted)
}
